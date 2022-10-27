// Code generated by ent, DO NOT EDIT.

package scout

const (
	// Label holds the string label denoting the scout type in the database.
	Label = "scout"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldShares holds the string denoting the shares field in the database.
	FieldShares = "shares"
	// Table holds the table name of the scout in the database.
	Table = "scouts"
)

// Columns holds all SQL columns for scout fields.
var Columns = []string{
	FieldID,
	FieldShares,
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
