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

	if !reflect.DeepEqual(g.edges, exp) {
		t.Errorf("%v != %v", exp, g.edges)
	}
}

func TestGraph_AddNode(t *testing.T) {
	g := NewGraph()
	s := NewNode("s")

	g.AddNode(s)

	exp := map[Node]map[Node]bool{
		s: {},
	}

	if !reflect.DeepEqual(g.edges, exp) {
		t.Errorf("%v != %v", exp, g.edges)
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

	if !neighborsEqual(q.list, exp1) {
		t.Errorf("%v != %v", exp1, q.list)
	}

	result := q.Pop()
	exp2 := a

	if result != exp2 {
		t.Errorf("%v != %v", exp2, result)
	}
}

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
