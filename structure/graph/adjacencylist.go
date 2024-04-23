package graph

/*
Creating a basic graph data structure in Go can be done using an adjacency list, which is a common way to represent graphs. This implementation will include methods to add vertices, add edges, and represent both directed and undirected graphs. Additionally, I'll include a simple way to print the graph for verification.

Step 1: Define the Graph Structure
Here, we'll define the Graph structure that uses a map to hold adjacency lists for each vertex.

go
Copy code
package main

import (
    "fmt"
)

// Graph structure
type Graph struct {
    adjacencyList map[string][]string
}

// NewGraph creates a new graph
func NewGraph() *Graph {
    return &Graph{
        adjacencyList: make(map[string][]string),
    }
}
Step 2: Add Vertex
This method adds a new vertex to the graph if it does not already exist.

go
Copy code
// AddVertex adds a new vertex to the graph
func (g *Graph) AddVertex(vertex string) {
    if _, exists := g.adjacencyList[vertex]; !exists {
        g.adjacencyList[vertex] = []string{}
    }
}
Step 3: Add Edge
This method adds an edge to the graph. You can modify it to handle directed or undirected graphs by adjusting how the edges are stored.

go
Copy code
// AddEdge adds an edge between two vertices (for an undirected graph)
func (g *Graph) AddEdge(vertex1, vertex2 string) {
    // Since this is an undirected graph, we add the edge from vertex1 to vertex2
    // and from vertex2 to vertex1 to make sure the adjacency is reciprocal.
    g.adjacencyList[vertex1] = append(g.adjacencyList[vertex1], vertex2)
    g.adjacencyList[vertex2] = append(g.adjacencyList[vertex2], vertex1)
}
For a directed graph, you would only add the edge from vertex1 to vertex2, indicating a one-way connection.

Step 4: Print the Graph
This utility function helps visualize the graph by printing each vertex and its adjacency list.

go
Copy code
// PrintGraph prints the adjacency list of each vertex
func (g *Graph) PrintGraph() {
    for vertex, edges := range g.adjacencyList {
        fmt.Printf("%v -> ", vertex)
        for _, edge := range edges {
            fmt.Printf("%v ", edge)
        }
        fmt.Println()
    }
}
Step 5: Main Function
We'll test the graph by adding some vertices and edges, and then printing the graph.

go
Copy code
func main() {
    graph := NewGraph()
    vertices := []string{"A", "B", "C", "D"}

    for _, vertex := range vertices {
        graph.AddVertex(vertex)
    }

    graph.AddEdge("A", "B")
    graph.AddEdge("A", "C")
    graph.AddEdge("B", "C")
    graph.AddEdge("C", "D")

    graph.PrintGraph()
}
This program will output:

css
Copy code
A -> B C
B -> A C
C -> A B D
D -> C
This setup creates an undirected graph. If you're interested in a directed graph, adjust the AddEdge method to only add an edge from vertex1 to vertex2. This basic framework provides a solid foundation for more advanced graph operations such as searching, traversing, and pathfinding algorithms like Depth-First Search, Breadth-First Search, and Dijkstra's Algorithm.
*/

/*

Step 1: Adding Depth-First Search (DFS)
Depth-First Search is a traversal method that explores as far down a branch as possible before backtracking.

go
Copy code
// DFS performs a depth-first search starting from the given vertex
func (g *Graph) DFS(startVertex string) {
    visited := make(map[string]bool)
    g.dfsHelper(startVertex, visited)
}

// dfsHelper is a recursive helper function for DFS
func (g *Graph) dfsHelper(vertex string, visited map[string]bool) {
    visited[vertex] = true
    fmt.Println(vertex)

    for _, neighbor := range g.adjacencyList[vertex] {
        if !visited[neighbor] {
            g.dfsHelper(neighbor, visited)
        }
    }
}
Step 2: Adding Breadth-First Search (BFS)
Breadth-First Search explores all the neighbors at the present depth prior to moving on to nodes at the next depth level.

go
Copy code
// BFS performs a breadth-first search starting from the given vertex
func (g *Graph) BFS(startVertex string) {
    visited := make(map[string]bool)
    queue := []string{startVertex}

    for len(queue) > 0 {
        vertex := queue[0]
        queue = queue[1:]
        if !visited[vertex] {
            visited[vertex] = true
            fmt.Println(vertex)
            for _, neighbor := range g.adjacencyList[vertex] {
                if !visited[neighbor] {
                    queue = append(queue, neighbor)
                }
            }
        }
    }
}
Step 3: Implementing Dijkstra's Algorithm
Dijkstra's Algorithm is used for finding the shortest path from a starting node to all other nodes in the graph.

First, we need to adjust our Graph structure to handle weights. Let's redefine the Graph class to accommodate weights:

go
Copy code
type Graph struct {
    adjacencyList map[string][]*Node
}

type Node struct {
    vertex string
    weight int
}

func NewGraph() *Graph {
    return &Graph{
        adjacencyList: make(map[string][]*Node),
    }
}

func (g *Graph) AddEdge(vertex1, vertex2 string, weight int) {
    g.adjacencyList[vertex1] = append(g.adjacencyList[vertex1], &Node{vertex: vertex2, weight: weight})
    // If undirected graph
    g.adjacencyList[vertex2] = append(g.adjacencyList[vertex2], &Node{vertex: vertex1, weight: weight})
}
Now, implement Dijkstra's algorithm:

go

Step 4: Main Function

Completing the Main Function:

go
Copy code
func main() {
    graph := NewGraph()
    graph.AddVertex("A")
    graph.AddVertex("B")
    graph.AddVertex("C")
    graph.AddVertex("D")
    graph.AddEdge("A", "B", 1)
    graph.AddEdge("A", "C", 3)
    graph.AddEdge("B", "C", 1)
    graph.AddEdge("B", "D", 5)
    graph.AddEdge("C", "D", 4)

    fmt.Println("DFS:")
    graph.DFS("A")

    fmt.Println("BFS:")
    graph.BFS("A")

    fmt.Println("Dijkstra's Algorithm:")
    distances := graph.Dijkstra("A")
    for vertex, distance := range distances {
        fmt.Printf("Distance from A to %s: %d\n", vertex, distance)
    }
}

*/
