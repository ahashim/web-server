package types

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"math/big"

	"github.com/holiman/uint256"
)

// 256 bit unsigned integer type.
type Uint256 struct {
	uint256.Int
}

// Scan interface reads Uint256 values from the database.
func (u *Uint256) Scan(src any) error {
	var i sql.NullString
	if err := i.Scan(src); err != nil {
		return err
	}
	if !i.Valid {
		return nil
	}
	n := new(big.Int)
	n, ok := n.SetString(i.String, 10)
	if !ok {
		return fmt.Errorf("error converting %v to big.Int", src)
	}
	uint256.FromBig(n)
	return fmt.Errorf("could not scan type %T with value %v into Uint256", src, src)
}

// Value interface stores Uint256 values in the database.
func (u *Uint256) Value() (driver.Value, error) {
	return u.String(), nil
}
