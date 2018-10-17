package primMST

import "testing"

func TestPrimMST(t *testing.T) {
	g := NewGraph()

	g.AddArc(1, 2, 2)
	g.AddArc(1, 3, 1)
	g.AddArc(1, 4, 2)
	g.AddArc(2, 3, 3)
	g.AddArc(3, 4, 4)

	result := PrimMST(g)
	expected := 5

	if result != expected {
		t.Fatalf("%v != %v", result, expected)
	}
}

func TestPrimMST2(t *testing.T) {
	g := NewGraph()

	g.AddArc(1, 2, 3)
	g.AddArc(1, 3, 1)
	g.AddArc(2, 3, 1)
	g.AddArc(3, 4, 3)
	g.AddArc(3, 5, 1)
	g.AddArc(3, 6, 2)
	g.AddArc(4, 6, 4)
	g.AddArc(4, 5, 1)

	result := PrimMST(g)
	expected := 6

	if result != expected {
		t.Fatalf("%v != %v", result, expected)
	}
}
