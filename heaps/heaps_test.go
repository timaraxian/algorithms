package heaps

import (
	"testing"
	"reflect"
)

func TestNewHeap(t *testing.T) {
	h := NewHeap()

	exp := make([]int,0)

	if !reflect.DeepEqual(exp, h.Items) {
		t.Errorf("%v != %v", exp, h.Items)
	}
}

func TestHeap_Insert(t *testing.T) {
	h := NewHeap()

	h.Insert(1)

	exp := []int{1}
	if !reflect.DeepEqual(exp, h.Items) {
		t.Errorf("%v != %v", exp, h.Items)
	}

	h.Insert(5)

	exp = []int{1,5}
	if !reflect.DeepEqual(exp, h.Items) {
		t.Errorf("%v != %v", exp, h.Items)
	}

	h.Insert(3)

	exp = []int{1,5,3}
	if !reflect.DeepEqual(exp, h.Items) {
		t.Errorf("%v != %v", exp, h.Items)
	}
}

func TestHeap_Insert2(t *testing.T) {
	h := NewHeap()

	h.Insert(1)
	exp := []int{1}
	if !reflect.DeepEqual(exp, h.Items) {
		t.Errorf("%v != %v", exp, h.Items)
	}

	h.Insert(1)
	exp = []int{1,1}
	if !reflect.DeepEqual(exp, h.Items) {
		t.Errorf("%v != %v", exp, h.Items)
	}

	h.Insert(13)
	exp = []int{1,1,13}
	if !reflect.DeepEqual(exp, h.Items) {
		t.Errorf("%v != %v", exp, h.Items)
	}

	h.Insert(3)
	exp = []int{1,1,13,3}
	if !reflect.DeepEqual(exp, h.Items) {
		t.Errorf("%v != %v", exp, h.Items)
	}

	h.Insert(8)
	exp = []int{1,1,13,3,8}
	if !reflect.DeepEqual(exp, h.Items) {
		t.Errorf("%v != %v", exp, h.Items)
	}

	h.Insert(5)
	exp = []int{1,1,5,3,8,13}
	if !reflect.DeepEqual(exp, h.Items) {
		t.Errorf("%v != %v", exp, h.Items)
	}

	h.Insert(2)
	exp = []int{1,1,2,3,8,13,5}
	if !reflect.DeepEqual(exp, h.Items) {
		t.Errorf("%v != %v", exp, h.Items)
	}

	h.Insert(21)
	exp = []int{1,1,2,3,8,13,5,21}
	if !reflect.DeepEqual(exp, h.Items) {
		t.Errorf("%v != %v", exp, h.Items)
	}
}

func TestHeap_ExtractMin(t *testing.T) {
	h := NewHeap()

	h.Insert(1)
	h.Insert(1)
	h.Insert(13)
	h.Insert(3)
	h.Insert(8)
	h.Insert(5)
	h.Insert(2)
	h.Insert(21)

	min, ok := h.ExtractMin()
	exp := 1
	if min != exp {
		t.Errorf("%v != %v", min, exp)
	}
	if !ok {
		t.Errorf("ok != true")
	}
	
	min, ok = h.ExtractMin()
	exp = 1
	if min != exp {
		t.Errorf("%v != %v", min, exp)
	}

	min, ok = h.ExtractMin()
	exp = 2
	if min != exp {
		t.Errorf("%v != %v", min, exp)
	}

	min, ok = h.ExtractMin()
	exp = 3
	if min != exp {
		t.Errorf("%v != %v", min, exp)
	}

	min, ok = h.ExtractMin()
	exp = 5
	if min != exp {
		t.Errorf("%v != %v", min, exp)
	}

	min, ok = h.ExtractMin()
	exp = 8
	if min != exp {
		t.Errorf("%v != %v", min, exp)
	}

	min, ok = h.ExtractMin()
	exp = 13
	if min != exp {
		t.Errorf("%v != %v", min, exp)
	}

	min, ok = h.ExtractMin()
	exp = 21
	if min != exp {
		t.Errorf("%v != %v", min, exp)
	}

	min, ok = h.ExtractMin()
	if ok {
		t.Errorf("ok != false")
	}
}
