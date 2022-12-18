package main

import (
	"fmt"
	"math/rand"
)

// Artificial nodes, since we are not create the real nodes
func (n Network) NewNode(stake int) []*Node {
	newNode := &Node{
		Stake:   stake,
		Address: randomAddress(),
	}
	n.Validators = append(n.Validators, newNode)
	return n.Validators
}

// https://pkg.go.dev/math/rand@go1.19.4
// Random address 16 chars
func randomAddress() string {
	byteToUse := make([]byte, 16)
	_, _ = rand.Read(byteToUse) // Read regerates len(byteToUse) and write them into byteToUse
	return fmt.Sprintf("%x", byteToUse)
}
