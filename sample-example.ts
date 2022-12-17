import { v4 as uuidv4 } from 'uuid';

// A node in the network
class Node {
  public id: string;
  public value: number;
  public votes: number;

  constructor(value: number) {
    this.id = uuidv4();
    this.value = value;
    this.votes = 1;
  }
}

// The network of nodes
class Network {
  public nodes: Node[];

  constructor() {
    this.nodes = [];
  }

  // Add a new node to the network
  public addNode(node: Node) {
    this.nodes.push(node);
  }

  // Run the consensus protocol to reach agreement on a value
  public async runConsensus() {
    let currentValue = 0;
    let currentVotes = 0;

    // Run the consensus protocol until all nodes agree on a value
    while (true) {
      // Each node proposes a value
      for (const node of this.nodes) {
        if (node.value > currentValue) {
          currentValue = node.value;
          currentVotes = node.votes;
        } else if (node.value === currentValue) {
          currentVotes += node.votes;
        }
      }

      // Check if a majority of nodes agree on the current value
      if (currentVotes > this.nodes.length / 2) {
        console.log(`Consensus reached! Value = ${currentValue}`);
        return currentValue;
      }

      // Reset the current value and votes if consensus was not reached
      currentValue = 0;
      currentVotes = 0;
    }
  }
}

async function main() {
  // Create a network of nodes
  const network = new Network();
  network.addNode(new Node(1));
  network.addNode(new Node(2));
  network.addNode(new Node(3));
  network.addNode(new Node(2));
  network.addNode(new Node(1));

  // Run the consensus protocol
  const agreedValue = await network.runConsensus();
  console.log(`Agreed value: ${agreedValue}`);
}

main().catch((error) => {
  console.error(error);
  process.exit(1);
});
