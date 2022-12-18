package main

import (
	"errors"
	"math/rand"
)

func (n Network) SelectWinner() (*Node, error) {
	var winnerPool []*Node
	totalStake := 0
	// for index, value
	for _, node := range n.Validators {
		if node.Stake > 0 {
			winnerPool = append(winnerPool, node)
			totalStake += node.Stake
		}
	}
	// No one in the winner pool, The PoS cannot occurs
	if winnerPool == nil {
		return nil, errors.New("there are no nodes with stake in the network")
	}
	validatorTickets := 0
	// https://pkg.go.dev/math/rand@go1.19.4#Intn
	// rand.Seed(x) may optional
	winnerNumber := rand.Intn(totalStake)
	for _, node := range n.Validators {
		// Visualize: [0-10, 11-30, 31-40] -> can represent as percentage
		// Assumption:
		// A stakes 10 tokens
		// B stakes 20 tokens
		// C stakes 10 tokens
		// B has a lot of chance to win
		// winnerNumber(25) < validatorTickets(30)
		// B is the winner, because B tickets is in the interval
		validatorTickets += node.Stake
		if winnerNumber < validatorTickets {
			return node, nil
		}
	}
	// No winner
	return nil, errors.New("a winner have not been picked")
}
