// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/jace996/saas/examples/ent/shared/ent/tenantconn"
)

// TenantConn is the model entity for the TenantConn schema.
type TenantConn struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Key holds the value of the "key" field.
	Key string `json:"key,omitempty"`
	// Value holds the value of the "value" field.
	Value        string `json:"value,omitempty"`
	tenant_conn  *int
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TenantConn) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case tenantconn.FieldID:
			values[i] = new(sql.NullInt64)
		case tenantconn.FieldKey, tenantconn.FieldValue:
			values[i] = new(sql.NullString)
		case tenantconn.FieldCreateTime, tenantconn.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case tenantconn.ForeignKeys[0]: // tenant_conn
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TenantConn fields.
func (tc *TenantConn) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case tenantconn.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			tc.ID = int(value.Int64)
		case tenantconn.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				tc.CreateTime = value.Time
			}
		case tenantconn.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				tc.UpdateTime = value.Time
			}
		case tenantconn.FieldKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field key", values[i])
			} else if value.Valid {
				tc.Key = value.String
			}
		case tenantconn.FieldValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field value", values[i])
			} else if value.Valid {
				tc.Value = value.String
			}
		case tenantconn.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field tenant_conn", value)
			} else if value.Valid {
				tc.tenant_conn = new(int)
				*tc.tenant_conn = int(value.Int64)
			}
		default:
			tc.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// GetValue returns the ent.Value that was dynamically selected and assigned to the TenantConn.
// This includes values selected through modifiers, order, etc.
func (tc *TenantConn) GetValue(name string) (ent.Value, error) {
	return tc.selectValues.Get(name)
}

// Update returns a builder for updating this TenantConn.
// Note that you need to call TenantConn.Unwrap() before calling this method if this TenantConn
// was returned from a transaction, and the transaction was committed or rolled back.
func (tc *TenantConn) Update() *TenantConnUpdateOne {
	return NewTenantConnClient(tc.config).UpdateOne(tc)
}

// Unwrap unwraps the TenantConn entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (tc *TenantConn) Unwrap() *TenantConn {
	_tx, ok := tc.config.driver.(*txDriver)
	if !ok {
		panic("ent: TenantConn is not a transactional entity")
	}
	tc.config.driver = _tx.drv
	return tc
}

// String implements the fmt.Stringer.
func (tc *TenantConn) String() string {
	var builder strings.Builder
	builder.WriteString("TenantConn(")
	builder.WriteString(fmt.Sprintf("id=%v, ", tc.ID))
	builder.WriteString("create_time=")
	builder.WriteString(tc.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(tc.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("key=")
	builder.WriteString(tc.Key)
	builder.WriteString(", ")
	builder.WriteString("value=")
	builder.WriteString(tc.Value)
	builder.WriteByte(')')
	return builder.String()
}

// TenantConns is a parsable slice of TenantConn.
type TenantConns []*TenantConn
