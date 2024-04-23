package graph

/*
Class UndirectedGraph
    Define 'adjacency_list' as a dictionary

    Function addVertex(vertex):
        if vertex not in adjacency_list:
            adjacency_list[vertex] = []

    Function addEdge(vertex1, vertex2):
        adjacency_list[vertex1].append(vertex2)
        adjacency_list[vertex2].append(vertex1)  // Symmetric as the graph is undirected

    Function removeEdge(vertex1, vertex2):
        adjacency_list[vertex1] = [v for v in adjacency_list[vertex1] if v != vertex2]
        adjacency_list[vertex2] = [v for v in adjacency_list[vertex2] if v != vertex1]

    Function removeVertex(vertex):
        for connected_vertex in adjacency_list[vertex]:
            adjacency_list[connected_vertex].remove(vertex)
        adjacency_list.pop(vertex, None)

*/
