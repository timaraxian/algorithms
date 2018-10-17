package primMST

import "sort"

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
		a = append(a, []int{u, v, w})
	}

	return a
}

// -----------------------------------------------------------------------------

type Set struct {
	items map[int]bool
}

func NewSet() Set {
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

func PrimMST(g Graph) int {
	explored := NewSet()
	dist := 0

	var start int
	for u := range g.arcs {
		start = u
		break
	}

	explored.Add(start)

	candidates := [][]int{{start, start, 0}}

	for range g.arcs {
		if len(explored.items) == len(g.arcs) {
			break
		}

		current := candidates[0]
		explored.Add(current[1])
		dist += current[2]

		adj := g.GetArcs(current[1])

		for _, arc := range adj {
			candidates = append(candidates, arc)
		}

		tempSlice := make([][]int, 0)
		for _, arc := range candidates {
			if !explored.Has(arc[1]) {
				tempSlice = append(tempSlice, arc)
			}
		}

		sort.Slice(tempSlice, func(i, j int) bool {
			if tempSlice[i][2] < tempSlice[j][2] {
				return true
			}
			return false
		})
		candidates = tempSlice
	}

	return dist
}
