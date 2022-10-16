// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/ahashim/web-server/ent/squeak"
	"github.com/ahashim/web-server/ent/user"
	"github.com/ahashim/web-server/types"
)

// Squeak is the model entity for the Squeak schema.
type Squeak struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// BlockNumber holds the value of the "block_number" field.
	BlockNumber *types.Uint256 `json:"block_number,omitempty"`
	// Content holds the value of the "content" field.
	Content string `json:"content,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SqueakQuery when eager-loading is set.
	Edges        SqueakEdges `json:"edges"`
	user_created *int
	user_owned   *int
}

// SqueakEdges holds the relations/edges for other nodes in the graph.
type SqueakEdges struct {
	// Interactions holds the value of the interactions edge.
	Interactions []*Interaction `json:"interactions,omitempty"`
	// Creator holds the value of the creator edge.
	Creator *User `json:"creator,omitempty"`
	// Owner holds the value of the owner edge.
	Owner *User `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// InteractionsOrErr returns the Interactions value or an error if the edge
// was not loaded in eager-loading.
func (e SqueakEdges) InteractionsOrErr() ([]*Interaction, error) {
	if e.loadedTypes[0] {
		return e.Interactions, nil
	}
	return nil, &NotLoadedError{edge: "interactions"}
}

// CreatorOrErr returns the Creator value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SqueakEdges) CreatorOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.Creator == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Creator, nil
	}
	return nil, &NotLoadedError{edge: "creator"}
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SqueakEdges) OwnerOrErr() (*User, error) {
	if e.loadedTypes[2] {
		if e.Owner == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Squeak) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case squeak.FieldID:
			values[i] = new(sql.NullInt64)
		case squeak.FieldContent:
			values[i] = new(sql.NullString)
		case squeak.FieldBlockNumber:
			values[i] = new(types.Uint256)
		case squeak.ForeignKeys[0]: // user_created
			values[i] = new(sql.NullInt64)
		case squeak.ForeignKeys[1]: // user_owned
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
		case squeak.FieldBlockNumber:
			if value, ok := values[i].(*types.Uint256); !ok {
				return fmt.Errorf("unexpected type %T for field block_number", values[i])
			} else if value != nil {
				s.BlockNumber = value
			}
		case squeak.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value.Valid {
				s.Content = value.String
			}
		case squeak.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_created", value)
			} else if value.Valid {
				s.user_created = new(int)
				*s.user_created = int(value.Int64)
			}
		case squeak.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_owned", value)
			} else if value.Valid {
				s.user_owned = new(int)
				*s.user_owned = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryInteractions queries the "interactions" edge of the Squeak entity.
func (s *Squeak) QueryInteractions() *InteractionQuery {
	return (&SqueakClient{config: s.config}).QueryInteractions(s)
}

// QueryCreator queries the "creator" edge of the Squeak entity.
func (s *Squeak) QueryCreator() *UserQuery {
	return (&SqueakClient{config: s.config}).QueryCreator(s)
}

// QueryOwner queries the "owner" edge of the Squeak entity.
func (s *Squeak) QueryOwner() *UserQuery {
	return (&SqueakClient{config: s.config}).QueryOwner(s)
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
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("block_number=")
	builder.WriteString(fmt.Sprintf("%v", s.BlockNumber))
	builder.WriteString(", ")
	builder.WriteString("content=")
	builder.WriteString(s.Content)
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
