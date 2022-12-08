// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package critter

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// IStoreablePoolInfo is an auto generated low-level Go binding around an user-defined struct.
type IStoreablePoolInfo struct {
	Amount      *big.Int
	Shares      *big.Int
	PassCount   *big.Int
	BlockNumber *big.Int
	Score       uint64
}

// IStoreablePoolPassInfo is an auto generated low-level Go binding around an user-defined struct.
type IStoreablePoolPassInfo struct {
	Account common.Address
	Shares  *big.Int
}

// IStoreableSentimentCounts is an auto generated low-level Go binding around an user-defined struct.
type IStoreableSentimentCounts struct {
	Dislikes  *big.Int
	Likes     *big.Int
	Resqueaks *big.Int
}

// CritterMetaData contains all meta data concerning the Critter contract.
var CritterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyBlocked\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyFollowing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyInteracted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ApprovalCallerNotOwnerNorApproved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ApprovalQueryForNonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BalanceQueryForZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Blocked\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientFunds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAccountStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRelationship\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintERC2309QuantityExceedsLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintToZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintZeroQuantity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotApprovedOrOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotBlocked\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotFollowing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInPool\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInteractedYet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerQueryForNonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnershipNotInitializedForExtraData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SqueakDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SqueakEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SqueakTooLong\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferCallerNotOwnerNorApproved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFromIncorrectOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToNonERC721ReceiverImplementer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"URIQueryForNonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UsernameEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UsernameTooLong\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UsernameUnavailable\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"username\",\"type\":\"bytes32\"}],\"name\":\"AccountCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newUsername\",\"type\":\"string\"}],\"name\":\"AccountUsernameUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toTokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"ConsecutiveTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FundsAddedToPool\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FundsDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FundsTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FundsWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumInteraction\",\"name\":\"interaction\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"InteractionFeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"PoolPayout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"relative\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumRelation\",\"name\":\"action\",\"type\":\"uint8\"}],\"name\":\"RelationshipUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"author\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"}],\"name\":\"SqueakCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"deletedBy\",\"type\":\"address\"}],\"name\":\"SqueakDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumInteraction\",\"name\":\"interaction\",\"type\":\"uint8\"}],\"name\":\"SqueakInteraction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"StatusUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"addresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseTokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumConfiguration\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"config\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"username\",\"type\":\"string\"}],\"name\":\"createAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"}],\"name\":\"createSqueak\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"deleteSqueak\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumInteraction\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"fees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getDeleteFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getPoolInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"passCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"score\",\"type\":\"uint64\"}],\"internalType\":\"structIStoreable.PoolInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getPoolPasses\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"internalType\":\"structIStoreable.PoolPassInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getSentimentCounts\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"dislikes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"likes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resqueaks\",\"type\":\"uint256\"}],\"internalType\":\"structIStoreable.SentimentCounts\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getViralityScore\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"baseURI\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"platformFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"platformTakeRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"poolPayoutThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"viralityThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"viralityBonus\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxLevel\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"enumInteraction\",\"name\":\"interaction\",\"type\":\"uint8\"}],\"name\":\"interact\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userOne\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"userTwo\",\"type\":\"address\"}],\"name\":\"isBlocked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userOne\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"userTwo\",\"type\":\"address\"}],\"name\":\"isFollowing\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"isViral\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"leavePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"squeaks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"author\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"content\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasury\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumConfiguration\",\"name\":\"configuration\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"updateConfiguration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumInteraction\",\"name\":\"interaction\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"updateInteractionFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"enumRelation\",\"name\":\"action\",\"type\":\"uint8\"}],\"name\":\"updateRelationship\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"enumStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"updateStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newUsername\",\"type\":\"string\"}],\"name\":\"updateUsername\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"users\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"enumStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"level\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"username\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// CritterABI is the input ABI used to generate the binding from.
// Deprecated: Use CritterMetaData.ABI instead.
var CritterABI = CritterMetaData.ABI

// Critter is an auto generated Go binding around an Ethereum contract.
type Critter struct {
	CritterCaller     // Read-only binding to the contract
	CritterTransactor // Write-only binding to the contract
	CritterFilterer   // Log filterer for contract events
}

