package DFS

import "fmt"

type Node struct {
	Value string
}

type Graph struct {
	Edges map[Node]map[Node]bool
	Found map[Node]bool
}

// -----------------------------------------------------------------------------

func NewNode(value string) Node {
	return Node{
		Value: value,
	}
}

func NewGraph() Graph {
	return Graph{
		Edges: make(map[Node]map[Node]bool),
		Found: make(map[Node]bool),
	}
}

// -----------------------------------------------------------------------------

func (graph *Graph) AddNode(n Node) {
	if _, ok := graph.Edges[n]; !ok {
		graph.Edges[n] = make(map[Node]bool)
	}
}

func (graph *Graph) AddEdge(n1, n2 Node) {
	if _, ok := graph.Edges[n1]; !ok {
		graph.Edges[n1] = make(map[Node]bool)
	}

	if _, ok := graph.Edges[n1][n2]; !ok {
		graph.Edges[n1][n2] = true
	}

	if _, ok := graph.Edges[n2]; !ok {
		graph.Edges[n2] = make(map[Node]bool)
	}

	if _, ok := graph.Edges[n2][n1]; !ok {
		graph.Edges[n2][n1] = true
	}
}

func (graph *Graph) GetNodes() []Node {
	var nodes []Node

	for key := range graph.Edges {
		nodes = append(nodes, key)
	}

	return nodes
}

func (graph *Graph) GetNeighbours(n Node) []Node {
	var nodes []Node

	for value := range graph.Edges[n] {
		if _, ok := graph.Found[value]; !ok {
			nodes = append(nodes, value)
		}
	}
	fmt.Println(nodes)
	return nodes
}

func (graph *Graph) inFound(n Node) bool {
	if _, ok := graph.Found[n]; !ok {
		graph.Found[n] = true
		return false
	}

	return true
}

// -----------------------------------------------------------------------------

func DFS(graph Graph, s, search Node) bool {
	fmt.Println("--DFS CALL--")
	graph.inFound(s)
	neighbs := graph.GetNeighbours(s)
	for i := range neighbs {
		fmt.Printf("exploring: %v\n", neighbs[i])
		if neighbs[0] == search {
			return true
		}
		if !graph.inFound(neighbs[0]) {
			return DFS(graph, neighbs[0], search)
		}
	}
	return false
}