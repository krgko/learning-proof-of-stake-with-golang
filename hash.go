package main

import (
	"crypto/sha256"
	"encoding/hex"
)

// Within main package, no need to import

func NewBlockHash(block *Block) string {
	blockInfo := block.Timestamp + block.PreviousHash + block.Hash + block.ValidatorAddress
	return newHash(blockInfo)
}

// Create a new hash, receive string and return string(hex)
func newHash(str string) string {
	// https://pkg.go.dev/crypto/sha256#New
	h := sha256.New()
	h.Write([]byte(str))
	hashed := h.Sum(nil)
	// https://pkg.go.dev/encoding/hex@go1.19.4#EncodeToString
	return hex.EncodeToString(hashed)
}
