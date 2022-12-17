package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Node represents a node in the network
type Node struct {
	Value  int
	Votes  int
	Locked bool
}

// Network represents the network of nodes
type Network struct {
	nodes []*Node
	mu    sync.Mutex
}

// NewNetwork creates a new network of nodes
func NewNetwork() *Network {
	return &Network{nodes: make([]*Node, 0)}
}

// AddNode adds a new node to the network
func (n *Network) AddNode(node *Node) {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.nodes = append(n.nodes, node)
}

// RunConsensus runs the consensus protocol to reach agreement on a value
func (n *Network) RunConsensus() int {
	for {
		// Each node proposes a value
		for _, node := range n.nodes {
			node.Locked = false
			node.Votes = 1
		}

		// Wait for all nodes to propose a value
		time.Sleep(time.Second)

		// Calculate the votes for each value
		votes := make(map[int]int)
		for _, node := range n.nodes {
			votes[node.Value]++
		}

		// Find the value with the most votes
		maxVotes := 0
		var agreedValue int
		for value, count := range votes {
			if count > maxVotes {
				maxVotes = count
				agreedValue = value
			}
		}

		// Check if a majority of nodes agree on the value
		if maxVotes > len(n.nodes)/2 {
			fmt.Println("Consensus reached! Value =", agreedValue)
			return agreedValue
		}
	}
}

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Create a network of nodes
	network := NewNetwork()
	for i := 0; i < 5; i++ {
		node := &Node{Value: rand.Intn(10)}
		network.AddNode(node)
	}

	// Run the consensus protocol
	agreedValue := network.RunConsensus()
	fmt.Println("Agreed value:", agreedValue)
}
