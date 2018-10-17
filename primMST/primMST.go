package primMST

type Graph struct {
	arcs map[int]map[int]int
}

func NewGraph() Graph {
	return Graph{
		arcs: make(map[int]map[int]int),
	}
}

func (g *Graph) AddArc(u, v, w int) {
	if _, ok := g.arcs[u]; !ok {
		g.arcs[u] = make(map[int]int)
	}

	g.arcs[u][v] = w

	if _, ok := g.arcs[v]; !ok {
		g.arcs[v] = make(map[int]int)
	}

	g.arcs[v][u] = w
}

func (g Graph) GetArcs(u int) (a [][]int) {
	for v, w := range g.arcs[u] {
		a = append(a, []int{u,v,w})
	}

	return a
}

// -----------------------------------------------------------------------------

type Set struct {
	items map[int]bool
}

func NewSet() Set{
	return Set{
		items: make(map[int]bool),
	}
}

func (s *Set) Add(i int) {
	s.items[i] = true
}

func (s Set) Has(i int) bool {
	_, ok := s.items[i]

	return ok
}

// -----------------------------------------------------------------------------

func primMST(g Graph) int {
	explored := NewSet()
	dist := 0

	var start int
	for u := range g.arcs {
		start = u
		break
	}

	explored.Add(start)

	for range g.arcs {
		if len(explored.items) == len(g.arcs) {
			break
		}

		
	}
}
