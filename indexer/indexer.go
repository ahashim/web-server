package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ahashim/web-server/contract"
	_ "github.com/joho/godotenv/autoload"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// load contract ABI
	contractABI := initContractABI()

	// create a channel for the event logs
	logs := make(chan types.Log)

	// create streaming event subscription
	subscription := initContractSubscription(
		os.Getenv("CONTRACT_WEBSOCKET_URL"),
		os.Getenv("CONTRACT_ADDRESS"),
		logs,
	)

	for {
		select {
		case err := <-subscription.Err():
			// close connection
			log.Fatal(err)
		case vLog := <-logs:
			// index events by signature
			switch vLog.Topics[0].Hex() {
			case contract.AccountCreatedSignature:
				// TODO: move this to an `indexer` package/struct to separate concerns
				indexAccountCreated(contractABI, vLog)
			}
		}
	}
}

// Loads the ABI interface from a smart-contracts abi.json file.
func initContractABI() abi.ABI {
	contractAbi, err := abi.JSON(strings.NewReader(string(contract.ContractMetaData.ABI)))
	if err != nil {
		log.Fatal(err)
	}

	return contractAbi
}

// Creates a subscription to an ethereum smart-contract.
func initContractSubscription(
	websocketURL string,
	contractAddress string,
	logs chan types.Log,
) ethereum.Subscription {
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

	// create event log subscription
	subscription, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	return subscription
}

// Indexes an `AccountCreated` event.
func indexAccountCreated(contractABI abi.ABI, vLog types.Log) {
	var event contract.AccountCreated

	// unpack event
	err := contractABI.UnpackIntoInterface(&event, contract.AccountCreatedEventName, vLog.Data)
	if err != nil {
		log.Fatal(err)
	}

	// populate indexed event data
	event.Account = common.HexToAddress(vLog.Topics[1].Hex())
	event.Username = vLog.Topics[2]

	// TODO: save to database via Ent
	fmt.Printf("\n")
	fmt.Printf("=================== %s ===================\n", contract.AccountCreatedEventName)
	fmt.Println("  Account:", event.Account.Hex())
	fmt.Println(" Username:", string(event.Username[:]))
}
