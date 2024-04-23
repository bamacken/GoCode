package graph

// A* and Dijkstra's algorithm are used to find the shortest path between two nodes in a weighted graph.
/*


Class WeightedGraph
    Define 'adjacency_list' as a dictionary

    Function addVertex(vertex):
        if vertex not in adjacency_list:
            adjacency_list[vertex] = []

    Function addEdge(vertex1, vertex2, weight):
        adjacency_list[vertex1].append((vertex2, weight))
        adjacency_list[vertex2].append((vertex1, weight))  // If undirected

    Function removeEdge(vertex1, vertex2):
        adjacency_list[vertex1] = [item for item in adjacency_list[vertex1] if item[0] != vertex2]
        adjacency_list[vertex2] = [item for item in adjacency_list[vertex2] if item[0] != vertex1]

    Function removeVertex(vertex):
        for (connected_vertex, _) in adjacency_list[vertex]:
            adjacency_list[connected_vertex] = [item for item in adjacency_list[connected_vertex] if item[0] != vertex]
        adjacency_list.pop(vertex, None)

*/

/*
Function BFS(graph, start_vertex):
    Define 'queue' initialized with start_vertex
    Define 'visited' as an empty set

    While queue is not empty:
        current_vertex = queue.pop(0)
        if current_vertex not in visited:
            Print current_vertex
            visited.add(current_vertex)
            queue.extend(graph.adjacency_list[current_vertex] - visited)

Function DFS(graph, start_vertex):
    Define 'stack' initialized with start_vertex
    Define 'visited' as an empty set

    While stack is not empty:
        current_vertex = stack.pop()
        if current_vertex not in visited:
            Print current_vertex
            visited.add(current_vertex)
            stack.extend(graph.adjacency_list[current_vertex] - visited)

*/

/*
Copy code
// Dijkstra finds the shortest path from startVertex to all other vertices
func (g *Graph) Dijkstra(startVertex string) map[string]int {
    distances := make(map[string]int)
    visited := make(map[string]bool)

    // Initialize distances
    for vertex := range g.adjacencyList {
        distances[vertex] = int(^uint(0) >> 1) // Max int
    }
    distances[startVertex] = 0

    for len(visited) < len(g.adjacencyList) {
        // Vertex with the minimum distance which hasn't been visited
        minVertex := ""
        minDistance := int(^uint(0) >> 1) // Max int
        for vertex, distance := range distances {
            if !visited[vertex] && distance <= minDistance {
                minDistance = distance
                minVertex = vertex
            }
        }

        visited[minVertex] = true

        for _, edge := range g.adjacencyList[minVertex] {
            if !visited[edge.vertex] {
                newDist := distances[minVertex] + edge.weight
                if newDist < distances[edge.vertex] {
                    distances[edge.vertex] = newDist
                }
            }
        }
    }

    return distances
}
*/
