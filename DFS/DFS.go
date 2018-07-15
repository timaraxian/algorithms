package DFS

import "fmt"

type Node struct {
	Value string
}

type Graph struct {
	Edges map[Node]map[Node]bool
	Found map[Node]bool
	Order []Node
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

	return nodes
}

func (graph *Graph) inFound(n Node) bool {
	if _, ok := graph.Found[n]; !ok {
		graph.Found[n] = true
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

func DFT(graph Graph, s Node, order []Node) []Node {
	graph.inFound(s)
	neighbs := graph.GetNeighbours(s)
	for range neighbs {
		if !graph.inFound(neighbs[0]) {
			order = DFT(graph, neighbs[0], order)
		}
	}
	order = append(order, s)
	return order
}

func DFStopological(graph Graph) []Node {
	fmt.Println("--START--")
	allNodes := graph.GetNodes()
	fmt.Printf("all nodes: %v\n", allNodes)
	for i := range allNodes {
		if !graph.inFound(allNodes[i]) {
			neighbs := graph.GetNeighbours(allNodes[0])
			fmt.Printf("neighbours of %v are %v\n, found nodes are: %v\n", allNodes[i], neighbs, graph.Found)
			for j := range neighbs {
				graph.Order = DFT(graph, neighbs[j], graph.Order)
			}
			graph.Order = append(graph.Order, allNodes[i])
		}
	}

	return graph.Order
}
