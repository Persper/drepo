package main

import (
	"fmt"
	// "github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/ethclient"
)

func Connector() {
	// Create an IPC based RPC connection to a remote node
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		fmt.Printf("err: %v \n", err)
		return
	}
	fmt.Printf("Hello %s", client)
}
