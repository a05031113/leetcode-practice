package main

import "fmt"

type Node struct {
	Val int
	Neighbors []*Node
}


func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	visited := make(map[*Node]*Node)

	var clone func(node *Node) *Node
	clone = func(node *Node) *Node {
		if copy, exists := visited[node]; exists {
			return copy
		}

		copy := &Node{
			Val: node.Val,
			Neighbors: make([]*Node, 0, len(node.Neighbors)),
		}
		visited[node] = copy
		
		for _, neighbor := range node.Neighbors {
			copy.Neighbors = append(copy.Neighbors, clone(neighbor))
		}
		return copy 
	}


    return clone(node)
}


func createTestGraph() *Node {
	// Create 4 nodes
	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	node4 := &Node{Val: 4}
	
	// Connect them in an undirected manner
	node1.Neighbors = []*Node{node2, node4}
	node2.Neighbors = []*Node{node1, node3}
	node3.Neighbors = []*Node{node2, node4}
	node4.Neighbors = []*Node{node1, node3}
	
	return node1
}

func printGraph(node *Node) {
	if node == nil {
		fmt.Println("Empty graph")
		return
	}
	
	visited := make(map[*Node]bool)
	queue := []*Node{node}
	visited[node] = true
	
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		
		fmt.Printf("Node %d -> Neighbors: [", current.Val)
		for i, neighbor := range current.Neighbors {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Print(neighbor.Val)
			
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
		fmt.Println("]")
	}
}

func main() {
	// Create a test graph
	originalGraph := createTestGraph()
	fmt.Println("Original Graph:")
	printGraph(originalGraph)
	
	// Clone the graph
	clonedGraph := cloneGraph(originalGraph)
	fmt.Println("\nCloned Graph:")
	printGraph(clonedGraph)
	
	// Verify they are different graph instances
	fmt.Printf("\nOriginal graph root address: %p\n", originalGraph)
	fmt.Printf("Cloned graph root address: %p\n", clonedGraph)
	
	// Modify the original graph to prove they're independent
	if len(originalGraph.Neighbors) > 0 {
		originalGraph.Neighbors[0].Val = 99
	}
	
	fmt.Println("\nAfter modifying original graph:")
	fmt.Println("Original Graph:")
	printGraph(originalGraph)
	fmt.Println("\nCloned Graph (should remain unchanged):")
	printGraph(clonedGraph)
}