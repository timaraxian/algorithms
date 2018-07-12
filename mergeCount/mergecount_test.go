package mergeCount

import (
	"testing"
	"reflect"
)

func TestSplitter_n4(t *testing.T) {
	input := []int{1,2,3,4}
	out1 := []int{1,2}
	out2 := []int{3,4}

	r1, r2 := Splitter(input, len(input))

	if !reflect.DeepEqual(r1, out1) {
		t.Errorf("%v != %v", out1, r1)
	}

	if !reflect.DeepEqual(r2, out2) {
		t.Errorf("%v != %v", out2, r2)
	}
}

func TestSplitter_n2(t *testing.T) {
	input := []int{1,2}
	out1 := []int{1}
	out2 := []int{2}

	r1, r2 := Splitter(input, len(input))

	if !reflect.DeepEqual(r1, out1) {
		t.Errorf("%v != %v", out1, r1)
	}

	if !reflect.DeepEqual(r2, out2) {
		t.Errorf("%v != %v", out2, r2)
	}
}

func TestSplitter_n3(t *testing.T) {
	input := []int{1,2,3}
	out1 := []int{1}
	out2 := []int{2,3}

	r1, r2 := Splitter(input, len(input))

	if !reflect.DeepEqual(r1, out1) {
		t.Errorf("%v != %v", out1, r1)
	}

	if !reflect.DeepEqual(r2, out2) {
		t.Errorf("%v != %v", out2, r2)
	}
}

func TestSplitter_n5(t *testing.T) {
	input := []int{1,2,3,4,5}
	out1 := []int{1,2}
	out2 := []int{3,4,5}

	r1, r2 := Splitter(input, len(input))

	if !reflect.DeepEqual(r1, out1) {
		t.Errorf("%v != %v", out1, r1)
	}

	if !reflect.DeepEqual(r2, out2) {
		t.Errorf("%v != %v", out2, r2)
	}
}

//func TestSwitch_2elements(t *testing.T) {
//	input := []int{2,1}
//	expA := []int{1,2}
//	expI := 1
//
//	resA, resI := Switch(input, 0)
//
//	if !reflect.DeepEqual(resA, expA) {
//		t.Errorf("%v != %v", expA, resA)
//	}
//
//	if expI != resI {
//		t.Errorf("%v != %v", expI, resI)
//	}
//}

//func TestMergeCount(t *testing.T) {
//	inp := []int{1,5,2,3,4,6}
//	out := []int{1,2,3,4,5,6}
//	exI := 3
//
//	res, reI := MergeCount(inp, 0)
//
//	if !reflect.DeepEqual(res, out) {
//		t.Errorf("%v != %v", out, res)
//	}
//
//	if exI != reI {
//		t.Errorf("%v != %v", exI, reI)
//	}
//}

func TestMergeCount_6backwards(t *testing.T) {
	inp := []int{6,5,4,3,2,1}
	out := []int{1,2,3,4,5,6}
	exI := 15

	res, reI := MergeCount(inp, 0)

	if !reflect.DeepEqual(res, out) {
		t.Errorf("%v != %v", out, res)
	}

	if exI != reI {
		t.Errorf("%v != %v", exI, reI)
	}
}

func TestMergeCount_febcda(t *testing.T) {
	inp := []int{6,5,2,3,4,1}
	out := []int{1,2,3,4,5,6}
	exI := 12

	res, reI := MergeCount(inp, 0)

	if !reflect.DeepEqual(res, out) {
		t.Errorf("%v != %v", out, res)
	}

	if exI != reI {
		t.Errorf("%v != %v", exI, reI)
	}
}

func TestMergeCount_badfec(t *testing.T) {
	inp := []int{2,1,4,6,5,3}
	out := []int{1,2,3,4,5,6}
	exI := 5

	res, reI := MergeCount(inp, 0)

	if !reflect.DeepEqual(res, out) {
		t.Errorf("%v != %v", out, res)
	}

	if exI != reI {
		t.Errorf("%v != %v", exI, reI)
	}
}

func TestMergeCount_7gcdeabf(t *testing.T) {
	inp := []int{7,3,4,5,1,2,6}
	out := []int{1,2,3,4,5,6,7}
	exI := 12

	res, reI := MergeCount(inp, 0)

	if !reflect.DeepEqual(res, out) {
		t.Errorf("%v != %v", out, res)
	}

	if exI != reI {
		t.Errorf("%v != %v", exI, reI)
	}
}
