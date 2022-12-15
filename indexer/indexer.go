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
	// load contract ABI
	abi := initContractABI()

	// create a channel for the event logs
	logs := make(chan types.Log)

	// create an ethereum client & contract subscription
	ec, sub := initContractSubscription(
		os.Getenv("CONTRACT_WEBSOCKET_URL"),
		os.Getenv("CONTRACT_ADDRESS"),
		logs,
	)
	defer ec.Close()

	// create ent ORM client
	orm := initORM()
	// defer orm.Close()

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
				indexAccountCreated(vLog, abi, orm)
			}
		}
	}
}

// Loads the ABI interface from a smart-contracts abi.json file.
func initContractABI() abi.ABI {
	contractAbi, err := abi.JSON(strings.NewReader(string(contract.ContractMetaData.ABI)))
	if err != nil {
		log.Fatalf("failed to load contract ABI: %v", err)
	}

	return contractAbi
}

// Creates a subscription to an ethereum smart-contract.
func initContractSubscription(
	websocketURL string,
	contractAddress string,
	logs chan types.Log,
) (*ethclient.Client, ethereum.Subscription) {
	// create an ethereum client
	client, err := ethclient.Dial(websocketURL)
	if err != nil {
		log.Fatal(err)
	}

	// create the filter query
	address := common.HexToAddress(contractAddress)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{address},
	}

	// create event log sub
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatalf("failed to subscribe to contract events: %v", err)
	}

	return client, sub
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
