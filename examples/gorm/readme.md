# Example project

combination of `jace996`,`gin`,`gorm(sqlite/mysql)`

### sqlite3
```shell
go run github.com/jace996/saas/examples/gorm
```
---
### mysql
```shell
docker-compose up -d
go run github.com/jace996/saas/examples/gorm --driver mysql
```
---
### postgres
```shell
docker-compose up -d
go run github.com/jace996/saas/examples/gorm --driver pgx
```

Host side ( use shared database):

Open http://localhost:8090/posts

---
Multi-tenancy ( use shared database):

Open http://localhost:8090/posts?__tenant=1

Open http://localhost:8090/posts?__tenant=2

---
Single-tenancy ( use separate database):

Open http://localhost:8090/posts?__tenant=3

---

Create tenant
```shell
curl -H "Accept: application/json" -H "Content-type: application/json" -X POST -d '{"name":"newTenant","separateDb":true}' http://localhost:8090/tenant
```
Open http://localhost:8090/posts?__tenant=newTenant
