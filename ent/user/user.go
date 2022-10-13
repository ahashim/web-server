// Code generated by ent, DO NOT EDIT.

package user

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAddress holds the string denoting the address field in the database.
	FieldAddress = "address"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldScoutLevel holds the string denoting the scout_level field in the database.
	FieldScoutLevel = "scout_level"
	// Table holds the table name of the user in the database.
	Table = "users"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldAddress,
	FieldUsername,
	FieldStatus,
	FieldScoutLevel,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// AddressValidator is a validator for the "address" field. It is called by the builders before save.
	AddressValidator func(string) error
	// UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	UsernameValidator func(string) error
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus int8
	// DefaultScoutLevel holds the default value on creation for the "scout_level" field.
	DefaultScoutLevel int8
)
