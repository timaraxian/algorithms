package DFS

import (
	"testing"
)

func TestDFS_d(t *testing.T) {
	//Given there is a graph with the following configuration
	g := NewGraph()

	s := NewNode("s")
	a := NewNode("a")
	b := NewNode("b")
	c := NewNode("c")
	d := NewNode("d")
	e := NewNode("e")

	g.AddEdge(s, a)
	g.AddEdge(s, b)
	g.AddEdge(a, c)
	g.AddEdge(b, c)
	g.AddEdge(b, d)
	g.AddEdge(c, d)
	g.AddEdge(c, e)
	g.AddEdge(d, e)

	// When DFS is called with node 'd' being the search node
	result := DFS(g, s, d)
	// Then 'true' should be returned
	if ! result {
		t.Errorf("%v != %v", true, result)
	}
}

func TestDFS_f(t *testing.T) {
	//Given there is a graph with the following configuration
	g := NewGraph()

	s := NewNode("s")
	a := NewNode("a")
	b := NewNode("b")
	c := NewNode("c")
	d := NewNode("d")
	e := NewNode("e")
	f := NewNode("f")

	g.AddEdge(s, a)
	g.AddEdge(s, b)
	g.AddEdge(a, c)
	g.AddEdge(b, c)
	g.AddEdge(b, d)
	g.AddEdge(c, d)
	g.AddEdge(c, e)
	g.AddEdge(d, e)

	// and When DFS is called on node "f" not connected to the graph
	result := DFS(g, s, f)

	// Then the result will be false
	if result {
		t.Errorf("%v != %v", false, result)
	}
}

// -----------------------------------------------------------------------------
// Helper functions
// -----------------------------------------------------------------------------
func neighborsEqual(a []Node, b []Node) (e bool) {
	if len(a) != len(b) {
		return false
	}

	aMap := map[Node]bool{}
	bMap := map[Node]bool{}
	for i := range a {
		aMap[a[i]] = true
		bMap[b[i]] = true
	}

	for i := range b {
		if _, ok := aMap[b[i]]; !ok {
			return false
		}
		if _, ok := bMap[a[i]]; !ok {
			return false
		}
	}

	return true
}