// CritterCaller is an auto generated read-only Go binding around an Ethereum contract.
type CritterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CritterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CritterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CritterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CritterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CritterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CritterSession struct {
	Contract     *Critter          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CritterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CritterCallerSession struct {
	Contract *CritterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// CritterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CritterTransactorSession struct {
	Contract     *CritterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// CritterRaw is an auto generated low-level Go binding around an Ethereum contract.
type CritterRaw struct {
	Contract *Critter // Generic contract binding to access the raw methods on
}

// CritterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CritterCallerRaw struct {
	Contract *CritterCaller // Generic read-only contract binding to access the raw methods on
}

// CritterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CritterTransactorRaw struct {
	Contract *CritterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCritter creates a new instance of Critter, bound to a specific deployed contract.
func NewCritter(address common.Address, backend bind.ContractBackend) (*Critter, error) {
	contract, err := bindCritter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Critter{CritterCaller: CritterCaller{contract: contract}, CritterTransactor: CritterTransactor{contract: contract}, CritterFilterer: CritterFilterer{contract: contract}}, nil
}

// NewCritterCaller creates a new read-only instance of Critter, bound to a specific deployed contract.
func NewCritterCaller(address common.Address, caller bind.ContractCaller) (*CritterCaller, error) {
	contract, err := bindCritter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CritterCaller{contract: contract}, nil
}

// NewCritterTransactor creates a new write-only instance of Critter, bound to a specific deployed contract.
func NewCritterTransactor(address common.Address, transactor bind.ContractTransactor) (*CritterTransactor, error) {
	contract, err := bindCritter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CritterTransactor{contract: contract}, nil
}

// NewCritterFilterer creates a new log filterer instance of Critter, bound to a specific deployed contract.
func NewCritterFilterer(address common.Address, filterer bind.ContractFilterer) (*CritterFilterer, error) {
	contract, err := bindCritter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CritterFilterer{contract: contract}, nil
}

// bindCritter binds a generic wrapper to an already deployed contract.
func bindCritter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CritterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Critter *CritterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Critter.Contract.CritterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Critter *CritterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Critter.Contract.CritterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Critter *CritterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Critter.Contract.CritterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Critter *CritterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Critter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Critter *CritterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Critter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Critter *CritterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Critter.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Critter *CritterCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Critter.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Critter *CritterSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Critter.Contract.DEFAULTADMINROLE(&_Critter.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Critter *CritterCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Critter.Contract.DEFAULTADMINROLE(&_Critter.CallOpts)
}

// Addresses is a free data retrieval call binding the contract method 0xbdfe7d47.
//
// Solidity: function addresses(string ) view returns(address)
func (_Critter *CritterCaller) Addresses(opts *bind.CallOpts, arg0 string) (common.Address, error) {
	var out []interface{}
	err := _Critter.contract.Call(opts, &out, "addresses", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Addresses is a free data retrieval call binding the contract method 0xbdfe7d47.
//
// Solidity: function addresses(string ) view returns(address)
func (_Critter *CritterSession) Addresses(arg0 string) (common.Address, error) {
	return _Critter.Contract.Addresses(&_Critter.CallOpts, arg0)
}

// Addresses is a free data retrieval call binding the contract method 0xbdfe7d47.
//
// Solidity: function addresses(string ) view returns(address)
func (_Critter *CritterCallerSession) Addresses(arg0 string) (common.Address, error) {
	return _Critter.Contract.Addresses(&_Critter.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Critter *CritterCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Critter.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Critter *CritterSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Critter.Contract.BalanceOf(&_Critter.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Critter *CritterCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Critter.Contract.BalanceOf(&_Critter.CallOpts, owner)
}

// BaseTokenURI is a free data retrieval call binding the contract method 0xd547cfb7.
//
// Solidity: function baseTokenURI() view returns(string)
func (_Critter *CritterCaller) BaseTokenURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Critter.contract.Call(opts, &out, "baseTokenURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BaseTokenURI is a free data retrieval call binding the contract method 0xd547cfb7.
//
// Solidity: function baseTokenURI() view returns(string)
func (_Critter *CritterSession) BaseTokenURI() (string, error) {
	return _Critter.Contract.BaseTokenURI(&_Critter.CallOpts)
}

// BaseTokenURI is a free data retrieval call binding the contract method 0xd547cfb7.
//
// Solidity: function baseTokenURI() view returns(string)
func (_Critter *CritterCallerSession) BaseTokenURI() (string, error) {
	return _Critter.Contract.BaseTokenURI(&_Critter.CallOpts)
}

// Config is a free data retrieval call binding the contract method 0x89d62952.
//
// Solidity: function config(uint8 ) view returns(uint256)
func (_Critter *CritterCaller) Config(opts *bind.CallOpts, arg0 uint8) (*big.Int, error) {
	var out []interface{}
	err := _Critter.contract.Call(opts, &out, "config", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Config is a free data retrieval call binding the contract method 0x89d62952.
//
// Solidity: function config(uint8 ) view returns(uint256)
func (_Critter *CritterSession) Config(arg0 uint8) (*big.Int, error) {
	return _Critter.Contract.Config(&_Critter.CallOpts, arg0)
}

// Config is a free data retrieval call binding the contract method 0x89d62952.
//
// Solidity: function config(uint8 ) view returns(uint256)
func (_Critter *CritterCallerSession) Config(arg0 uint8) (*big.Int, error) {
	return _Critter.Contract.Config(&_Critter.CallOpts, arg0)
}

// Fees is a free data retrieval call binding the contract method 0x357c1354.
//
// Solidity: function fees(uint8 ) view returns(uint256)
func (_Critter *CritterCaller) Fees(opts *bind.CallOpts, arg0 uint8) (*big.Int, error) {
	var out []interface{}
	err := _Critter.contract.Call(opts, &out, "fees", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fees is a free data retrieval call binding the contract method 0x357c1354.
//
// Solidity: function fees(uint8 ) view returns(uint256)
func (_Critter *CritterSession) Fees(arg0 uint8) (*big.Int, error) {
	return _Critter.Contract.Fees(&_Critter.CallOpts, arg0)
}

// Fees is a free data retrieval call binding the contract method 0x357c1354.
//
// Solidity: function fees(uint8 ) view returns(uint256)
func (_Critter *CritterCallerSession) Fees(arg0 uint8) (*big.Int, error) {
	return _Critter.Contract.Fees(&_Critter.CallOpts, arg0)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Critter *CritterCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Critter.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Critter *CritterSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Critter.Contract.GetApproved(&_Critter.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Critter *CritterCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Critter.Contract.GetApproved(&_Critter.CallOpts, tokenId)
}

// GetDeleteFee is a free data retrieval call binding the contract method 0x7fcac9c3.
//
// Solidity: function getDeleteFee(uint256 tokenId) view returns(uint256)
func (_Critter *CritterCaller) GetDeleteFee(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Critter.contract.Call(opts, &out, "getDeleteFee", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDeleteFee is a free data retrieval call binding the contract method 0x7fcac9c3.
//
// Solidity: function getDeleteFee(uint256 tokenId) view returns(uint256)
func (_Critter *CritterSession) GetDeleteFee(tokenId *big.Int) (*big.Int, error) {
	return _Critter.Contract.GetDeleteFee(&_Critter.CallOpts, tokenId)
}

// GetDeleteFee is a free data retrieval call binding the contract method 0x7fcac9c3.
//
// Solidity: function getDeleteFee(uint256 tokenId) view returns(uint256)
func (_Critter *CritterCallerSession) GetDeleteFee(tokenId *big.Int) (*big.Int, error) {
	return _Critter.Contract.GetDeleteFee(&_Critter.CallOpts, tokenId)
}

// GetPoolInfo is a free data retrieval call binding the contract method 0x2f380b35.
//
// Solidity: function getPoolInfo(uint256 tokenId) view returns((uint256,uint256,uint256,uint256,uint64))
func (_Critter *CritterCaller) GetPoolInfo(opts *bind.CallOpts, tokenId *big.Int) (IStoreablePoolInfo, error) {
	var out []interface{}
	err := _Critter.contract.Call(opts, &out, "getPoolInfo", tokenId)

	if err != nil {
		return *new(IStoreablePoolInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IStoreablePoolInfo)).(*IStoreablePoolInfo)

	return out0, err

}

// GetPoolInfo is a free data retrieval call binding the contract method 0x2f380b35.
//
// Solidity: function getPoolInfo(uint256 tokenId) view returns((uint256,uint256,uint256,uint256,uint64))
func (_Critter *CritterSession) GetPoolInfo(tokenId *big.Int) (IStoreablePoolInfo, error) {
	return _Critter.Contract.GetPoolInfo(&_Critter.CallOpts, tokenId)
}

// GetPoolInfo is a free data retrieval call binding the contract method 0x2f380b35.
//
// Solidity: function getPoolInfo(uint256 tokenId) view returns((uint256,uint256,uint256,uint256,uint64))
func (_Critter *CritterCallerSession) GetPoolInfo(tokenId *big.Int) (IStoreablePoolInfo, error) {
	return _Critter.Contract.GetPoolInfo(&_Critter.CallOpts, tokenId)
}

// GetPoolPasses is a free data retrieval call binding the contract method 0x402b7b56.
//
// Solidity: function getPoolPasses(uint256 tokenId) view returns((address,uint256)[])
func (_Critter *CritterCaller) GetPoolPasses(opts *bind.CallOpts, tokenId *big.Int) ([]IStoreablePoolPassInfo, error) {
	var out []interface{}
	err := _Critter.contract.Call(opts, &out, "getPoolPasses", tokenId)

	if err != nil {
		return *new([]IStoreablePoolPassInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IStoreablePoolPassInfo)).(*[]IStoreablePoolPassInfo)

	return out0, err

}

// GetPoolPasses is a free data retrieval call binding the contract method 0x402b7b56.
//
// Solidity: function getPoolPasses(uint256 tokenId) view returns((address,uint256)[])
func (_Critter *CritterSession) GetPoolPasses(tokenId *big.Int) ([]IStoreablePoolPassInfo, error) {
	return _Critter.Contract.GetPoolPasses(&_Critter.CallOpts, tokenId)
}

// GetPoolPasses is a free data retrieval call binding the contract method 0x402b7b56.
//
// Solidity: function getPoolPasses(uint256 tokenId) view returns((address,uint256)[])
func (_Critter *CritterCallerSession) GetPoolPasses(tokenId *big.Int) ([]IStoreablePoolPassInfo, error) {
	return _Critter.Contract.GetPoolPasses(&_Critter.CallOpts, tokenId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Critter *CritterCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Critter.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Critter *CritterSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Critter.Contract.GetRoleAdmin(&_Critter.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Critter *CritterCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Critter.Contract.GetRoleAdmin(&_Critter.CallOpts, role)
}

// GetSentimentCounts is a free data retrieval call binding the contract method 0x8d01e712.
//
// Solidity: function getSentimentCounts(uint256 tokenId) view returns((uint256,uint256,uint256))
func (_Critter *CritterCaller) GetSentimentCounts(opts *bind.CallOpts, tokenId *big.Int) (IStoreableSentimentCounts, error) {
	var out []interface{}
	err := _Critter.contract.Call(opts, &out, "getSentimentCounts", tokenId)

	if err != nil {
		return *new(IStoreableSentimentCounts), err
	}

	out0 := *abi.ConvertType(out[0], new(IStoreableSentimentCounts)).(*IStoreableSentimentCounts)

	return out0, err

}

// GetSentimentCounts is a free data retrieval call binding the contract method 0x8d01e712.
//
// Solidity: function getSentimentCounts(uint256 tokenId) view returns((uint256,uint256,uint256))
func (_Critter *CritterSession) GetSentimentCounts(tokenId *big.Int) (IStoreableSentimentCounts, error) {
	return _Critter.Contract.GetSentimentCounts(&_Critter.CallOpts, tokenId)
}

// GetSentimentCounts is a free data retrieval call binding the contract method 0x8d01e712.
//
// Solidity: function getSentimentCounts(uint256 tokenId) view returns((uint256,uint256,uint256))
func (_Critter *CritterCallerSession) GetSentimentCounts(tokenId *big.Int) (IStoreableSentimentCounts, error) {
	return _Critter.Contract.GetSentimentCounts(&_Critter.CallOpts, tokenId)
}

// GetViralityScore is a free data retrieval call binding the contract method 0xc3fb2616.
//
// Solidity: function getViralityScore(uint256 tokenId) view returns(uint64)
func (_Critter *CritterCaller) GetViralityScore(opts *bind.CallOpts, tokenId *big.Int) (uint64, error) {
	var out []interface{}
	err := _Critter.contract.Call(opts, &out, "getViralityScore", tokenId)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetViralityScore is a free data retrieval call binding the contract method 0xc3fb2616.
//
// Solidity: function getViralityScore(uint256 tokenId) view returns(uint64)
func (_Critter *CritterSession) GetViralityScore(tokenId *big.Int) (uint64, error) {
	return _Critter.Contract.GetViralityScore(&_Critter.CallOpts, tokenId)
}

// GetViralityScore is a free data retrieval call binding the contract method 0xc3fb2616.
//
// Solidity: function getViralityScore(uint256 tokenId) view returns(uint64)
func (_Critter *CritterCallerSession) GetViralityScore(tokenId *big.Int) (uint64, error) {
	return _Critter.Contract.GetViralityScore(&_Critter.CallOpts, tokenId)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Critter *CritterCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Critter.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Critter *CritterSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Critter.Contract.HasRole(&_Critter.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Critter *CritterCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Critter.Contract.HasRole(&_Critter.CallOpts, role, account)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Critter *CritterCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Critter.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Critter *CritterSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Critter.Contract.IsApprovedForAll(&_Critter.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Critter *CritterCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Critter.Contract.IsApprovedForAll(&_Critter.CallOpts, owner, operator)
}

// IsBlocked is a free data retrieval call binding the contract method 0x86c58d3e.
//
// Solidity: function isBlocked(address userOne, address userTwo) view returns(bool)
func (_Critter *CritterCaller) IsBlocked(opts *bind.CallOpts, userOne common.Address, userTwo common.Address) (bool, error) {
	var out []interface{}
	err := _Critter.contract.Call(opts, &out, "isBlocked", userOne, userTwo)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBlocked is a free data retrieval call binding the contract method 0x86c58d3e.
//
// Solidity: function isBlocked(address userOne, address userTwo) view returns(bool)
func (_Critter *CritterSession) IsBlocked(userOne common.Address, userTwo common.Address) (bool, error) {
	return _Critter.Contract.IsBlocked(&_Critter.CallOpts, userOne, userTwo)
}

// IsBlocked is a free data retrieval call binding the contract method 0x86c58d3e.
//
// Solidity: function isBlocked(address userOne, address userTwo) view returns(bool)
func (_Critter *CritterCallerSession) IsBlocked(userOne common.Address, userTwo common.Address) (bool, error) {
	return _Critter.Contract.IsBlocked(&_Critter.CallOpts, userOne, userTwo)
}

// IsFollowing is a free data retrieval call binding the contract method 0x99ec3a42.
//
// Solidity: function isFollowing(address userOne, address userTwo) view returns(bool)
func (_Critter *CritterCaller) IsFollowing(opts *bind.CallOpts, userOne common.Address, userTwo common.Address) (bool, error) {
	var out []interface{}
	err := _Critter.contract.Call(opts, &out, "isFollowing", userOne, userTwo)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFollowing is a free data retrieval call binding the contract method 0x99ec3a42.
//
// Solidity: function isFollowing(address userOne, address userTwo) view returns(bool)
func (_Critter *CritterSession) IsFollowing(userOne common.Address, userTwo common.Address) (bool, error) {
	return _Critter.Contract.IsFollowing(&_Critter.CallOpts, userOne, userTwo)
}

// IsFollowing is a free data retrieval call binding the contract method 0x99ec3a42.
//
// Solidity: function isFollowing(address userOne, address userTwo) view returns(bool)
func (_Critter *CritterCallerSession) IsFollowing(userOne common.Address, userTwo common.Address) (bool, error) {
	return _Critter.Contract.IsFollowing(&_Critter.CallOpts, userOne, userTwo)
}

// IsViral is a free data retrieval call binding the contract method 0xeb7b8ee6.
//
// Solidity: function isViral(uint256 tokenId) view returns(bool)
func (_Critter *CritterCaller) IsViral(opts *bind.CallOpts, tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _Critter.contract.Call(opts, &out, "isViral", tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsViral is a free data retrieval call binding the contract method 0xeb7b8ee6.
//
// Solidity: function isViral(uint256 tokenId) view returns(bool)
func (_Critter *CritterSession) IsViral(tokenId *big.Int) (bool, error) {
	return _Critter.Contract.IsViral(&_Critter.CallOpts, tokenId)
}

// IsViral is a free data retrieval call binding the contract method 0xeb7b8ee6.
//
// Solidity: function isViral(uint256 tokenId) view returns(bool)
func (_Critter *CritterCallerSession) IsViral(tokenId *big.Int) (bool, error) {
	return _Critter.Contract.IsViral(&_Critter.CallOpts, tokenId)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Critter *CritterCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Critter.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Critter *CritterSession) Name() (string, error) {
	return _Critter.Contract.Name(&_Critter.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Critter *CritterCallerSession) Name() (string, error) {
	return _Critter.Contract.Name(&_Critter.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Critter *CritterCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Critter.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Critter *CritterSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Critter.Contract.OwnerOf(&_Critter.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Critter *CritterCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Critter.Contract.OwnerOf(&_Critter.CallOpts, tokenId)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Critter *CritterCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Critter.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Critter *CritterSession) ProxiableUUID() ([32]byte, error) {
	return _Critter.Contract.ProxiableUUID(&_Critter.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Critter *CritterCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Critter.Contract.ProxiableUUID(&_Critter.CallOpts)
}

// Squeaks is a free data retrieval call binding the contract method 0xf7e75ccb.
//
// Solidity: function squeaks(uint256 ) view returns(uint256 blockNumber, address author, address owner, bytes content)
func (_Critter *CritterCaller) Squeaks(opts *bind.CallOpts, arg0 *big.Int) (struct {
	BlockNumber *big.Int
	Author      common.Address
	Owner       common.Address
	Content     []byte
}, error) {
	var out []interface{}
	err := _Critter.contract.Call(opts, &out, "squeaks", arg0)

	outstruct := new(struct {
		BlockNumber *big.Int
		Author      common.Address
		Owner       common.Address
		Content     []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BlockNumber = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Author = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Owner = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Content = *abi.ConvertType(out[3], new([]byte)).(*[]byte)

	return *outstruct, err

}

// Squeaks is a free data retrieval call binding the contract method 0xf7e75ccb.
//
// Solidity: function squeaks(uint256 ) view returns(uint256 blockNumber, address author, address owner, bytes content)
func (_Critter *CritterSession) Squeaks(arg0 *big.Int) (struct {
	BlockNumber *big.Int
	Author      common.Address
	Owner       common.Address
	Content     []byte
}, error) {
	return _Critter.Contract.Squeaks(&_Critter.CallOpts, arg0)
}

// Squeaks is a free data retrieval call binding the contract method 0xf7e75ccb.
//
// Solidity: function squeaks(uint256 ) view returns(uint256 blockNumber, address author, address owner, bytes content)
func (_Critter *CritterCallerSession) Squeaks(arg0 *big.Int) (struct {
	BlockNumber *big.Int
	Author      common.Address
	Owner       common.Address
	Content     []byte
}, error) {
	return _Critter.Contract.Squeaks(&_Critter.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Critter *CritterCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Critter.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Critter *CritterSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Critter.Contract.SupportsInterface(&_Critter.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Critter *CritterCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Critter.Contract.SupportsInterface(&_Critter.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Critter *CritterCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Critter.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Critter *CritterSession) Symbol() (string, error) {
	return _Critter.Contract.Symbol(&_Critter.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Critter *CritterCallerSession) Symbol() (string, error) {
	return _Critter.Contract.Symbol(&_Critter.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Critter *CritterCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Critter.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Critter *CritterSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Critter.Contract.TokenURI(&_Critter.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Critter *CritterCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Critter.Contract.TokenURI(&_Critter.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Critter *CritterCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Critter.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Critter *CritterSession) TotalSupply() (*big.Int, error) {
	return _Critter.Contract.TotalSupply(&_Critter.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Critter *CritterCallerSession) TotalSupply() (*big.Int, error) {
	return _Critter.Contract.TotalSupply(&_Critter.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(uint256)
func (_Critter *CritterCaller) Treasury(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Critter.contract.Call(opts, &out, "treasury")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(uint256)
func (_Critter *CritterSession) Treasury() (*big.Int, error) {
	return _Critter.Contract.Treasury(&_Critter.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(uint256)
func (_Critter *CritterCallerSession) Treasury() (*big.Int, error) {
	return _Critter.Contract.Treasury(&_Critter.CallOpts)
}

// Users is a free data retrieval call binding the contract method 0xa87430ba.
//
// Solidity: function users(address ) view returns(address account, uint8 status, uint256 level, string username)
func (_Critter *CritterCaller) Users(opts *bind.CallOpts, arg0 common.Address) (struct {
	Account  common.Address
	Status   uint8
	Level    *big.Int
	Username string
}, error) {
	var out []interface{}
	err := _Critter.contract.Call(opts, &out, "users", arg0)

	outstruct := new(struct {
		Account  common.Address
		Status   uint8
		Level    *big.Int
		Username string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Account = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Status = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.Level = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Username = *abi.ConvertType(out[3], new(string)).(*string)

	return *outstruct, err

}

// Users is a free data retrieval call binding the contract method 0xa87430ba.
//
// Solidity: function users(address ) view returns(address account, uint8 status, uint256 level, string username)
func (_Critter *CritterSession) Users(arg0 common.Address) (struct {
	Account  common.Address
	Status   uint8
	Level    *big.Int
	Username string
}, error) {
	return _Critter.Contract.Users(&_Critter.CallOpts, arg0)
}

// Users is a free data retrieval call binding the contract method 0xa87430ba.
//
// Solidity: function users(address ) view returns(address account, uint8 status, uint256 level, string username)
func (_Critter *CritterCallerSession) Users(arg0 common.Address) (struct {
	Account  common.Address
	Status   uint8
	Level    *big.Int
	Username string
}, error) {
	return _Critter.Contract.Users(&_Critter.CallOpts, arg0)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) payable returns()
func (_Critter *CritterTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Critter.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) payable returns()
func (_Critter *CritterSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Critter.Contract.Approve(&_Critter.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) payable returns()
func (_Critter *CritterTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Critter.Contract.Approve(&_Critter.TransactOpts, to, tokenId)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x298daf5b.
//
// Solidity: function createAccount(string username) returns()
func (_Critter *CritterTransactor) CreateAccount(opts *bind.TransactOpts, username string) (*types.Transaction, error) {
	return _Critter.contract.Transact(opts, "createAccount", username)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x298daf5b.
//
// Solidity: function createAccount(string username) returns()
func (_Critter *CritterSession) CreateAccount(username string) (*types.Transaction, error) {
	return _Critter.Contract.CreateAccount(&_Critter.TransactOpts, username)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x298daf5b.
//
// Solidity: function createAccount(string username) returns()
func (_Critter *CritterTransactorSession) CreateAccount(username string) (*types.Transaction, error) {
	return _Critter.Contract.CreateAccount(&_Critter.TransactOpts, username)
}

// CreateSqueak is a paid mutator transaction binding the contract method 0xf3357b50.
//
// Solidity: function createSqueak(string content) returns()
func (_Critter *CritterTransactor) CreateSqueak(opts *bind.TransactOpts, content string) (*types.Transaction, error) {
	return _Critter.contract.Transact(opts, "createSqueak", content)
}

// CreateSqueak is a paid mutator transaction binding the contract method 0xf3357b50.
//
// Solidity: function createSqueak(string content) returns()
func (_Critter *CritterSession) CreateSqueak(content string) (*types.Transaction, error) {
	return _Critter.Contract.CreateSqueak(&_Critter.TransactOpts, content)
}

// CreateSqueak is a paid mutator transaction binding the contract method 0xf3357b50.
//
// Solidity: function createSqueak(string content) returns()
func (_Critter *CritterTransactorSession) CreateSqueak(content string) (*types.Transaction, error) {
	return _Critter.Contract.CreateSqueak(&_Critter.TransactOpts, content)
}

// DeleteSqueak is a paid mutator transaction binding the contract method 0xa015c9d7.
//
// Solidity: function deleteSqueak(uint256 tokenId) payable returns()
func (_Critter *CritterTransactor) DeleteSqueak(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Critter.contract.Transact(opts, "deleteSqueak", tokenId)
}

// DeleteSqueak is a paid mutator transaction binding the contract method 0xa015c9d7.
//
// Solidity: function deleteSqueak(uint256 tokenId) payable returns()
func (_Critter *CritterSession) DeleteSqueak(tokenId *big.Int) (*types.Transaction, error) {
	return _Critter.Contract.DeleteSqueak(&_Critter.TransactOpts, tokenId)
}

// DeleteSqueak is a paid mutator transaction binding the contract method 0xa015c9d7.
//
// Solidity: function deleteSqueak(uint256 tokenId) payable returns()
func (_Critter *CritterTransactorSession) DeleteSqueak(tokenId *big.Int) (*types.Transaction, error) {
	return _Critter.Contract.DeleteSqueak(&_Critter.TransactOpts, tokenId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Critter *CritterTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Critter.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Critter *CritterSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Critter.Contract.GrantRole(&_Critter.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Critter *CritterTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Critter.Contract.GrantRole(&_Critter.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xea5cfa44.
//
// Solidity: function initialize(string name, string symbol, string baseURI, uint256 platformFee, uint256 platformTakeRate, uint256 poolPayoutThreshold, uint256 viralityThreshold, uint256 viralityBonus, uint256 maxLevel) returns()
func (_Critter *CritterTransactor) Initialize(opts *bind.TransactOpts, name string, symbol string, baseURI string, platformFee *big.Int, platformTakeRate *big.Int, poolPayoutThreshold *big.Int, viralityThreshold *big.Int, viralityBonus *big.Int, maxLevel *big.Int) (*types.Transaction, error) {
	return _Critter.contract.Transact(opts, "initialize", name, symbol, baseURI, platformFee, platformTakeRate, poolPayoutThreshold, viralityThreshold, viralityBonus, maxLevel)
}

// Initialize is a paid mutator transaction binding the contract method 0xea5cfa44.
//
// Solidity: function initialize(string name, string symbol, string baseURI, uint256 platformFee, uint256 platformTakeRate, uint256 poolPayoutThreshold, uint256 viralityThreshold, uint256 viralityBonus, uint256 maxLevel) returns()
func (_Critter *CritterSession) Initialize(name string, symbol string, baseURI string, platformFee *big.Int, platformTakeRate *big.Int, poolPayoutThreshold *big.Int, viralityThreshold *big.Int, viralityBonus *big.Int, maxLevel *big.Int) (*types.Transaction, error) {
	return _Critter.Contract.Initialize(&_Critter.TransactOpts, name, symbol, baseURI, platformFee, platformTakeRate, poolPayoutThreshold, viralityThreshold, viralityBonus, maxLevel)
}

// Initialize is a paid mutator transaction binding the contract method 0xea5cfa44.
//
// Solidity: function initialize(string name, string symbol, string baseURI, uint256 platformFee, uint256 platformTakeRate, uint256 poolPayoutThreshold, uint256 viralityThreshold, uint256 viralityBonus, uint256 maxLevel) returns()
func (_Critter *CritterTransactorSession) Initialize(name string, symbol string, baseURI string, platformFee *big.Int, platformTakeRate *big.Int, poolPayoutThreshold *big.Int, viralityThreshold *big.Int, viralityBonus *big.Int, maxLevel *big.Int) (*types.Transaction, error) {
	return _Critter.Contract.Initialize(&_Critter.TransactOpts, name, symbol, baseURI, platformFee, platformTakeRate, poolPayoutThreshold, viralityThreshold, viralityBonus, maxLevel)
}

// Interact is a paid mutator transaction binding the contract method 0x701d8ef8.
//
// Solidity: function interact(uint256 tokenId, uint8 interaction) payable returns()
func (_Critter *CritterTransactor) Interact(opts *bind.TransactOpts, tokenId *big.Int, interaction uint8) (*types.Transaction, error) {
	return _Critter.contract.Transact(opts, "interact", tokenId, interaction)
}

// Interact is a paid mutator transaction binding the contract method 0x701d8ef8.
//
// Solidity: function interact(uint256 tokenId, uint8 interaction) payable returns()
func (_Critter *CritterSession) Interact(tokenId *big.Int, interaction uint8) (*types.Transaction, error) {
	return _Critter.Contract.Interact(&_Critter.TransactOpts, tokenId, interaction)
}

// Interact is a paid mutator transaction binding the contract method 0x701d8ef8.
//
// Solidity: function interact(uint256 tokenId, uint8 interaction) payable returns()
func (_Critter *CritterTransactorSession) Interact(tokenId *big.Int, interaction uint8) (*types.Transaction, error) {
	return _Critter.Contract.Interact(&_Critter.TransactOpts, tokenId, interaction)
}

// LeavePool is a paid mutator transaction binding the contract method 0x9ca53f82.
//
// Solidity: function leavePool(uint256 tokenId) returns()
func (_Critter *CritterTransactor) LeavePool(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Critter.contract.Transact(opts, "leavePool", tokenId)
}

// LeavePool is a paid mutator transaction binding the contract method 0x9ca53f82.
//
// Solidity: function leavePool(uint256 tokenId) returns()
func (_Critter *CritterSession) LeavePool(tokenId *big.Int) (*types.Transaction, error) {
	return _Critter.Contract.LeavePool(&_Critter.TransactOpts, tokenId)
}

// LeavePool is a paid mutator transaction binding the contract method 0x9ca53f82.
//
// Solidity: function leavePool(uint256 tokenId) returns()
func (_Critter *CritterTransactorSession) LeavePool(tokenId *big.Int) (*types.Transaction, error) {
	return _Critter.Contract.LeavePool(&_Critter.TransactOpts, tokenId)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Critter *CritterTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Critter.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Critter *CritterSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Critter.Contract.RenounceRole(&_Critter.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Critter *CritterTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Critter.Contract.RenounceRole(&_Critter.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Critter *CritterTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Critter.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Critter *CritterSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Critter.Contract.RevokeRole(&_Critter.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Critter *CritterTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Critter.Contract.RevokeRole(&_Critter.TransactOpts, role, account)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) payable returns()
func (_Critter *CritterTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Critter.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) payable returns()
func (_Critter *CritterSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Critter.Contract.SafeTransferFrom(&_Critter.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) payable returns()
func (_Critter *CritterTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Critter.Contract.SafeTransferFrom(&_Critter.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) payable returns()
func (_Critter *CritterTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Critter.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) payable returns()
func (_Critter *CritterSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Critter.Contract.SafeTransferFrom0(&_Critter.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) payable returns()
func (_Critter *CritterTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Critter.Contract.SafeTransferFrom0(&_Critter.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Critter *CritterTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Critter.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Critter *CritterSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Critter.Contract.SetApprovalForAll(&_Critter.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Critter *CritterTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Critter.Contract.SetApprovalForAll(&_Critter.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) payable returns()
func (_Critter *CritterTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Critter.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) payable returns()
func (_Critter *CritterSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Critter.Contract.TransferFrom(&_Critter.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) payable returns()
func (_Critter *CritterTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Critter.Contract.TransferFrom(&_Critter.TransactOpts, from, to, tokenId)
}

// UpdateConfiguration is a paid mutator transaction binding the contract method 0x3f57d800.
//
// Solidity: function updateConfiguration(uint8 configuration, uint256 amount) returns()
func (_Critter *CritterTransactor) UpdateConfiguration(opts *bind.TransactOpts, configuration uint8, amount *big.Int) (*types.Transaction, error) {
	return _Critter.contract.Transact(opts, "updateConfiguration", configuration, amount)
}

// UpdateConfiguration is a paid mutator transaction binding the contract method 0x3f57d800.
//
// Solidity: function updateConfiguration(uint8 configuration, uint256 amount) returns()
func (_Critter *CritterSession) UpdateConfiguration(configuration uint8, amount *big.Int) (*types.Transaction, error) {
	return _Critter.Contract.UpdateConfiguration(&_Critter.TransactOpts, configuration, amount)
}

// UpdateConfiguration is a paid mutator transaction binding the contract method 0x3f57d800.
//
// Solidity: function updateConfiguration(uint8 configuration, uint256 amount) returns()
func (_Critter *CritterTransactorSession) UpdateConfiguration(configuration uint8, amount *big.Int) (*types.Transaction, error) {
	return _Critter.Contract.UpdateConfiguration(&_Critter.TransactOpts, configuration, amount)
}

// UpdateInteractionFee is a paid mutator transaction binding the contract method 0x52347773.
//
// Solidity: function updateInteractionFee(uint8 interaction, uint256 amount) returns()
func (_Critter *CritterTransactor) UpdateInteractionFee(opts *bind.TransactOpts, interaction uint8, amount *big.Int) (*types.Transaction, error) {
	return _Critter.contract.Transact(opts, "updateInteractionFee", interaction, amount)
}

// UpdateInteractionFee is a paid mutator transaction binding the contract method 0x52347773.
//
// Solidity: function updateInteractionFee(uint8 interaction, uint256 amount) returns()
func (_Critter *CritterSession) UpdateInteractionFee(interaction uint8, amount *big.Int) (*types.Transaction, error) {
	return _Critter.Contract.UpdateInteractionFee(&_Critter.TransactOpts, interaction, amount)
}

// UpdateInteractionFee is a paid mutator transaction binding the contract method 0x52347773.
//
// Solidity: function updateInteractionFee(uint8 interaction, uint256 amount) returns()
func (_Critter *CritterTransactorSession) UpdateInteractionFee(interaction uint8, amount *big.Int) (*types.Transaction, error) {
	return _Critter.Contract.UpdateInteractionFee(&_Critter.TransactOpts, interaction, amount)
}

// UpdateRelationship is a paid mutator transaction binding the contract method 0xdf2ae7c7.
//
// Solidity: function updateRelationship(address account, uint8 action) returns()
func (_Critter *CritterTransactor) UpdateRelationship(opts *bind.TransactOpts, account common.Address, action uint8) (*types.Transaction, error) {
	return _Critter.contract.Transact(opts, "updateRelationship", account, action)
}

// UpdateRelationship is a paid mutator transaction binding the contract method 0xdf2ae7c7.
//
// Solidity: function updateRelationship(address account, uint8 action) returns()
func (_Critter *CritterSession) UpdateRelationship(account common.Address, action uint8) (*types.Transaction, error) {
	return _Critter.Contract.UpdateRelationship(&_Critter.TransactOpts, account, action)
}

// UpdateRelationship is a paid mutator transaction binding the contract method 0xdf2ae7c7.
//
// Solidity: function updateRelationship(address account, uint8 action) returns()
func (_Critter *CritterTransactorSession) UpdateRelationship(account common.Address, action uint8) (*types.Transaction, error) {
	return _Critter.Contract.UpdateRelationship(&_Critter.TransactOpts, account, action)
}

// UpdateStatus is a paid mutator transaction binding the contract method 0x44c5bbf8.
//
// Solidity: function updateStatus(address account, uint8 status) returns()
func (_Critter *CritterTransactor) UpdateStatus(opts *bind.TransactOpts, account common.Address, status uint8) (*types.Transaction, error) {
	return _Critter.contract.Transact(opts, "updateStatus", account, status)
}

// UpdateStatus is a paid mutator transaction binding the contract method 0x44c5bbf8.
//
// Solidity: function updateStatus(address account, uint8 status) returns()
func (_Critter *CritterSession) UpdateStatus(account common.Address, status uint8) (*types.Transaction, error) {
	return _Critter.Contract.UpdateStatus(&_Critter.TransactOpts, account, status)
}

// UpdateStatus is a paid mutator transaction binding the contract method 0x44c5bbf8.
//
// Solidity: function updateStatus(address account, uint8 status) returns()
func (_Critter *CritterTransactorSession) UpdateStatus(account common.Address, status uint8) (*types.Transaction, error) {
	return _Critter.Contract.UpdateStatus(&_Critter.TransactOpts, account, status)
}

// UpdateUsername is a paid mutator transaction binding the contract method 0xc96cea70.
//
// Solidity: function updateUsername(string newUsername) returns()
func (_Critter *CritterTransactor) UpdateUsername(opts *bind.TransactOpts, newUsername string) (*types.Transaction, error) {
	return _Critter.contract.Transact(opts, "updateUsername", newUsername)
}

// UpdateUsername is a paid mutator transaction binding the contract method 0xc96cea70.
//
// Solidity: function updateUsername(string newUsername) returns()
func (_Critter *CritterSession) UpdateUsername(newUsername string) (*types.Transaction, error) {
	return _Critter.Contract.UpdateUsername(&_Critter.TransactOpts, newUsername)
}

// UpdateUsername is a paid mutator transaction binding the contract method 0xc96cea70.
//
// Solidity: function updateUsername(string newUsername) returns()
func (_Critter *CritterTransactorSession) UpdateUsername(newUsername string) (*types.Transaction, error) {
	return _Critter.Contract.UpdateUsername(&_Critter.TransactOpts, newUsername)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Critter *CritterTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Critter.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Critter *CritterSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Critter.Contract.UpgradeTo(&_Critter.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Critter *CritterTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Critter.Contract.UpgradeTo(&_Critter.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Critter *CritterTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Critter.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Critter *CritterSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Critter.Contract.UpgradeToAndCall(&_Critter.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Critter *CritterTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Critter.Contract.UpgradeToAndCall(&_Critter.TransactOpts, newImplementation, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address to, uint256 amount) payable returns()
func (_Critter *CritterTransactor) Withdraw(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Critter.contract.Transact(opts, "withdraw", to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address to, uint256 amount) payable returns()
func (_Critter *CritterSession) Withdraw(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Critter.Contract.Withdraw(&_Critter.TransactOpts, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address to, uint256 amount) payable returns()
func (_Critter *CritterTransactorSession) Withdraw(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Critter.Contract.Withdraw(&_Critter.TransactOpts, to, amount)
}

// CritterAccountCreatedIterator is returned from FilterAccountCreated and is used to iterate over the raw logs and unpacked data for AccountCreated events raised by the Critter contract.
type CritterAccountCreatedIterator struct {
	Event *CritterAccountCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CritterAccountCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CritterAccountCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CritterAccountCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CritterAccountCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CritterAccountCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CritterAccountCreated represents a AccountCreated event raised by the Critter contract.
type CritterAccountCreated struct {
	Account  common.Address
	Username [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAccountCreated is a free log retrieval operation binding the contract event 0x8fe66a5d954d6d3e0306797e31e226812a9916895165c96c367ef52807631951.
//
// Solidity: event AccountCreated(address indexed account, bytes32 indexed username)
func (_Critter *CritterFilterer) FilterAccountCreated(opts *bind.FilterOpts, account []common.Address, username [][32]byte) (*CritterAccountCreatedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var usernameRule []interface{}
	for _, usernameItem := range username {
		usernameRule = append(usernameRule, usernameItem)
	}

	logs, sub, err := _Critter.contract.FilterLogs(opts, "AccountCreated", accountRule, usernameRule)
	if err != nil {
		return nil, err
	}
	return &CritterAccountCreatedIterator{contract: _Critter.contract, event: "AccountCreated", logs: logs, sub: sub}, nil
}

// WatchAccountCreated is a free log subscription operation binding the contract event 0x8fe66a5d954d6d3e0306797e31e226812a9916895165c96c367ef52807631951.
//
// Solidity: event AccountCreated(address indexed account, bytes32 indexed username)
func (_Critter *CritterFilterer) WatchAccountCreated(opts *bind.WatchOpts, sink chan<- *CritterAccountCreated, account []common.Address, username [][32]byte) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var usernameRule []interface{}
	for _, usernameItem := range username {
		usernameRule = append(usernameRule, usernameItem)
	}

	logs, sub, err := _Critter.contract.WatchLogs(opts, "AccountCreated", accountRule, usernameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CritterAccountCreated)
				if err := _Critter.contract.UnpackLog(event, "AccountCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAccountCreated is a log parse operation binding the contract event 0x8fe66a5d954d6d3e0306797e31e226812a9916895165c96c367ef52807631951.
//
// Solidity: event AccountCreated(address indexed account, bytes32 indexed username)
func (_Critter *CritterFilterer) ParseAccountCreated(log types.Log) (*CritterAccountCreated, error) {
	event := new(CritterAccountCreated)
	if err := _Critter.contract.UnpackLog(event, "AccountCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CritterAccountUsernameUpdatedIterator is returned from FilterAccountUsernameUpdated and is used to iterate over the raw logs and unpacked data for AccountUsernameUpdated events raised by the Critter contract.
type CritterAccountUsernameUpdatedIterator struct {
	Event *CritterAccountUsernameUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CritterAccountUsernameUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CritterAccountUsernameUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CritterAccountUsernameUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CritterAccountUsernameUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CritterAccountUsernameUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CritterAccountUsernameUpdated represents a AccountUsernameUpdated event raised by the Critter contract.
type CritterAccountUsernameUpdated struct {
	Account     common.Address
	NewUsername string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAccountUsernameUpdated is a free log retrieval operation binding the contract event 0xe15cfc00b3a9258ddde2ebf35aa9be829fd79430e061370458c2bd656cc375ac.
//
// Solidity: event AccountUsernameUpdated(address account, string newUsername)
func (_Critter *CritterFilterer) FilterAccountUsernameUpdated(opts *bind.FilterOpts) (*CritterAccountUsernameUpdatedIterator, error) {

	logs, sub, err := _Critter.contract.FilterLogs(opts, "AccountUsernameUpdated")
	if err != nil {
		return nil, err
	}
	return &CritterAccountUsernameUpdatedIterator{contract: _Critter.contract, event: "AccountUsernameUpdated", logs: logs, sub: sub}, nil
}

// WatchAccountUsernameUpdated is a free log subscription operation binding the contract event 0xe15cfc00b3a9258ddde2ebf35aa9be829fd79430e061370458c2bd656cc375ac.
//
// Solidity: event AccountUsernameUpdated(address account, string newUsername)
func (_Critter *CritterFilterer) WatchAccountUsernameUpdated(opts *bind.WatchOpts, sink chan<- *CritterAccountUsernameUpdated) (event.Subscription, error) {

	logs, sub, err := _Critter.contract.WatchLogs(opts, "AccountUsernameUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CritterAccountUsernameUpdated)
				if err := _Critter.contract.UnpackLog(event, "AccountUsernameUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAccountUsernameUpdated is a log parse operation binding the contract event 0xe15cfc00b3a9258ddde2ebf35aa9be829fd79430e061370458c2bd656cc375ac.
//
// Solidity: event AccountUsernameUpdated(address account, string newUsername)
func (_Critter *CritterFilterer) ParseAccountUsernameUpdated(log types.Log) (*CritterAccountUsernameUpdated, error) {
	event := new(CritterAccountUsernameUpdated)
	if err := _Critter.contract.UnpackLog(event, "AccountUsernameUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CritterAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Critter contract.
type CritterAdminChangedIterator struct {
	Event *CritterAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CritterAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CritterAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CritterAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CritterAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CritterAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CritterAdminChanged represents a AdminChanged event raised by the Critter contract.
type CritterAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Critter *CritterFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*CritterAdminChangedIterator, error) {

	logs, sub, err := _Critter.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &CritterAdminChangedIterator{contract: _Critter.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Critter *CritterFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *CritterAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Critter.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CritterAdminChanged)
				if err := _Critter.contract.UnpackLog(event, "AdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Critter *CritterFilterer) ParseAdminChanged(log types.Log) (*CritterAdminChanged, error) {
	event := new(CritterAdminChanged)
	if err := _Critter.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CritterApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Critter contract.
type CritterApprovalIterator struct {
	Event *CritterApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CritterApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CritterApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CritterApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CritterApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CritterApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CritterApproval represents a Approval event raised by the Critter contract.
type CritterApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Critter *CritterFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*CritterApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Critter.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CritterApprovalIterator{contract: _Critter.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Critter *CritterFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CritterApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Critter.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CritterApproval)
				if err := _Critter.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Critter *CritterFilterer) ParseApproval(log types.Log) (*CritterApproval, error) {
	event := new(CritterApproval)
	if err := _Critter.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CritterApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Critter contract.
type CritterApprovalForAllIterator struct {
	Event *CritterApprovalForAll // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CritterApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CritterApprovalForAll)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CritterApprovalForAll)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CritterApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CritterApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CritterApprovalForAll represents a ApprovalForAll event raised by the Critter contract.
type CritterApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Critter *CritterFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*CritterApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Critter.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &CritterApprovalForAllIterator{contract: _Critter.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Critter *CritterFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *CritterApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Critter.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CritterApprovalForAll)
				if err := _Critter.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Critter *CritterFilterer) ParseApprovalForAll(log types.Log) (*CritterApprovalForAll, error) {
	event := new(CritterApprovalForAll)
	if err := _Critter.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CritterBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Critter contract.
type CritterBeaconUpgradedIterator struct {
	Event *CritterBeaconUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CritterBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CritterBeaconUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CritterBeaconUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CritterBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CritterBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CritterBeaconUpgraded represents a BeaconUpgraded event raised by the Critter contract.
type CritterBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Critter *CritterFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*CritterBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Critter.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &CritterBeaconUpgradedIterator{contract: _Critter.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Critter *CritterFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *CritterBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Critter.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CritterBeaconUpgraded)
				if err := _Critter.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Critter *CritterFilterer) ParseBeaconUpgraded(log types.Log) (*CritterBeaconUpgraded, error) {
	event := new(CritterBeaconUpgraded)
	if err := _Critter.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CritterConsecutiveTransferIterator is returned from FilterConsecutiveTransfer and is used to iterate over the raw logs and unpacked data for ConsecutiveTransfer events raised by the Critter contract.
type CritterConsecutiveTransferIterator struct {
	Event *CritterConsecutiveTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CritterConsecutiveTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CritterConsecutiveTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CritterConsecutiveTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CritterConsecutiveTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CritterConsecutiveTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CritterConsecutiveTransfer represents a ConsecutiveTransfer event raised by the Critter contract.
type CritterConsecutiveTransfer struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	From        common.Address
	To          common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterConsecutiveTransfer is a free log retrieval operation binding the contract event 0xdeaa91b6123d068f5821d0fb0678463d1a8a6079fe8af5de3ce5e896dcf9133d.
//
// Solidity: event ConsecutiveTransfer(uint256 indexed fromTokenId, uint256 toTokenId, address indexed from, address indexed to)
func (_Critter *CritterFilterer) FilterConsecutiveTransfer(opts *bind.FilterOpts, fromTokenId []*big.Int, from []common.Address, to []common.Address) (*CritterConsecutiveTransferIterator, error) {

	var fromTokenIdRule []interface{}
	for _, fromTokenIdItem := range fromTokenId {
		fromTokenIdRule = append(fromTokenIdRule, fromTokenIdItem)
	}

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Critter.contract.FilterLogs(opts, "ConsecutiveTransfer", fromTokenIdRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CritterConsecutiveTransferIterator{contract: _Critter.contract, event: "ConsecutiveTransfer", logs: logs, sub: sub}, nil
}

// WatchConsecutiveTransfer is a free log subscription operation binding the contract event 0xdeaa91b6123d068f5821d0fb0678463d1a8a6079fe8af5de3ce5e896dcf9133d.
//
// Solidity: event ConsecutiveTransfer(uint256 indexed fromTokenId, uint256 toTokenId, address indexed from, address indexed to)
func (_Critter *CritterFilterer) WatchConsecutiveTransfer(opts *bind.WatchOpts, sink chan<- *CritterConsecutiveTransfer, fromTokenId []*big.Int, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromTokenIdRule []interface{}
	for _, fromTokenIdItem := range fromTokenId {
		fromTokenIdRule = append(fromTokenIdRule, fromTokenIdItem)
	}

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Critter.contract.WatchLogs(opts, "ConsecutiveTransfer", fromTokenIdRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CritterConsecutiveTransfer)
				if err := _Critter.contract.UnpackLog(event, "ConsecutiveTransfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseConsecutiveTransfer is a log parse operation binding the contract event 0xdeaa91b6123d068f5821d0fb0678463d1a8a6079fe8af5de3ce5e896dcf9133d.
//
// Solidity: event ConsecutiveTransfer(uint256 indexed fromTokenId, uint256 toTokenId, address indexed from, address indexed to)
func (_Critter *CritterFilterer) ParseConsecutiveTransfer(log types.Log) (*CritterConsecutiveTransfer, error) {
	event := new(CritterConsecutiveTransfer)
	if err := _Critter.contract.UnpackLog(event, "ConsecutiveTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CritterFundsAddedToPoolIterator is returned from FilterFundsAddedToPool and is used to iterate over the raw logs and unpacked data for FundsAddedToPool events raised by the Critter contract.
type CritterFundsAddedToPoolIterator struct {
	Event *CritterFundsAddedToPool // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CritterFundsAddedToPoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CritterFundsAddedToPool)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CritterFundsAddedToPool)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CritterFundsAddedToPoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CritterFundsAddedToPoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CritterFundsAddedToPool represents a FundsAddedToPool event raised by the Critter contract.
type CritterFundsAddedToPool struct {
	TokenId *big.Int
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFundsAddedToPool is a free log retrieval operation binding the contract event 0xe2b757cdd8b279ca2f3316f88e8e3b959d9bad09f9369132a3475d3ef8a4f366.
//
// Solidity: event FundsAddedToPool(uint256 tokenId, uint256 amount)
func (_Critter *CritterFilterer) FilterFundsAddedToPool(opts *bind.FilterOpts) (*CritterFundsAddedToPoolIterator, error) {

	logs, sub, err := _Critter.contract.FilterLogs(opts, "FundsAddedToPool")
	if err != nil {
		return nil, err
	}
	return &CritterFundsAddedToPoolIterator{contract: _Critter.contract, event: "FundsAddedToPool", logs: logs, sub: sub}, nil
}

// WatchFundsAddedToPool is a free log subscription operation binding the contract event 0xe2b757cdd8b279ca2f3316f88e8e3b959d9bad09f9369132a3475d3ef8a4f366.
//
// Solidity: event FundsAddedToPool(uint256 tokenId, uint256 amount)
func (_Critter *CritterFilterer) WatchFundsAddedToPool(opts *bind.WatchOpts, sink chan<- *CritterFundsAddedToPool) (event.Subscription, error) {

	logs, sub, err := _Critter.contract.WatchLogs(opts, "FundsAddedToPool")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CritterFundsAddedToPool)
				if err := _Critter.contract.UnpackLog(event, "FundsAddedToPool", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFundsAddedToPool is a log parse operation binding the contract event 0xe2b757cdd8b279ca2f3316f88e8e3b959d9bad09f9369132a3475d3ef8a4f366.
//
// Solidity: event FundsAddedToPool(uint256 tokenId, uint256 amount)
func (_Critter *CritterFilterer) ParseFundsAddedToPool(log types.Log) (*CritterFundsAddedToPool, error) {
	event := new(CritterFundsAddedToPool)
	if err := _Critter.contract.UnpackLog(event, "FundsAddedToPool", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CritterFundsDepositedIterator is returned from FilterFundsDeposited and is used to iterate over the raw logs and unpacked data for FundsDeposited events raised by the Critter contract.
type CritterFundsDepositedIterator struct {
	Event *CritterFundsDeposited // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CritterFundsDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CritterFundsDeposited)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CritterFundsDeposited)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CritterFundsDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CritterFundsDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CritterFundsDeposited represents a FundsDeposited event raised by the Critter contract.
type CritterFundsDeposited struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFundsDeposited is a free log retrieval operation binding the contract event 0x7cb5d45c1071f219bd2da1ea101758715bd52f23f2a707f2d9a0387dea18d906.
//
// Solidity: event FundsDeposited(uint256 amount)
func (_Critter *CritterFilterer) FilterFundsDeposited(opts *bind.FilterOpts) (*CritterFundsDepositedIterator, error) {

	logs, sub, err := _Critter.contract.FilterLogs(opts, "FundsDeposited")
	if err != nil {
		return nil, err
	}
	return &CritterFundsDepositedIterator{contract: _Critter.contract, event: "FundsDeposited", logs: logs, sub: sub}, nil
}

// WatchFundsDeposited is a free log subscription operation binding the contract event 0x7cb5d45c1071f219bd2da1ea101758715bd52f23f2a707f2d9a0387dea18d906.
//
// Solidity: event FundsDeposited(uint256 amount)
func (_Critter *CritterFilterer) WatchFundsDeposited(opts *bind.WatchOpts, sink chan<- *CritterFundsDeposited) (event.Subscription, error) {

	logs, sub, err := _Critter.contract.WatchLogs(opts, "FundsDeposited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CritterFundsDeposited)
				if err := _Critter.contract.UnpackLog(event, "FundsDeposited", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFundsDeposited is a log parse operation binding the contract event 0x7cb5d45c1071f219bd2da1ea101758715bd52f23f2a707f2d9a0387dea18d906.
//
// Solidity: event FundsDeposited(uint256 amount)
func (_Critter *CritterFilterer) ParseFundsDeposited(log types.Log) (*CritterFundsDeposited, error) {
	event := new(CritterFundsDeposited)
	if err := _Critter.contract.UnpackLog(event, "FundsDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CritterFundsTransferredIterator is returned from FilterFundsTransferred and is used to iterate over the raw logs and unpacked data for FundsTransferred events raised by the Critter contract.
type CritterFundsTransferredIterator struct {
	Event *CritterFundsTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CritterFundsTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CritterFundsTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CritterFundsTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CritterFundsTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CritterFundsTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CritterFundsTransferred represents a FundsTransferred event raised by the Critter contract.
type CritterFundsTransferred struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFundsTransferred is a free log retrieval operation binding the contract event 0x8c9a4f13b67cb64d7c6aa1ae0c9bf07694af521a28b93e7060020810ab4bc59f.
//
// Solidity: event FundsTransferred(address indexed to, uint256 amount)
func (_Critter *CritterFilterer) FilterFundsTransferred(opts *bind.FilterOpts, to []common.Address) (*CritterFundsTransferredIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Critter.contract.FilterLogs(opts, "FundsTransferred", toRule)
	if err != nil {
		return nil, err
	}
	return &CritterFundsTransferredIterator{contract: _Critter.contract, event: "FundsTransferred", logs: logs, sub: sub}, nil
}

// WatchFundsTransferred is a free log subscription operation binding the contract event 0x8c9a4f13b67cb64d7c6aa1ae0c9bf07694af521a28b93e7060020810ab4bc59f.
//
// Solidity: event FundsTransferred(address indexed to, uint256 amount)
func (_Critter *CritterFilterer) WatchFundsTransferred(opts *bind.WatchOpts, sink chan<- *CritterFundsTransferred, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Critter.contract.WatchLogs(opts, "FundsTransferred", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CritterFundsTransferred)
				if err := _Critter.contract.UnpackLog(event, "FundsTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFundsTransferred is a log parse operation binding the contract event 0x8c9a4f13b67cb64d7c6aa1ae0c9bf07694af521a28b93e7060020810ab4bc59f.
//
// Solidity: event FundsTransferred(address indexed to, uint256 amount)
func (_Critter *CritterFilterer) ParseFundsTransferred(log types.Log) (*CritterFundsTransferred, error) {
	event := new(CritterFundsTransferred)
	if err := _Critter.contract.UnpackLog(event, "FundsTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CritterFundsWithdrawnIterator is returned from FilterFundsWithdrawn and is used to iterate over the raw logs and unpacked data for FundsWithdrawn events raised by the Critter contract.
type CritterFundsWithdrawnIterator struct {
	Event *CritterFundsWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CritterFundsWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CritterFundsWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CritterFundsWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CritterFundsWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CritterFundsWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CritterFundsWithdrawn represents a FundsWithdrawn event raised by the Critter contract.
type CritterFundsWithdrawn struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFundsWithdrawn is a free log retrieval operation binding the contract event 0xeaff4b37086828766ad3268786972c0cd24259d4c87a80f9d3963a3c3d999b0d.
//
// Solidity: event FundsWithdrawn(address indexed to, uint256 amount)
func (_Critter *CritterFilterer) FilterFundsWithdrawn(opts *bind.FilterOpts, to []common.Address) (*CritterFundsWithdrawnIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Critter.contract.FilterLogs(opts, "FundsWithdrawn", toRule)
	if err != nil {
		return nil, err
	}
	return &CritterFundsWithdrawnIterator{contract: _Critter.contract, event: "FundsWithdrawn", logs: logs, sub: sub}, nil
}

// WatchFundsWithdrawn is a free log subscription operation binding the contract event 0xeaff4b37086828766ad3268786972c0cd24259d4c87a80f9d3963a3c3d999b0d.
//
// Solidity: event FundsWithdrawn(address indexed to, uint256 amount)
func (_Critter *CritterFilterer) WatchFundsWithdrawn(opts *bind.WatchOpts, sink chan<- *CritterFundsWithdrawn, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Critter.contract.WatchLogs(opts, "FundsWithdrawn", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CritterFundsWithdrawn)
				if err := _Critter.contract.UnpackLog(event, "FundsWithdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFundsWithdrawn is a log parse operation binding the contract event 0xeaff4b37086828766ad3268786972c0cd24259d4c87a80f9d3963a3c3d999b0d.
//
// Solidity: event FundsWithdrawn(address indexed to, uint256 amount)
func (_Critter *CritterFilterer) ParseFundsWithdrawn(log types.Log) (*CritterFundsWithdrawn, error) {
	event := new(CritterFundsWithdrawn)
	if err := _Critter.contract.UnpackLog(event, "FundsWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CritterInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Critter contract.
type CritterInitializedIterator struct {
	Event *CritterInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CritterInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CritterInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CritterInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CritterInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CritterInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CritterInitialized represents a Initialized event raised by the Critter contract.
type CritterInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Critter *CritterFilterer) FilterInitialized(opts *bind.FilterOpts) (*CritterInitializedIterator, error) {

	logs, sub, err := _Critter.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &CritterInitializedIterator{contract: _Critter.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Critter *CritterFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *CritterInitialized) (event.Subscription, error) {

	logs, sub, err := _Critter.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CritterInitialized)
				if err := _Critter.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Critter *CritterFilterer) ParseInitialized(log types.Log) (*CritterInitialized, error) {
	event := new(CritterInitialized)
	if err := _Critter.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CritterInteractionFeeUpdatedIterator is returned from FilterInteractionFeeUpdated and is used to iterate over the raw logs and unpacked data for InteractionFeeUpdated events raised by the Critter contract.
type CritterInteractionFeeUpdatedIterator struct {
	Event *CritterInteractionFeeUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CritterInteractionFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CritterInteractionFeeUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CritterInteractionFeeUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CritterInteractionFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CritterInteractionFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CritterInteractionFeeUpdated represents a InteractionFeeUpdated event raised by the Critter contract.
type CritterInteractionFeeUpdated struct {
	Interaction uint8
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterInteractionFeeUpdated is a free log retrieval operation binding the contract event 0x2367b007366025a5b7353c26b6e4e03d0a9dc2bb538423c7bb46c447fdcc9f94.
//
// Solidity: event InteractionFeeUpdated(uint8 interaction, uint256 amount)
func (_Critter *CritterFilterer) FilterInteractionFeeUpdated(opts *bind.FilterOpts) (*CritterInteractionFeeUpdatedIterator, error) {

	logs, sub, err := _Critter.contract.FilterLogs(opts, "InteractionFeeUpdated")
	if err != nil {
		return nil, err
	}
	return &CritterInteractionFeeUpdatedIterator{contract: _Critter.contract, event: "InteractionFeeUpdated", logs: logs, sub: sub}, nil
}

// WatchInteractionFeeUpdated is a free log subscription operation binding the contract event 0x2367b007366025a5b7353c26b6e4e03d0a9dc2bb538423c7bb46c447fdcc9f94.
//
// Solidity: event InteractionFeeUpdated(uint8 interaction, uint256 amount)
func (_Critter *CritterFilterer) WatchInteractionFeeUpdated(opts *bind.WatchOpts, sink chan<- *CritterInteractionFeeUpdated) (event.Subscription, error) {

	logs, sub, err := _Critter.contract.WatchLogs(opts, "InteractionFeeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CritterInteractionFeeUpdated)
				if err := _Critter.contract.UnpackLog(event, "InteractionFeeUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInteractionFeeUpdated is a log parse operation binding the contract event 0x2367b007366025a5b7353c26b6e4e03d0a9dc2bb538423c7bb46c447fdcc9f94.
//
// Solidity: event InteractionFeeUpdated(uint8 interaction, uint256 amount)
func (_Critter *CritterFilterer) ParseInteractionFeeUpdated(log types.Log) (*CritterInteractionFeeUpdated, error) {
	event := new(CritterInteractionFeeUpdated)
	if err := _Critter.contract.UnpackLog(event, "InteractionFeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CritterPoolPayoutIterator is returned from FilterPoolPayout and is used to iterate over the raw logs and unpacked data for PoolPayout events raised by the Critter contract.
type CritterPoolPayoutIterator struct {
	Event *CritterPoolPayout // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CritterPoolPayoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CritterPoolPayout)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CritterPoolPayout)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CritterPoolPayoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CritterPoolPayoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CritterPoolPayout represents a PoolPayout event raised by the Critter contract.
type CritterPoolPayout struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPoolPayout is a free log retrieval operation binding the contract event 0x535d43775b5470ea8720cb895efca28d2fa024f0bf05c5a9c212a1cd909f46d9.
//
// Solidity: event PoolPayout(uint256 tokenId)
func (_Critter *CritterFilterer) FilterPoolPayout(opts *bind.FilterOpts) (*CritterPoolPayoutIterator, error) {

	logs, sub, err := _Critter.contract.FilterLogs(opts, "PoolPayout")
	if err != nil {
		return nil, err
	}
	return &CritterPoolPayoutIterator{contract: _Critter.contract, event: "PoolPayout", logs: logs, sub: sub}, nil
}

// WatchPoolPayout is a free log subscription operation binding the contract event 0x535d43775b5470ea8720cb895efca28d2fa024f0bf05c5a9c212a1cd909f46d9.
//
// Solidity: event PoolPayout(uint256 tokenId)
func (_Critter *CritterFilterer) WatchPoolPayout(opts *bind.WatchOpts, sink chan<- *CritterPoolPayout) (event.Subscription, error) {

	logs, sub, err := _Critter.contract.WatchLogs(opts, "PoolPayout")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CritterPoolPayout)
				if err := _Critter.contract.UnpackLog(event, "PoolPayout", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePoolPayout is a log parse operation binding the contract event 0x535d43775b5470ea8720cb895efca28d2fa024f0bf05c5a9c212a1cd909f46d9.
//
// Solidity: event PoolPayout(uint256 tokenId)
func (_Critter *CritterFilterer) ParsePoolPayout(log types.Log) (*CritterPoolPayout, error) {
	event := new(CritterPoolPayout)
	if err := _Critter.contract.UnpackLog(event, "PoolPayout", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CritterRelationshipUpdatedIterator is returned from FilterRelationshipUpdated and is used to iterate over the raw logs and unpacked data for RelationshipUpdated events raised by the Critter contract.
type CritterRelationshipUpdatedIterator struct {
	Event *CritterRelationshipUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CritterRelationshipUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CritterRelationshipUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CritterRelationshipUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CritterRelationshipUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CritterRelationshipUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CritterRelationshipUpdated represents a RelationshipUpdated event raised by the Critter contract.
type CritterRelationshipUpdated struct {
	Sender   common.Address
	Relative common.Address
	Action   uint8
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRelationshipUpdated is a free log retrieval operation binding the contract event 0xe13c1bcdc36659a504f077c150e670a84d1a22ee9ee8e59f385249f9587a8c8d.
//
// Solidity: event RelationshipUpdated(address sender, address relative, uint8 action)
func (_Critter *CritterFilterer) FilterRelationshipUpdated(opts *bind.FilterOpts) (*CritterRelationshipUpdatedIterator, error) {

	logs, sub, err := _Critter.contract.FilterLogs(opts, "RelationshipUpdated")
	if err != nil {
		return nil, err
	}
	return &CritterRelationshipUpdatedIterator{contract: _Critter.contract, event: "RelationshipUpdated", logs: logs, sub: sub}, nil
}

// WatchRelationshipUpdated is a free log subscription operation binding the contract event 0xe13c1bcdc36659a504f077c150e670a84d1a22ee9ee8e59f385249f9587a8c8d.
//
// Solidity: event RelationshipUpdated(address sender, address relative, uint8 action)
func (_Critter *CritterFilterer) WatchRelationshipUpdated(opts *bind.WatchOpts, sink chan<- *CritterRelationshipUpdated) (event.Subscription, error) {

	logs, sub, err := _Critter.contract.WatchLogs(opts, "RelationshipUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CritterRelationshipUpdated)
				if err := _Critter.contract.UnpackLog(event, "RelationshipUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRelationshipUpdated is a log parse operation binding the contract event 0xe13c1bcdc36659a504f077c150e670a84d1a22ee9ee8e59f385249f9587a8c8d.
//
// Solidity: event RelationshipUpdated(address sender, address relative, uint8 action)
func (_Critter *CritterFilterer) ParseRelationshipUpdated(log types.Log) (*CritterRelationshipUpdated, error) {
	event := new(CritterRelationshipUpdated)
	if err := _Critter.contract.UnpackLog(event, "RelationshipUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CritterRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Critter contract.
type CritterRoleAdminChangedIterator struct {
	Event *CritterRoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CritterRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CritterRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CritterRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CritterRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CritterRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CritterRoleAdminChanged represents a RoleAdminChanged event raised by the Critter contract.
type CritterRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Critter *CritterFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*CritterRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Critter.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &CritterRoleAdminChangedIterator{contract: _Critter.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Critter *CritterFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *CritterRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Critter.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CritterRoleAdminChanged)
				if err := _Critter.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Critter *CritterFilterer) ParseRoleAdminChanged(log types.Log) (*CritterRoleAdminChanged, error) {
	event := new(CritterRoleAdminChanged)
	if err := _Critter.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CritterRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Critter contract.
type CritterRoleGrantedIterator struct {
	Event *CritterRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CritterRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CritterRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CritterRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CritterRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CritterRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CritterRoleGranted represents a RoleGranted event raised by the Critter contract.
type CritterRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Critter *CritterFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CritterRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Critter.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CritterRoleGrantedIterator{contract: _Critter.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Critter *CritterFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *CritterRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Critter.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CritterRoleGranted)
				if err := _Critter.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Critter *CritterFilterer) ParseRoleGranted(log types.Log) (*CritterRoleGranted, error) {
	event := new(CritterRoleGranted)
	if err := _Critter.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CritterRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Critter contract.
type CritterRoleRevokedIterator struct {
	Event *CritterRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CritterRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CritterRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CritterRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CritterRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CritterRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CritterRoleRevoked represents a RoleRevoked event raised by the Critter contract.
type CritterRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Critter *CritterFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CritterRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Critter.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CritterRoleRevokedIterator{contract: _Critter.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Critter *CritterFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *CritterRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Critter.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CritterRoleRevoked)
				if err := _Critter.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Critter *CritterFilterer) ParseRoleRevoked(log types.Log) (*CritterRoleRevoked, error) {
	event := new(CritterRoleRevoked)
	if err := _Critter.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CritterSqueakCreatedIterator is returned from FilterSqueakCreated and is used to iterate over the raw logs and unpacked data for SqueakCreated events raised by the Critter contract.
type CritterSqueakCreatedIterator struct {
	Event *CritterSqueakCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CritterSqueakCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CritterSqueakCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CritterSqueakCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CritterSqueakCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CritterSqueakCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CritterSqueakCreated represents a SqueakCreated event raised by the Critter contract.
type CritterSqueakCreated struct {
	Author      common.Address
	TokenId     *big.Int
	BlockNumber *big.Int
	Content     string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSqueakCreated is a free log retrieval operation binding the contract event 0x1127829e7034c69522f6c9a2361f7a0753366509033a2a37760a329f111187f3.
//
// Solidity: event SqueakCreated(address indexed author, uint256 tokenId, uint256 blockNumber, string content)
func (_Critter *CritterFilterer) FilterSqueakCreated(opts *bind.FilterOpts, author []common.Address) (*CritterSqueakCreatedIterator, error) {

	var authorRule []interface{}
	for _, authorItem := range author {
		authorRule = append(authorRule, authorItem)
	}

	logs, sub, err := _Critter.contract.FilterLogs(opts, "SqueakCreated", authorRule)
	if err != nil {
		return nil, err
	}
	return &CritterSqueakCreatedIterator{contract: _Critter.contract, event: "SqueakCreated", logs: logs, sub: sub}, nil
}

// WatchSqueakCreated is a free log subscription operation binding the contract event 0x1127829e7034c69522f6c9a2361f7a0753366509033a2a37760a329f111187f3.
//
// Solidity: event SqueakCreated(address indexed author, uint256 tokenId, uint256 blockNumber, string content)
func (_Critter *CritterFilterer) WatchSqueakCreated(opts *bind.WatchOpts, sink chan<- *CritterSqueakCreated, author []common.Address) (event.Subscription, error) {

	var authorRule []interface{}
	for _, authorItem := range author {
		authorRule = append(authorRule, authorItem)
	}

	logs, sub, err := _Critter.contract.WatchLogs(opts, "SqueakCreated", authorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CritterSqueakCreated)
				if err := _Critter.contract.UnpackLog(event, "SqueakCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSqueakCreated is a log parse operation binding the contract event 0x1127829e7034c69522f6c9a2361f7a0753366509033a2a37760a329f111187f3.
//
// Solidity: event SqueakCreated(address indexed author, uint256 tokenId, uint256 blockNumber, string content)
func (_Critter *CritterFilterer) ParseSqueakCreated(log types.Log) (*CritterSqueakCreated, error) {
	event := new(CritterSqueakCreated)
	if err := _Critter.contract.UnpackLog(event, "SqueakCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CritterSqueakDeletedIterator is returned from FilterSqueakDeleted and is used to iterate over the raw logs and unpacked data for SqueakDeleted events raised by the Critter contract.
type CritterSqueakDeletedIterator struct {
	Event *CritterSqueakDeleted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CritterSqueakDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CritterSqueakDeleted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CritterSqueakDeleted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CritterSqueakDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CritterSqueakDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CritterSqueakDeleted represents a SqueakDeleted event raised by the Critter contract.
type CritterSqueakDeleted struct {
	TokenId   *big.Int
	DeletedBy common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSqueakDeleted is a free log retrieval operation binding the contract event 0x4a69e658d9998f715a7a8f4e7e4282dfd9e2bb520e6e8a9aa3fce4c96245cb59.
//
// Solidity: event SqueakDeleted(uint256 tokenId, address deletedBy)
func (_Critter *CritterFilterer) FilterSqueakDeleted(opts *bind.FilterOpts) (*CritterSqueakDeletedIterator, error) {

	logs, sub, err := _Critter.contract.FilterLogs(opts, "SqueakDeleted")
	if err != nil {
		return nil, err
	}
	return &CritterSqueakDeletedIterator{contract: _Critter.contract, event: "SqueakDeleted", logs: logs, sub: sub}, nil
}

// WatchSqueakDeleted is a free log subscription operation binding the contract event 0x4a69e658d9998f715a7a8f4e7e4282dfd9e2bb520e6e8a9aa3fce4c96245cb59.
//
// Solidity: event SqueakDeleted(uint256 tokenId, address deletedBy)
func (_Critter *CritterFilterer) WatchSqueakDeleted(opts *bind.WatchOpts, sink chan<- *CritterSqueakDeleted) (event.Subscription, error) {

	logs, sub, err := _Critter.contract.WatchLogs(opts, "SqueakDeleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CritterSqueakDeleted)
				if err := _Critter.contract.UnpackLog(event, "SqueakDeleted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSqueakDeleted is a log parse operation binding the contract event 0x4a69e658d9998f715a7a8f4e7e4282dfd9e2bb520e6e8a9aa3fce4c96245cb59.
//
// Solidity: event SqueakDeleted(uint256 tokenId, address deletedBy)
func (_Critter *CritterFilterer) ParseSqueakDeleted(log types.Log) (*CritterSqueakDeleted, error) {
	event := new(CritterSqueakDeleted)
	if err := _Critter.contract.UnpackLog(event, "SqueakDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CritterSqueakInteractionIterator is returned from FilterSqueakInteraction and is used to iterate over the raw logs and unpacked data for SqueakInteraction events raised by the Critter contract.
type CritterSqueakInteractionIterator struct {
	Event *CritterSqueakInteraction // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CritterSqueakInteractionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CritterSqueakInteraction)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CritterSqueakInteraction)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CritterSqueakInteractionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CritterSqueakInteractionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CritterSqueakInteraction represents a SqueakInteraction event raised by the Critter contract.
type CritterSqueakInteraction struct {
	TokenId     *big.Int
	Sender      common.Address
	Interaction uint8
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSqueakInteraction is a free log retrieval operation binding the contract event 0x0fb3f9e279c76ba3d6187344cb6de16410acdb446be516723a7652c1826af76a.
//
// Solidity: event SqueakInteraction(uint256 tokenId, address sender, uint8 interaction)
func (_Critter *CritterFilterer) FilterSqueakInteraction(opts *bind.FilterOpts) (*CritterSqueakInteractionIterator, error) {

	logs, sub, err := _Critter.contract.FilterLogs(opts, "SqueakInteraction")
	if err != nil {
		return nil, err
	}
	return &CritterSqueakInteractionIterator{contract: _Critter.contract, event: "SqueakInteraction", logs: logs, sub: sub}, nil
}

// WatchSqueakInteraction is a free log subscription operation binding the contract event 0x0fb3f9e279c76ba3d6187344cb6de16410acdb446be516723a7652c1826af76a.
//
// Solidity: event SqueakInteraction(uint256 tokenId, address sender, uint8 interaction)
func (_Critter *CritterFilterer) WatchSqueakInteraction(opts *bind.WatchOpts, sink chan<- *CritterSqueakInteraction) (event.Subscription, error) {

	logs, sub, err := _Critter.contract.WatchLogs(opts, "SqueakInteraction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CritterSqueakInteraction)
				if err := _Critter.contract.UnpackLog(event, "SqueakInteraction", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSqueakInteraction is a log parse operation binding the contract event 0x0fb3f9e279c76ba3d6187344cb6de16410acdb446be516723a7652c1826af76a.
//
// Solidity: event SqueakInteraction(uint256 tokenId, address sender, uint8 interaction)
func (_Critter *CritterFilterer) ParseSqueakInteraction(log types.Log) (*CritterSqueakInteraction, error) {
	event := new(CritterSqueakInteraction)
	if err := _Critter.contract.UnpackLog(event, "SqueakInteraction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CritterStatusUpdatedIterator is returned from FilterStatusUpdated and is used to iterate over the raw logs and unpacked data for StatusUpdated events raised by the Critter contract.
type CritterStatusUpdatedIterator struct {
	Event *CritterStatusUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CritterStatusUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CritterStatusUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CritterStatusUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CritterStatusUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CritterStatusUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CritterStatusUpdated represents a StatusUpdated event raised by the Critter contract.
type CritterStatusUpdated struct {
	Account common.Address
	Status  uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterStatusUpdated is a free log retrieval operation binding the contract event 0xa7cc6673952dd2140cd171b21cdf21825b0f34e690500ecf38ed0aeb006f356e.
//
// Solidity: event StatusUpdated(address account, uint8 status)
func (_Critter *CritterFilterer) FilterStatusUpdated(opts *bind.FilterOpts) (*CritterStatusUpdatedIterator, error) {

	logs, sub, err := _Critter.contract.FilterLogs(opts, "StatusUpdated")
	if err != nil {
		return nil, err
	}
	return &CritterStatusUpdatedIterator{contract: _Critter.contract, event: "StatusUpdated", logs: logs, sub: sub}, nil
}

// WatchStatusUpdated is a free log subscription operation binding the contract event 0xa7cc6673952dd2140cd171b21cdf21825b0f34e690500ecf38ed0aeb006f356e.
//
// Solidity: event StatusUpdated(address account, uint8 status)
func (_Critter *CritterFilterer) WatchStatusUpdated(opts *bind.WatchOpts, sink chan<- *CritterStatusUpdated) (event.Subscription, error) {

	logs, sub, err := _Critter.contract.WatchLogs(opts, "StatusUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CritterStatusUpdated)
				if err := _Critter.contract.UnpackLog(event, "StatusUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStatusUpdated is a log parse operation binding the contract event 0xa7cc6673952dd2140cd171b21cdf21825b0f34e690500ecf38ed0aeb006f356e.
//
// Solidity: event StatusUpdated(address account, uint8 status)
func (_Critter *CritterFilterer) ParseStatusUpdated(log types.Log) (*CritterStatusUpdated, error) {
	event := new(CritterStatusUpdated)
	if err := _Critter.contract.UnpackLog(event, "StatusUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CritterTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Critter contract.
type CritterTransferIterator struct {
	Event *CritterTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CritterTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CritterTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CritterTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CritterTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CritterTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CritterTransfer represents a Transfer event raised by the Critter contract.
type CritterTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Critter *CritterFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*CritterTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Critter.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CritterTransferIterator{contract: _Critter.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Critter *CritterFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CritterTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Critter.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CritterTransfer)
				if err := _Critter.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Critter *CritterFilterer) ParseTransfer(log types.Log) (*CritterTransfer, error) {
	event := new(CritterTransfer)
	if err := _Critter.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CritterUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Critter contract.
type CritterUpgradedIterator struct {
	Event *CritterUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CritterUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CritterUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CritterUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CritterUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CritterUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CritterUpgraded represents a Upgraded event raised by the Critter contract.
type CritterUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Critter *CritterFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*CritterUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Critter.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &CritterUpgradedIterator{contract: _Critter.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Critter *CritterFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *CritterUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Critter.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CritterUpgraded)
				if err := _Critter.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Critter *CritterFilterer) ParseUpgraded(log types.Log) (*CritterUpgraded, error) {
	event := new(CritterUpgraded)
	if err := _Critter.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
