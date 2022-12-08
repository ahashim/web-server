package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ahashim/web-server/critter"
	_ "github.com/joho/godotenv/autoload"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type AccountCreated struct {
	Account  common.Address
	Username [32]byte
}

func main() {
	// create client
	client, err := ethclient.Dial(os.Getenv("CONTRACT_WEBSOCKET_URL"))
	if err != nil {
		log.Fatal(err)
	}

	// create filter query
	contractAddress := common.HexToAddress(os.Getenv("CONTRACT_ADDRESS"))
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	// load contract ABI
	contractAbi, err := abi.JSON(strings.NewReader(string(critter.CritterMetaData.ABI)))
	if err != nil {
		log.Fatal(err)
	}

	// create a channel for the logs
	logs := make(chan types.Log)

	// subscribe to them using the query
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	// calculate event signatures
	sigAccountCreated := crypto.Keccak256Hash([]byte("AccountCreated(address,bytes32)")).Hex()

	for {
		select {
		case err := <-sub.Err():
			// connection closed
			log.Fatal(err)
		case vLog := <-logs:
			// index events
			switch vLog.Topics[0].Hex() {
			case sigAccountCreated:
				// testing with `AccountCreated` for now as a proof-of-concept
				fmt.Println("event: AccountCreated")
				var event AccountCreated

				// unpack event
				err := contractAbi.UnpackIntoInterface(&event, "AccountCreated", vLog.Data)
				if err != nil {
					log.Fatal(err)
				}

				// populate event data
				event.Account = common.HexToAddress(vLog.Topics[1].Hex())
				event.Username = vLog.Topics[2]

				// TODO: save to database
				fmt.Println("Account:", event.Account.Hex())
				fmt.Println("Username:", string(event.Username[:]))
			}
		}
	}
}
