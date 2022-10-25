package enums

import (
	"database/sql/driver"
	"errors"
	"fmt"
)

// Status is the state a user can be in.
type Status uint8

// Enum of user statuses.
const (
	Unknown Status = iota
	Active
	Suspended
	Banned
)

// String provides a string value for a Status int.
func (p Status) String() string {
	switch p {
	case Unknown:
		return "UNKNOWN"
	case Active:
		return "ACTIVE"
	case Suspended:
		return "SUSPENDED"
	case Banned:
		return "BANNED"
	default:
		return ""
	}
}

// Values provides a list of valid values for Status.
func (Status) Values() []string {
	return []string{
		Unknown.String(),
		Active.String(),
		Suspended.String(),
		Banned.String(),
	}
}

// Value provides the DB a string from an int.
func (p Status) Value() (driver.Value, error) {
	return p.String(), nil
}

// Scan determines how to read the enum value into the type.
func (p *Status) Scan(val any) error {
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
	case Unknown.String():
		*p = Unknown
	case Active.String():
		*p = Active
	case Suspended.String():
		*p = Suspended
	case Banned.String():
		*p = Banned
	default:
		return errors.New(fmt.Sprintf("Invalid status: %s", s))
	}

	return nil
}
