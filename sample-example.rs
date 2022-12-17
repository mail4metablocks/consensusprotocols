extern crate uuid;

use std::collections::HashMap;

// A node in the network
struct Node {
    id: uuid::Uuid,
    value: u32,
    votes: u32,
}

impl Node {
    fn new(value: u32) -> Self {
        Node {
            id: uuid::Uuid::new_v4(),
            value,
            votes: 1,
        }
    }
}

// The network of nodes
struct Network {
    nodes: Vec<Node>,
}

impl Network {
    fn new() -> Self {
        Network { nodes: Vec::new() }
    }

    // Add a new node to the network
    fn add_node(&mut self, node: Node) {
        self.nodes.push(node);
    }

    // Run the consensus protocol to reach agreement on a value
    fn run_consensus(&self) -> u32 {
        let mut current_value = 0;
        let mut current_votes = 0;

        // Run the consensus protocol until all nodes agree on a value
        loop {
            // Each node proposes a value
            for node in &self.nodes {
                if node.value > current_value {
                    current_value = node.value;
                    current_votes = node.votes;
                } else if node.value == current_value {
                    current_votes += node.votes;
                }
            }

            // Check if a majority of nodes agree on the current value
            if current_votes > self.nodes.len() as u32 / 2 {
                println!("Consensus reached! Value = {}", current_value);
                return current_value;
            }

            // Reset the current value and votes if consensus was not reached
            current_value = 0;
            current_votes = 0;
        }
    }
}

fn main() {
    // Create a network of nodes
    let mut network = Network::new();
    network.add_node(Node::new(1));
    network.add_node(Node::new(2));
    network.add_node(Node::new(3));
    network.add_node(Node::new(2));
    network.add_node(Node::new(1));

    // Run the consensus protocol
    let agreed_value = network.run_consensus();
    println!("Agreed value: {}", agreed_value);
}
  
