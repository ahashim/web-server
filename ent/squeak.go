// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/ahashim/web-server/ent/squeak"
)

// Squeak is the model entity for the Squeak schema.
type Squeak struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Squeak) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case squeak.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Squeak", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Squeak fields.
func (s *Squeak) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case squeak.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		}
	}
	return nil
}

// Update returns a builder for updating this Squeak.
// Note that you need to call Squeak.Unwrap() before calling this method if this Squeak
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Squeak) Update() *SqueakUpdateOne {
	return (&SqueakClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the Squeak entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Squeak) Unwrap() *Squeak {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Squeak is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Squeak) String() string {
	var builder strings.Builder
	builder.WriteString("Squeak(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Squeaks is a parsable slice of Squeak.
type Squeaks []*Squeak

func (s Squeaks) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
