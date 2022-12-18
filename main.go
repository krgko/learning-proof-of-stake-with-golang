package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	// Set a random seed
	rand.Seed(time.Now().UnixNano())

	// Generate an initial PoS Network including a blockchain with a genesis block
	genesisTimestamp := time.Now().String()
	network := &Network{
		Blockchain: []*Block{
			{
				Timestamp:        genesisTimestamp,
				PreviousHash:     "",
				Hash:             newHash(genesisTimestamp),
				ValidatorAddress: "",
			},
		},
	}
	// Current block is the genesis block
	network.BlockchainHead = network.Blockchain[0]
	// instantiate nodes to act as validators in our network
	network.Validators = network.NewNode(40)
	network.Validators = network.NewNode(60)

	// Build 10 additions to the blockchain
	for i := 0; i < 10; i++ {
		winner, err := network.SelectWinner()
		if err != nil {
			log.Fatal(err)
		}
		network.Blockchain, network.BlockchainHead, err = network.GenerateNewBlock(winner)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Round:", i)
		fmt.Println("\tAddress:", network.Validators[0].Address, "-Stake:", network.Validators[0].Stake)
		fmt.Println("\tAddress:", network.Validators[1].Address, "-Stake:", network.Validators[1].Stake)
	}

	network.PrintBlockchainInfo()
}
