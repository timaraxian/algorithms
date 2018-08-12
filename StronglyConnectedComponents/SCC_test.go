package StronglyConnectedComponents

import (
	"testing"
	"reflect"
)

func TestSCC(t *testing.T) {
	// Given there is a graph like shown in scc.dot.pnh
	gr := NewGraph()
	a := NewVertex("a")
	b := NewVertex("b")
	c := NewVertex("c")
	d := NewVertex("d")
	e := NewVertex("e")
	f := NewVertex("f")
	g := NewVertex("g")
	h := NewVertex("h")
	i := NewVertex("i")

	gr.AddVertex(a)
	gr.AddVertex(b)
	gr.AddVertex(c)
	gr.AddVertex(d)
	gr.AddVertex(e)
	gr.AddVertex(f)
	gr.AddVertex(g)
	gr.AddVertex(h)
	gr.AddVertex(i)

	gr.AddArc(a,d)
	gr.AddArc(d,g)
	gr.AddArc(g,a)
	gr.AddArc(i,g)
	gr.AddArc(i,c)
	gr.AddArc(c,f)
	gr.AddArc(f,i)
	gr.AddArc(h,f)
	gr.AddArc(h,b)
	gr.AddArc(b,e)
	gr.AddArc(e,h)

	// When SCC is called on the graph
	SCC(&gr)

	// Then the correct Strongly Connected Components should be returned
	exp := map[Vertex]map[Vertex]bool{
		g: {
			a: true,
			d: true,
		},
		c: {
			i: true,
			f: true,
		},
		h: {
			b: true,
			e: true,
		},
	}

	if !reflect.DeepEqual(gr.Components, exp) {
		t.Errorf("\n%v \n!=\n %v", exp, gr.Components)
	}
}
