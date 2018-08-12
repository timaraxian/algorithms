package StronglyConnectedComponents

type Vertex struct {
	Value string
}

type Graph struct {
	Arcs       map[Vertex]map[Vertex]bool
	Vertices   []Vertex
	Order      []Vertex
	Found      map[Vertex]bool
	Found2nd   map[Vertex]bool
	Components map[Vertex]map[Vertex]bool
}

// -----------------------------------------------------------------------------

func NewVertex(value string) Vertex {
	return Vertex{
		Value: value,
	}
}

func NewGraph() Graph {
	return Graph{
		Arcs:       make(map[Vertex]map[Vertex]bool),
		Vertices:   make([]Vertex, 0),
		Order:      make([]Vertex, 0),
		Found:      make(map[Vertex]bool),
		Found2nd:   make(map[Vertex]bool),
		Components: make(map[Vertex]map[Vertex]bool),
	}
}

// -----------------------------------------------------------------------------

func (graph *Graph) AddVertex(n Vertex) {
	if _, ok := graph.Arcs[n]; !ok {
		graph.Arcs[n] = make(map[Vertex]bool)
		graph.Vertices = append(graph.Vertices, n)
	}
}

func (graph *Graph) AddArc(n1, n2 Vertex) {
	if _, ok := graph.Arcs[n1]; !ok {
		graph.Arcs[n1] = make(map[Vertex]bool)
		graph.Vertices = append(graph.Vertices, n1)
	}

	if _, ok := graph.Arcs[n1][n2]; !ok {
		graph.Arcs[n1][n2] = true
	}

	if _, ok := graph.Arcs[n2]; !ok {
		graph.Arcs[n2] = make(map[Vertex]bool)
		graph.Vertices = append(graph.Vertices, n2)
	}
}

func (graph *Graph) GetNeighbours(n Vertex) []Vertex {
	var vertices []Vertex

	for value := range graph.Arcs[n] {
		if _, ok := graph.Found[value]; !ok {
			vertices = append(vertices, value)
		}
	}

	return vertices
}

func (graph *Graph) inFound(n Vertex) bool {
	if _, ok := graph.Found[n]; !ok {
		return false
	}

	return true
}

func (graph *Graph) addFound(n Vertex) {
	if _, ok := graph.Found[n]; !ok {
		graph.Found[n] = true
	}
}

func (graph *Graph) GetNeighbours2nd(n Vertex) []Vertex {
	var vertices []Vertex

	for value := range graph.Arcs[n] {
		if _, ok := graph.Found2nd[value]; !ok {
			vertices = append(vertices, value)
		}
	}

	return vertices
}

func (graph *Graph) inFound2nd(n Vertex) bool {
	if _, ok := graph.Found2nd[n]; !ok {
		return false
	}

	return true
}

func (graph *Graph) addFound2nd(n Vertex) {
	if _, ok := graph.Found2nd[n]; !ok {
		graph.Found2nd[n] = true
	}
}

// -----------------------------------------------------------------------------

func DFS1st(graph *Graph, start Vertex) {
	graph.addFound(start)
	neighbs := graph.GetNeighbours(start)
	for range neighbs {
		if !graph.inFound(neighbs[0]) {
			DFS1st(graph, neighbs[0])
		}
	}
	graph.Order = append(graph.Order, start)
}

func DFS2nd(graph *Graph, start, lead Vertex) {
	graph.addFound2nd(start)
	neighbs := graph.GetNeighbours2nd(start)
	graph.Components[lead][start] = true
	for range neighbs {
		if !graph.inFound2nd(neighbs[0]) {
			DFS2nd(graph, neighbs[0], lead)
		}
	}
}

func SCC(graph *Graph) {
	for i := range graph.Vertices {
		if graph.inFound(graph.Vertices[i]) {
			continue
		}
		graph.addFound(graph.Vertices[i])
		neighbs := graph.GetNeighbours(graph.Vertices[i])
		for j := range neighbs {
			if !graph.inFound(neighbs[j]) {
				DFS1st(graph, neighbs[j])
			}
		}
		graph.Order = append(graph.Order, graph.Vertices[i])
	}

	for k := range graph.Order {
		if !graph.inFound2nd(graph.Order[k]) {
			graph.addFound2nd(graph.Order[k])
			neighbs := graph.GetNeighbours2nd(graph.Order[k])
			graph.Components[graph.Order[k]] = make(map[Vertex]bool)
			for l := range neighbs {
				if !graph.inFound2nd(neighbs[l]) {
					DFS2nd(graph, neighbs[l], graph.Order[k])
				}
			}
		}
	}
}
