package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jace996/saas"
	sgin "github.com/jace996/saas/gin"
	mysql2 "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"

	"net/http"

	"github.com/jace996/saas/data"
	sgorm "github.com/jace996/saas/gorm"
	"github.com/jace996/saas/seed"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	defaultSqliteSharedDsn   = "./example.db"
	defaultMysqlSharedDsn    = "root:youShouldChangeThis@tcp(127.0.0.1:3406)/example?parseTime=true&loc=Local"
	defaultPostgresSharedDsn = "host=localhost user=pgsql-saas password=pgsql-saas dbname=pgsql-saas port=5435 sslmode=disable TimeZone=UTC"
)

var (
	driver        string
	sharedDsn     string
	ensureDbExist func(s string) error
)

func init() {
	flag.StringVar(&driver, "driver", "sqlite3", "sqlite3/mysql/pgx")
	flag.StringVar(&sharedDsn, "dsn", "", "shared dsn.")
}

func main() {
	flag.Parse()

	cache := saas.NewCache[string, *sgorm.DbWrap]()
	defer cache.Flush()

	var connStrGen saas.ConnStrGenerator
	switch driver {
	case sqlite.DriverName:
		sharedDsn = defaultSqliteSharedDsn
		connStrGen = saas.NewConnStrGenerator("./example-%s.db")
	case "mysql":
		if len(sharedDsn) == 0 {
			sharedDsn = defaultMysqlSharedDsn
		}
		dd, err := mysql2.ParseDSN(sharedDsn)
		if err != nil {
			panic(err)
		}
		hostDbName := dd.DBName
		dd.DBName = hostDbName + "-%s"
		connStrGen = saas.NewConnStrGenerator(dd.FormatDSN())

		ensureDbExist = func(s string) error {
			dsn, err := mysql2.ParseDSN(s)
			if err != nil {
				return err
			}
			dbname := dsn.DBName
			dsn.DBName = ""
			//open without db name
			db, err := sql.Open(driver, dsn.FormatDSN())
			if err != nil {
				return err
			}
			_, err = db.ExecContext(context.Background(), fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s`", dbname))
			if err != nil {
				return err
			}
			return db.Close()
		}
	case "pgx":
		if len(sharedDsn) == 0 {
			sharedDsn = defaultPostgresSharedDsn
		}

		suffix := "%s"
		modifiedDSN := AddSuffixToDBName(sharedDsn, suffix)
		connStrGen = saas.NewConnStrGenerator(modifiedDSN)

		ensureDbExist = func(s string) error {
			dbname, err := ParseDBNameFromPostgresDSN(s)
			if err != nil {
				return err
			}

			noDbDsn, err := RemoveDBNameFromPostgresDSN(s)
			if err != nil {
				return err
			}

			//open without db name
			db, err := sql.Open(driver, noDbDsn)
			if err != nil {
				return err
			}

			var exists bool
			err = db.QueryRowContext(context.Background(), "SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = $1)", dbname).Scan(&exists)
			if err != nil {
				return err
			}

			// If the database doesn't exist, create it
			if !exists {
				_, err := db.ExecContext(context.Background(), fmt.Sprintf("CREATE DATABASE \"%s\"", dbname))
				if err != nil {
					return err
				}
			}

			return db.Close()
		}

	default:
		panic(fmt.Errorf("driver %s unsupported", driver))
	}

	r := gin.Default()

	conn := make(data.ConnStrings, 1)
	//default database
	conn.SetDefault(sharedDsn)

	clientProvider := sgorm.ClientProviderFunc(func(ctx context.Context, s string) (*gorm.DB, error) {
		client, _, err := cache.GetOrSet(s, func() (*sgorm.DbWrap, error) {
			if ensureDbExist != nil {
				if err := ensureDbExist(s); err != nil {
					return nil, err
				}
			}
			var client *gorm.DB
			var err error
			db, err := sql.Open(driver, s)
			if err != nil {
				return nil, err
			}
			if driver == sqlite.DriverName {
				db.SetMaxIdleConns(1)
				db.SetMaxOpenConns(1)
			}

			if driver == sqlite.DriverName {
				client, err = gorm.Open(&sqlite.Dialector{
					DriverName: sqlite.DriverName,
					DSN:        s,
					Conn:       db,
				})
			} else if driver == "mysql" {
				client, err = gorm.Open(mysql.New(mysql.Config{
					Conn: db,
				}))
			} else if driver == "pgx" {
				client, err = gorm.Open(postgres.New(postgres.Config{
					Conn: db,
				}))
			}

			return sgorm.NewDbWrap(client), err
		})

		if err != nil {
			return nil, err
		}
		return client.WithContext(ctx).Debug(), err

	})
	//tenantStore use connection string from conn
	tenantStore := &TenantStore{dbProvider: sgorm.NewDbProvider(conn, clientProvider)}

	mr := saas.NewMultiTenancyConnStrResolver(tenantStore, conn)

	//tenant dbProvider use connection string from tenant store
	dbProvider := sgorm.NewDbProvider(mr, clientProvider)

	r.Use(sgin.MultiTenancy(tenantStore))

	//return current tenant
	r.GET("/tenant/current", func(c *gin.Context) {
		rCtx := c.Request.Context()
		tenantInfo, _ := saas.FromCurrentTenant(rCtx)
		trR := saas.FromTenantResolveRes(rCtx)
		c.JSON(200, gin.H{
			"tenantId":  tenantInfo.GetId(),
			"resolvers": trR.AppliedResolvers,
		})
	})

	r.GET("/posts", func(c *gin.Context) {
		db := dbProvider.Get(c.Request.Context(), "")
		var entities []Post
		if err := db.Model(&Post{}).Find(&entities).Error; err != nil {
			c.AbortWithError(500, err)
		} else {
			c.JSON(200, entities)
		}
	})

	//seed data into db
	seeder := seed.NewDefaultSeeder(NewMigrationSeeder(dbProvider), NewSeed(dbProvider, connStrGen))
	err := seeder.Seed(context.Background(), seed.AddHost(), seed.AddTenant("1", "2", "3"))
	if err != nil {
		panic(err)
	}

	r.POST("/tenant", func(c *gin.Context) {
		type CreateTenant struct {
			Name       string `form:"name" json:"name" binding:"required"`
			SeparateDb bool   `form:"separateDb" json:"separateDb"`
		}
		var json CreateTenant
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx := c.Request.Context()
		//change to host side
		ctx = saas.NewCurrentTenant(ctx, "", "")
		db := dbProvider.Get(ctx, "")
		t := &Tenant{
			ID:          uuid.New().String(),
			Name:        json.Name,
			DisplayName: json.Name,
		}
		if json.SeparateDb {
			t3Conn, _ := connStrGen.Gen(ctx, saas.NewBasicTenantInfo(t.ID, t.Name))
			t.Conn = []TenantConn{
				{Key: data.Default, Value: t3Conn}, // use tenant3.db
			}
		}
		err := db.Model(t).Create(t).Error
		if err != nil {
			c.AbortWithError(500, err)
			return
		}
		err = seeder.Seed(context.Background(), seed.AddTenant(t.ID))
		if err != nil {
			panic(err)
		}

	})

	r.Run(":8090") // listen and serve on 0.0.0.0:8090 (for windows "localhost:8090")
}
