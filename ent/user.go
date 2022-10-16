// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/ahashim/web-server/ent/user"
	"github.com/ahashim/web-server/enums"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Address holds the value of the "address" field.
	Address string `json:"address,omitempty"`
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// Status holds the value of the "status" field.
	Status enums.Status `json:"status,omitempty"`
	// ScoutLevel holds the value of the "scout_level" field.
	ScoutLevel int8 `json:"scout_level,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges UserEdges `json:"edges"`
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Followers holds the value of the followers edge.
	Followers []*User `json:"followers,omitempty"`
	// Following holds the value of the following edge.
	Following []*User `json:"following,omitempty"`
	// Roles holds the value of the roles edge.
	Roles []*Role `json:"roles,omitempty"`
	// Created holds the value of the created edge.
	Created []*Squeak `json:"created,omitempty"`
	// Owned holds the value of the owned edge.
	Owned []*Squeak `json:"owned,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
}

// FollowersOrErr returns the Followers value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) FollowersOrErr() ([]*User, error) {
	if e.loadedTypes[0] {
		return e.Followers, nil
	}
	return nil, &NotLoadedError{edge: "followers"}
}

// FollowingOrErr returns the Following value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) FollowingOrErr() ([]*User, error) {
	if e.loadedTypes[1] {
		return e.Following, nil
	}
	return nil, &NotLoadedError{edge: "following"}
}

// RolesOrErr returns the Roles value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) RolesOrErr() ([]*Role, error) {
	if e.loadedTypes[2] {
		return e.Roles, nil
	}
	return nil, &NotLoadedError{edge: "roles"}
}

// CreatedOrErr returns the Created value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) CreatedOrErr() ([]*Squeak, error) {
	if e.loadedTypes[3] {
		return e.Created, nil
	}
	return nil, &NotLoadedError{edge: "created"}
}

// OwnedOrErr returns the Owned value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) OwnedOrErr() ([]*Squeak, error) {
	if e.loadedTypes[4] {
		return e.Owned, nil
	}
	return nil, &NotLoadedError{edge: "owned"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldStatus:
			values[i] = new(enums.Status)
		case user.FieldID, user.FieldScoutLevel:
			values[i] = new(sql.NullInt64)
		case user.FieldAddress, user.FieldUsername:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type User", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int(value.Int64)
		case user.FieldAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address", values[i])
			} else if value.Valid {
				u.Address = value.String
			}
		case user.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				u.Username = value.String
			}
		case user.FieldStatus:
			if value, ok := values[i].(*enums.Status); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value != nil {
				u.Status = *value
			}
		case user.FieldScoutLevel:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field scout_level", values[i])
			} else if value.Valid {
				u.ScoutLevel = int8(value.Int64)
			}
		}
	}
	return nil
}

// QueryFollowers queries the "followers" edge of the User entity.
func (u *User) QueryFollowers() *UserQuery {
	return (&UserClient{config: u.config}).QueryFollowers(u)
}

// QueryFollowing queries the "following" edge of the User entity.
func (u *User) QueryFollowing() *UserQuery {
	return (&UserClient{config: u.config}).QueryFollowing(u)
}

// QueryRoles queries the "roles" edge of the User entity.
func (u *User) QueryRoles() *RoleQuery {
	return (&UserClient{config: u.config}).QueryRoles(u)
}

// QueryCreated queries the "created" edge of the User entity.
func (u *User) QueryCreated() *SqueakQuery {
	return (&UserClient{config: u.config}).QueryCreated(u)
}

// QueryOwned queries the "owned" edge of the User entity.
func (u *User) QueryOwned() *SqueakQuery {
	return (&UserClient{config: u.config}).QueryOwned(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("address=")
	builder.WriteString(u.Address)
	builder.WriteString(", ")
	builder.WriteString("username=")
	builder.WriteString(u.Username)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", u.Status))
	builder.WriteString(", ")
	builder.WriteString("scout_level=")
	builder.WriteString(fmt.Sprintf("%v", u.ScoutLevel))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
