package TopologicalSort

import (
	"testing"
	"reflect"
)

func TestGraph_AddDirectedEdge(t *testing.T) {
	//Given we have a graph that matches graph0.dot.png
	gr := NewGraph()
	a := NewNode("a")
	b := NewNode("b")
	c := NewNode("c")
	d := NewNode("d")

	gr.AddNode(a)
	gr.AddNode(b)
	gr.AddNode(c)
	gr.AddNode(d)

	gr.AddDirectedEdge(a,b)
	gr.AddDirectedEdge(b,c)
	gr.AddDirectedEdge(b,d)

	// Then the graph will have the correct layout
	exp := map[Node]map[Node]bool{
		a: {
			b: true,
		},
		b: {
			c: true,
			d: true,
		},
		c: {},
		d: {},
	}

	if !reflect.DeepEqual(gr.Edges, exp) {
		t.Errorf("%v != %v", exp, gr.Edges)
	}
}

func TestTopologicalSort1(t *testing.T) {
	// Given there is a graph that matches graph0.dot.png
	gr := NewGraph()
	a := NewNode("a")
	b := NewNode("b")
	c := NewNode("c")
	d := NewNode("d")

	gr.AddNode(a)
	gr.AddNode(b)
	gr.AddNode(c)
	gr.AddNode(d)

	gr.AddDirectedEdge(a,b)
	gr.AddDirectedEdge(b,c)
	gr.AddDirectedEdge(b,d)

	// When we call the topological sort function
	TopologicalSort(&gr)

	// Then the expectant result should be one the 2 correct sorted lists
	exp1 := []Node{c,d,b,a}
	exp2 := []Node{d,c,b,a}

	if !reflect.DeepEqual(gr.Order, exp1) && !reflect.DeepEqual(gr.Order, exp2) {
		t.Errorf("%v \n!= \n %v \n|| \n%v", gr.Order, exp1, exp2)
	}
}

func TestTopologicalSort2(t *testing.T) {
	// Given there is a graph that matches graph1.dot.png
	gr := NewGraph()
	a := NewNode("a")
	b := NewNode("b")
	c := NewNode("c")
	d := NewNode("d")
	e := NewNode("e")
	f := NewNode("f")

	gr.AddNode(a)
	gr.AddNode(b)
	gr.AddNode(c)
	gr.AddNode(d)
	gr.AddNode(e)
	gr.AddNode(f)

	gr.AddDirectedEdge(a,b)
	gr.AddDirectedEdge(a,c)
	gr.AddDirectedEdge(b,d)
	gr.AddDirectedEdge(c,d)
	gr.AddDirectedEdge(c,e)
	gr.AddDirectedEdge(d,e)
	gr.AddDirectedEdge(d,f)
	gr.AddDirectedEdge(e,f)

	// When we call the topological sort function
	TopologicalSort(&gr)

	// Then the expectant result should be one the 2 correct sorted lists
	exp1 := []Node{f,e,d,b,c,a}
	exp2 := []Node{f,e,d,c,b,a}

	if !reflect.DeepEqual(gr.Order, exp1) && !reflect.DeepEqual(gr.Order, exp2) {
		t.Errorf("%v \n!= \n %v \n|| \n%v", gr.Order, exp1, exp2)
	}
}

func TestTopologicalSort3(t *testing.T) {
	// Given there is a graph that matches graph2.dot.png
	gr := NewGraph()
	a := NewNode("a")
	b := NewNode("b")
	c := NewNode("c")
	d := NewNode("d")
	e := NewNode("e")

	gr.AddNode(a)
	gr.AddNode(b)
	gr.AddNode(c)
	gr.AddNode(d)
	gr.AddNode(e)

	gr.AddDirectedEdge(a,b)
	gr.AddDirectedEdge(b,c)
	gr.AddDirectedEdge(b,d)
	gr.AddDirectedEdge(e,c)

	// When we call the topological sort function
	TopologicalSort(&gr)

	// Then the expectant result should be one the 2 correct sorted lists
	exp1 := []Node{d,c,b,a,e}
	exp2 := []Node{c,d,b,a,e}

	if !reflect.DeepEqual(gr.Order, exp1) && !reflect.DeepEqual(gr.Order, exp2) {
		t.Errorf("%v \n!= \n %v \n|| \n%v", gr.Order, exp1, exp2)
	}
}

func TestTopologicalSort4(t *testing.T) {
	// Given there is a graph that matches graph2.dot.png
	gr := NewGraph()
	a := NewNode("a")
	b := NewNode("b")
	c := NewNode("c")
	d := NewNode("d")
	e := NewNode("e")
	f := NewNode("f")
	g := NewNode("g")
	h := NewNode("h")

	gr.AddNode(a)
	gr.AddNode(b)
	gr.AddNode(c)
	gr.AddNode(d)
	gr.AddNode(e)
	gr.AddNode(f)
	gr.AddNode(g)
	gr.AddNode(h)

	gr.AddDirectedEdge(a,b)
	gr.AddDirectedEdge(a,e)
	gr.AddDirectedEdge(b,f)
	gr.AddDirectedEdge(e,c)
	gr.AddDirectedEdge(b,f)
	gr.AddDirectedEdge(c,f)
	gr.AddDirectedEdge(f,h)
	gr.AddDirectedEdge(d,h)

	// When we call the topological sort function
	TopologicalSort(&gr)

	// Then the expectant result should be the correct sorted list
	exp := []Node{h,f,b,c,e,a,d,g}

	if !reflect.DeepEqual(gr.Order, exp) {
		t.Errorf("%v != %v", gr.Order, exp)
	}
}
