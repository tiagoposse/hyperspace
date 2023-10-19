// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/tiagoposse/connect/ent/audit"
	"github.com/tiagoposse/connect/ent/user"
)

// Audit is the model entity for the Audit schema.
type Audit struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Action holds the value of the "action" field.
	Action string `json:"action,omitempty"`
	// Author holds the value of the "author" field.
	Author string `json:"author,omitempty"`
	// Timestamp holds the value of the "timestamp" field.
	Timestamp time.Time `json:"timestamp,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AuditQuery when eager-loading is set.
	Edges        AuditEdges `json:"edges"`
	selectValues sql.SelectValues
}

// AuditEdges holds the relations/edges for other nodes in the graph.
type AuditEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AuditEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Audit) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case audit.FieldID, audit.FieldAction, audit.FieldAuthor:
			values[i] = new(sql.NullString)
		case audit.FieldTimestamp:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Audit fields.
func (a *Audit) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case audit.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				a.ID = value.String
			}
		case audit.FieldAction:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field action", values[i])
			} else if value.Valid {
				a.Action = value.String
			}
		case audit.FieldAuthor:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field author", values[i])
			} else if value.Valid {
				a.Author = value.String
			}
		case audit.FieldTimestamp:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field timestamp", values[i])
			} else if value.Valid {
				a.Timestamp = value.Time
			}
		default:
			a.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Audit.
// This includes values selected through modifiers, order, etc.
func (a *Audit) Value(name string) (ent.Value, error) {
	return a.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the Audit entity.
func (a *Audit) QueryUser() *UserQuery {
	return NewAuditClient(a.config).QueryUser(a)
}

// Update returns a builder for updating this Audit.
// Note that you need to call Audit.Unwrap() before calling this method if this Audit
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Audit) Update() *AuditUpdateOne {
	return NewAuditClient(a.config).UpdateOne(a)
}

// Unwrap unwraps the Audit entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Audit) Unwrap() *Audit {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Audit is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Audit) String() string {
	var builder strings.Builder
	builder.WriteString("Audit(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("action=")
	builder.WriteString(a.Action)
	builder.WriteString(", ")
	builder.WriteString("author=")
	builder.WriteString(a.Author)
	builder.WriteString(", ")
	builder.WriteString("timestamp=")
	builder.WriteString(a.Timestamp.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Audits is a parsable slice of Audit.
type Audits []*Audit
