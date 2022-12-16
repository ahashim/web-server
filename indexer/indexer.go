package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"entgo.io/ent/dialect"

	"github.com/ahashim/web-server/config"
	"github.com/ahashim/web-server/contract"
	"github.com/ahashim/web-server/ent"

	_ "github.com/joho/godotenv/autoload" // .env files
	_ "github.com/lib/pq"                 // postgres driver

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// event logs channel
	logs := make(chan types.Log)

	// critter abi
	contract_abi, err := newContractABI()
	if err != nil {
		log.Fatalf("failed to load contract ABI: %v", err)
	}

	// ethereum client
	ec, err := newEthClient(os.Getenv("CONTRACT_WEBSOCKET_URL"))
	defer ec.Close()
	if err != nil {
		log.Fatalf("failed to connect to Ethereum node: %v", err)
	}

	// contract subscription
	sub, err := newContractSubscription(
		common.HexToAddress(os.Getenv("CONTRACT_ADDRESS")),
		ec,
		logs,
	)
	if err != nil {
		log.Fatalf("failed to subscribe to contract: %v", err)
	}

	// ent ORM
	orm := initORM()
	defer orm.Close()

	// stream events
	for {
		select {
		case err := <-sub.Err():
			// close connection
			log.Fatal(err)
		case vLog := <-logs:
			// index events by signature
			switch vLog.Topics[0].Hex() {
			case contract.AccountCreatedSignature:
				// TODO: move this to an `indexer` package/struct to separate concerns
				indexAccountCreated(vLog, contract_abi, orm)
			}
		}
	}
}

// Loads the smart-contract ABI interface from a json file.
func newContractABI() (abi.ABI, error) {
	contract_abi, err := abi.JSON(strings.NewReader(string(contract.ContractMetaData.ABI)))
	if err != nil {
		return contract_abi, err
	}

	return contract_abi, nil
}

// Creates an ethereum client connected to a node.
func newEthClient(node_url string) (*ethclient.Client, error) {
	ec, err := ethclient.Dial(node_url)
	if err != nil {
		return ec, err
	}

	return ec, nil
}

// Creates a subscription to an ethereum smart-contract.
func newContractSubscription(
	address common.Address,
	ec *ethclient.Client,
	logs chan types.Log,
) (ethereum.Subscription, error) {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{address},
	}

	sub, err := ec.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		return sub, err
	}

	return sub, nil
}

func initORM() *ent.Client {
	// init config
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// connect to database
	orm, err := ent.Open(
		dialect.Postgres,
		fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			cfg.Database.Hostname,
			cfg.Database.Port,
			cfg.Database.User,
			cfg.Database.Password,
			cfg.Database.Database,
		),
	)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	return orm
}

// Indexes an `AccountCreated` event.
func indexAccountCreated(vLog types.Log, contractABI abi.ABI, orm *ent.Client) {
	var event contract.AccountCreated

	// unpack event
	err := contractABI.UnpackIntoInterface(&event, contract.AccountCreatedEventName, vLog.Data)
	if err != nil {
		log.Fatal(err)
	}

	// populate indexed event data
	event.Account = common.HexToAddress(vLog.Topics[1].Hex())
	event.Username = vLog.Topics[2]

	// remove null bytes
	username := string(bytes.Trim(event.Username[:], "\x00"))

	// save to db
	user, err := orm.User.Create().
		SetAddress(event.Account.Hex()).
		SetStatus("ACTIVE").
		SetUsername(username).
		Save(context.Background())
	if err != nil {
		log.Fatalf("failed creating user: %s", err)
	}

	// log creation event
	fmt.Printf("=================== %s ===================\n", contract.AccountCreatedEventName)
	fmt.Println("Address:", user.Address)
	fmt.Println("Username:", user.Username)
}
