package main

import (
	"errors"
)

func (n *Network) ValidateBlockchain() error {
	// Skip verification if the Blockchain just contains one block or less
	if len(n.Blockchain) <= 1 {
		return nil
	}

	// If refer to an array index, position would be -1 always
	currentBlockIndex := len(n.Blockchain) - 1
	previousBlockIndex := len(n.Blockchain) - 2

	// for true, do
	for previousBlockIndex >= 0 {
		currentBlock := n.Blockchain[currentBlockIndex]
		previousBlock := n.Blockchain[previousBlockIndex]
		// Hashes should not duplicated
		if currentBlock.PreviousHash != previousBlock.Hash {
			return errors.New("the blockchain has inconsistent hashes")
		}
		// The previous block should not created after the current block
		if currentBlock.Timestamp <= previousBlock.Timestamp {
			return errors.New("the blockchain has inconsistent timestamps")
		}
		// Hash should valid; created hash should be the same with hash in the block
		if NewBlockHash(previousBlock) != currentBlock.Hash {
			return errors.New("the blockchain has inconsistent hash generation")
		}

		// Check back until the first block
		currentBlockIndex-- // Minus 1
		previousBlockIndex--
	}
	return nil
}

// Make sure that the next block that will be added is also valid
func (n Network) ValidateBlockCandidate(newBlock *Block) error {
	if n.BlockchainHead.Hash != newBlock.PreviousHash {
		return errors.New("the blockchain HEAD is not equal to the new block previous hash")
	}
	if n.BlockchainHead.Timestamp >= newBlock.Timestamp {
		return errors.New("the blockchain HEAD timestamp is greater than or equal to the new block timestamp")
	}
	if NewBlockHash(n.BlockchainHead) != newBlock.Hash {
		return errors.New("the new block hash of Blockchain HEAD does not equal to the new block hash")
	}
	return nil
}
