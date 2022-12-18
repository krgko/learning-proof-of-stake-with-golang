package main

import (
	"fmt"
	"time"
)

type Network struct {
	Blockchain     []*Block // Array of references as an instance of the block
	BlockchainHead *Block   // Most recently added block
	Validators     []*Node  // Array of references as validators
}

type Node struct {
	Stake   int // A number of tokens staked and added to the network
	Address string
}

// Block is to keep track of Blockchain
type Block struct {
	Timestamp        string // When created
	PreviousHash     string // A previous block reference
	Hash             string // A current block
	ValidatorAddress string // Who validated the block
}

// Use receiver function, like self in other language
// A logic for new block generation
// 1. validate blockchain
// 2. create a new block
// 3. validate blockchain candiate
// 4. passed then append to the blockchain
func (n Network) GenerateNewBlock(Validator *Node) ([]*Block, *Block, error) {
	if err := n.ValidateBlockchain(); err != nil {
		Validator.Stake = 0 // Slash all if the validator got penalized
		return n.Blockchain, n.BlockchainHead, err
	}

	currentTime := time.Now().String()

	newBlock := &Block{
		Timestamp:        currentTime,
		PreviousHash:     n.BlockchainHead.Hash,
		Hash:             NewBlockHash(n.BlockchainHead), // Current block hash
		ValidatorAddress: Validator.Address,
	}

	if err := n.ValidateBlockCandidate(newBlock); err != nil {
		Validator.Stake = 0 // Slash all if the validator got penalized
		return n.Blockchain, n.BlockchainHead, err
	} else {
		Validator.Stake += 1 // Reward to the validator
		n.Blockchain = append(n.Blockchain, newBlock)
	}

	// No error return
	return n.Blockchain, newBlock, nil
}

func (n Network) PrintBlockchainInfo() {
	for index, block := range n.Blockchain {
		fmt.Println("Block", index, "Info:")
		fmt.Println("\tTimestamp:", block.Timestamp)
		fmt.Println("\tPrevious Hash:", block.PreviousHash)
		fmt.Println("\tHash:", block.Hash)
		fmt.Println("\tValidator Address:", block.ValidatorAddress)
	}
}
