package TopologicalSort

type Node struct {
	Value string
}

type Graph struct {
	Edges map[Node]map[Node]bool
	Nodes []Node
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
		Edges:      make(map[Node]map[Node]bool),
		Nodes:      make([]Node, 0),
		Found:      make(map[Node]bool),
		Order:      make([]Node, 0),
	}
}

// -----------------------------------------------------------------------------

func (graph *Graph) AddNode(n Node) {
	if _, ok := graph.Edges[n]; !ok {
		graph.Edges[n] = make(map[Node]bool)
		graph.Nodes = append(graph.Nodes, n)
	}
}

func (graph *Graph) AddDirectedEdge(n1, n2 Node) {
	if _, ok := graph.Edges[n1]; !ok {
		graph.Edges[n1] = make(map[Node]bool)
		graph.Nodes = append(graph.Nodes, n1)
	}

	if _, ok := graph.Edges[n1][n2]; !ok {
		graph.Edges[n1][n2] = true
	}

	if _, ok := graph.Edges[n2]; !ok {
		graph.Edges[n2] = make(map[Node]bool)
		graph.Nodes = append(graph.Nodes, n2)
	}
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
		return false
	}

	return true
}

func (graph *Graph) addFound(n Node) {
	if _, ok := graph.Found[n]; !ok {
		graph.Found[n] = true
	}
}

// -----------------------------------------------------------------------------

func nextLayer(graph *Graph, n Node) {
	neighbs := graph.GetNeighbours(n)
	for i := range neighbs {
		if !graph.inFound(neighbs[i]) {
			nextLayer(graph, neighbs[i])
			graph.addFound(neighbs[i])
		}
	}
	graph.Order = append(graph.Order, n)
}

func TopologicalSort(graph *Graph) {
	for i := range graph.Nodes {
		if graph.inFound(graph.Nodes[i]) {
			continue
		}
		neighbs := graph.GetNeighbours(graph.Nodes[i])
		for j := range neighbs {
			if !graph.inFound(neighbs[j]) {
				nextLayer(graph, neighbs[j])
			}
			graph.addFound(neighbs[j])
		}
		graph.Order = append(graph.Order, graph.Nodes[i])
		graph.addFound(graph.Nodes[i])
	}
}
