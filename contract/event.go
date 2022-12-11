package contract

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// event names
const AccountCreatedEventName = "AccountCreated"

// event types
type AccountCreated struct {
	Account  common.Address
	Username [32]byte
}

// event signatures
var AccountCreatedSignature = crypto.Keccak256Hash([]byte("AccountCreated(address,bytes32)")).Hex()
