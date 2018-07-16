package DFS

import "fmt"

type Node struct {
	Value string
}

type Graph struct {
	Edges map[Node]map[Node]bool
	Found      map[Node]bool
	LocalFound map[Node]bool
	Order      map[Node]int
	OrderFound map[Node]bool
}

// -----------------------------------------------------------------------------

func NewNode(value string) Node {
	return Node{
		Value: value,
	}
}

func NewGraph() Graph {
	return Graph{
		Edges:      make(map[Node]map[Node]bool),
		Found:      make(map[Node]bool),
		LocalFound: make(map[Node]bool),
		Order:      make(map[Node]int),
		OrderFound:      make(map[Node]bool),
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

	return nodes
}

func (graph *Graph) inFound(n Node) bool {
	if _, ok := graph.Found[n]; !ok {
		graph.Found[n] = true
		return false
	}

	return true
}

func (graph *Graph) inLocalFound(n Node) bool {
	if _, ok := graph.LocalFound[n]; !ok {
		graph.LocalFound[n] = true
		return false
	}

	return true
}

func (graph *Graph) AddDirectedEdge(n1, n2 Node) {
	if _, ok := graph.Edges[n1]; !ok {
		graph.Edges[n1] = make(map[Node]bool)
	}

	if _, ok := graph.Edges[n1][n2]; !ok {
		graph.Edges[n1][n2] = true
	}
}

// -----------------------------------------------------------------------------

func DFS(graph Graph, s, search Node) bool {
	graph.inFound(s)
	neighbs := graph.GetNeighbours(s)
	for range neighbs {
		if neighbs[0] == search {
			return true
		}
		if !graph.inFound(neighbs[0]) {
			return DFS(graph, neighbs[0], search)
		}
	}
	return false
}

func DFT(graph Graph, s Node, increment int) int {
	fmt.Printf("exploring: %v increment: %v\n", s, increment)
	graph.inLocalFound(s)
	neighbs := graph.GetNeighbours(s)
	for range neighbs {
		if !graph.inLocalFound(neighbs[0]) {
			increment = DFT(graph, neighbs[0], increment)
		}
	}
	graph.LocalFound = make(map[Node]bool)
	if _, ok := graph.OrderFound[s]; !ok {
		fmt.Printf("DFT: putting %v into order with increment of %v\n", s, increment)
		graph.Order[s] = increment
		graph.OrderFound[s] = true
		increment --
	}
	return increment
}

func DFStopological(graph Graph) map[Node]int {
	allNodes := graph.GetNodes()
	fmt.Printf("allnodes: %v\n", allNodes)
	increment := len(allNodes)
	for increment > 0 {
		fmt.Printf("DFS OUTER ON: %v\n", allNodes[increment-1])
		if !graph.inFound(allNodes[increment-1]) {
			neighbs := graph.GetNeighbours(allNodes[increment-1])
			for j := range neighbs {
				increment = DFT(graph, neighbs[j], increment)
				graph.inFound(neighbs[j])
			}
		}
		if _, ok := graph.OrderFound[allNodes[increment-1]]; !ok {
			fmt.Printf("main: putting %v into order with increment %v\n", allNodes[increment-1], increment)
			graph.Order[allNodes[increment-1]] = increment
			graph.OrderFound[allNodes[increment-1]] = true

		}
		increment --
	}

	return graph.Order
}
