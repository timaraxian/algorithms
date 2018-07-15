package BFS

import (
	"testing"
	"reflect"
)

func TestGraph_AddEdge(t *testing.T) {
	g := NewGraph()
	s := NewNode("s")
	a := NewNode("a")

	g.AddEdge(s, a)
	exp := map[Node]map[Node]bool{
		s: {
			a: true,
		},
		a: {
			s: true,
		},
	}

	if !reflect.DeepEqual(g.Edges, exp) {
		t.Errorf("%v != %v", exp, g.Edges)
	}
}

func TestGraph_AddNode(t *testing.T) {
	g := NewGraph()
	s := NewNode("s")

	g.AddNode(s)

	exp := map[Node]map[Node]bool{
		s: {},
	}

	if !reflect.DeepEqual(g.Edges, exp) {
		t.Errorf("%v != %v", exp, g.Edges)
	}
}

func TestGraph_GetNodes(t *testing.T) {
	g := NewGraph()
	s := NewNode("s")
	a := NewNode("a")
	b := NewNode("b")
	c := NewNode("c")

	g.AddEdge(s, a)
	g.AddEdge(s, b)
	g.AddNode(c)

	exp := []Node{s, a, b, c}

	if !neighborsEqual(g.GetNodes(), exp) {
		t.Errorf("%v != %v", exp, g.GetNodes())
	}
}

func TestGraph_GetNeighbours(t *testing.T) {
	g := NewGraph()
	s := NewNode("s")
	a := NewNode("a")
	b := NewNode("b")

	g.AddEdge(s, a)
	g.AddEdge(s, b)

	exp := []Node{a, b}

	if !neighborsEqual(g.GetNeighbours(s), exp) {
		t.Errorf("%v != %v", exp, g.GetNeighbours(s))
	}
}

func TestQueue(t *testing.T) {
	q := NewQueue()
	a := NewNode("a")
	b := NewNode("b")
	c := NewNode("c")
	d := NewNode("d")

	q.Push(a)
	q.Push(b)
	q.Push(c)
	q.Push(d)

	exp1 := []Node{a, b, c, d}

	if !neighborsEqual(q.List, exp1) {
		t.Errorf("%v != %v", exp1, q.List)
	}

	result := q.Pop()
	exp2 := a

	if result != exp2 {
		t.Errorf("%v != %v", exp2, result)
	}
}

func TestBFS(t *testing.T) {
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

	// When BFS is called with node 'd' being the search node
	resultTrue := BFS(g, s, d)
	// Then 'true' should be returned
	if ! resultTrue {
		t.Errorf("%v != %v", true, resultTrue)
	}

	// and When there is a node 'f' which is the search node
	f := NewNode("f")
	resultFalse := BFS(g, s, f)

	// Then 'false' should be returned
	if resultFalse {
		t.Errorf("%v != %v", false, resultFalse)
	}
}

func TestBFSdist_s(t *testing.T) {
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

	// When BFSdist is called on node 's'
	conx, dist := BFSdist(g, s, s)

	// Then the distance returned should be 0
	exConx, exDist := true, 0
	if ! conx {
		t.Errorf("%v != %v", true, exConx)
	}
	if !reflect.DeepEqual(dist, exDist) {
		t.Errorf("%v != %v", exDist, dist)
	}
}

func TestBFSdist_a(t *testing.T) {
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

	// When BFSdist is called on node 'a'
	conx, dist := BFSdist(g, s, a)

	// Then the distance returned should be 0
	exConx, exDist := true, 1
	if ! conx {
		t.Errorf("%v != %v", true, exConx)
	}
	if !reflect.DeepEqual(dist, exDist) {
		t.Errorf("%v != %v", exDist, dist)
	}
}

func TestBFSdist_b(t *testing.T) {
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

	// When BFSdist is called on node 'a'
	conx, dist := BFSdist(g, s, b)

	// Then the distance returned should be 0
	exConx, exDist := true, 1
	if ! conx {
		t.Errorf("%v != %v", true, exConx)
	}
	if !reflect.DeepEqual(dist, exDist) {
		t.Errorf("%v != %v", exDist, dist)
	}
}

func TestBFSdist_c(t *testing.T) {
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

	// When BFSdist is called on node 'a'
	conx, dist := BFSdist(g, s, c)

	// Then the distance returned should be 0
	exConx, exDist := true, 2
	if ! conx {
		t.Errorf("%v != %v", true, exConx)
	}
	if !reflect.DeepEqual(dist, exDist) {
		t.Errorf("%v != %v", exDist, dist)
	}
}

func TestBFSdist_d(t *testing.T) {
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

	// When BFSdist is called on node 'a'
	conx, dist := BFSdist(g, s, d)

	// Then the distance returned should be 0
	exConx, exDist := true, 2
	if ! conx {
		t.Errorf("%v != %v", true, exConx)
	}
	if !reflect.DeepEqual(dist, exDist) {
		t.Errorf("%v != %v", exDist, dist)
	}
}

func TestBFSdist_e(t *testing.T) {
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

	// When BFSdist is called on node 'a'
	conx, dist := BFSdist(g, s, e)

	// Then the distance returned should be 0
	exConx, exDist := true, 3
	if ! conx {
		t.Errorf("%v != %v", true, exConx)
	}
	if !reflect.DeepEqual(dist, exDist) {
		t.Errorf("%v != %v", exDist, dist)
	}
}

func TestBFSdist_f(t *testing.T) {
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

	// When BFSdist is called on node 'a'
	conx, dist := BFSdist(g, s, f)

	// Then the distance returned should be 0
	exConx, exDist := false, 0
	if conx {
		t.Errorf("%v != %v", false, exConx)
	}
	if !reflect.DeepEqual(dist, exDist) {
		t.Errorf("%v != %v", exDist, dist)
	}
}

func TestBFSconnectivity(t *testing.T) {
	// Given there is a graph with the following configuration
	graph := NewGraph()

	a := NewNode("a")
	b := NewNode("b")
	c := NewNode("c")
	d := NewNode("d")
	e := NewNode("e")
	f := NewNode("f")
	g := NewNode("g")
	h := NewNode("h")
	i := NewNode("i")
	j := NewNode("j")

	graph.AddEdge(a, c)
	graph.AddEdge(a, e)
	graph.AddEdge(c, e)
	graph.AddEdge(e, g)
	graph.AddEdge(e, i)

	graph.AddEdge(b, d)

	graph.AddEdge(f, h)
	graph.AddEdge(f, j)
	graph.AddEdge(h, j)

	// When we call BFSconnectivity on the graph
	result := BFSconnectivity(graph)
	// Then we should expect the connected nodes returned
	exResult := [][]Node{
		{a, c, e, g, i},
		{b, d},
		{f, h, j},
	}

	if !connectivityEqual(result, exResult) {
		t.Errorf("%v != %v", exResult, result)
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

func connectivityEqual(result, expected [][]Node) bool {
	for i := range result {
		for j := range expected {
			if neighborsEqual(result[i], expected[j]) {
				continue
			}
		}
	}
	return true
}
