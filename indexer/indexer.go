package main

import (
	"context"
	"fmt"
	"log"
	"os"

	// autoload environment variables
	_ "github.com/joho/godotenv/autoload"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// create client
	client, err := ethclient.Dial(os.Getenv("CONTRACT_WEBSOCKET_URL"))
	if err != nil {
		log.Fatal(err)
	}

	// connect to contract
	contractAddress := common.HexToAddress(os.Getenv("CONTRACT_ADDRESS"))
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	// create a channel for the logs
	logs := make(chan types.Log)

	// subscribe to them using the query
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println(vLog)
		}
	}
}
