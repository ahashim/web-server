// Code generated by ent, DO NOT EDIT.

package interaction

import (
	"fmt"

	"github.com/ahashim/web-server/enums"
)

const (
	// Label holds the string label denoting the interaction type in the database.
	Label = "interaction"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeSqueak holds the string denoting the squeak edge name in mutations.
	EdgeSqueak = "squeak"
	// Table holds the table name of the interaction in the database.
	Table = "interactions"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "interactions"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_interactions"
	// SqueakTable is the table that holds the squeak relation/edge.
	SqueakTable = "interactions"
	// SqueakInverseTable is the table name for the Squeak entity.
	// It exists in this package in order to avoid circular dependency with the "squeak" package.
	SqueakInverseTable = "squeaks"
	// SqueakColumn is the table column denoting the squeak relation/edge.
	SqueakColumn = "squeak_interactions"
)

// Columns holds all SQL columns for interaction fields.
var Columns = []string{
	FieldID,
	FieldType,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "interactions"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"squeak_interactions",
	"user_interactions",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type enums.Interaction) error {
	switch _type.String() {
	case "DISLIKE", "LIKE", "RESQUEAK", "UNDO_DISLIKE", "UNDO_LIKE", "UNDO_RESQUEAK":
		return nil
	default:
		return fmt.Errorf("interaction: invalid enum value for type field: %q", _type)
	}
}
