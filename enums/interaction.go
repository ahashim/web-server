package enums

import (
	"database/sql/driver"
	"fmt"
)

// Interaction is the state a user can be in.
type Interaction uint8

// Enum of user statuses.
const (
	Dislike Interaction = iota
	Like
	Resqueak
	UndoDislike
	UndoLike
	UndoResqueak
)

// String provides a string value for a Interaction int.
func (p Interaction) String() string {
	switch p {
	case Dislike:
		return "DISLIKE"
	case Like:
		return "LIKE"
	case Resqueak:
		return "RESQUEAK"
	case UndoDislike:
		return "UNDO_DISLIKE"
	case UndoLike:
		return "UNDO_LIKE"
	case UndoResqueak:
		return "UNDO_RESQUEAK"
	default:
		return ""
	}
}

// Values provides a list of valid values for Interaction.
func (Interaction) Values() []string {
	return []string{
		Dislike.String(),
		Like.String(),
		Resqueak.String(),
		UndoDislike.String(),
		UndoLike.String(),
		UndoResqueak.String(),
	}
}

// Value provides the DB a string from an int.
func (p Interaction) Value() (driver.Value, error) {
	return p.String(), nil
}

// Scan determines how to read the enum value into the type.
func (p *Interaction) Scan(val any) error {
	var s string

	switch v := val.(type) {
	case nil:
		return nil
	case string:
		s = v
	case []uint8:
		s = string(v)
	}

	switch s {
	case Dislike.String():
		*p = Dislike
	case Like.String():
		*p = Like
	case Resqueak.String():
		*p = Resqueak
	case UndoDislike.String():
		*p = UndoDislike
	case UndoLike.String():
		*p = UndoLike
	case UndoResqueak.String():
		*p = UndoResqueak
	default:
		return fmt.Errorf("Invalid interaction: %s", s)
	}

	return nil
}
