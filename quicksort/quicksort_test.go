package quicksort

import (
	"testing"
	"reflect"
)

func TestQuickSort(t *testing.T) {
	input := []int{3,8,2,5,1,4,7,6}
	expected := []int{1,2,3,4,5,6,7,8}

	result, _ := QuickSort(input,len(input), 0)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("%v != %v", expected, result)
	}
}

func TestQuickSort_2(t *testing.T) {
	input := []int{3,8,1,5,2,4,7,6}
	expected := []int{1,2,3,4,5,6,7,8}

	result, _ := QuickSort(input, len(input), 0)

	if  !reflect.DeepEqual(result, expected) {
		t.Errorf("%v != %v", expected, result)
	}
}

func TestQuickSort_3(t *testing.T) {
	input := []int{5,7,1,3,9,2,4,8,6}
	expected := []int{1,2,3,4,5,6,7,8,9}

	result, _ := QuickSort(input, len(input), 0)

	if  !reflect.DeepEqual(result, expected) {
		t.Errorf("%v != %v", expected, result)
	}
}

func TestQuickSort_4(t *testing.T) {
	input := []int{1,1,5,3,9,2,7,4,8,6,1}
	expected := []int{1,1,1,2,3,4,5,6,7,8,9}

	result, _ := QuickSort(input, len(input), 0)

	if  !reflect.DeepEqual(result, expected) {
		t.Errorf("%v != %v", expected, result)
	}
}

func TestQuickSort_5(t *testing.T) {
	input := []int{3,2,5,1,4}
	expected := []int{1,2,3,4,5}
	expComp := 6

	result, comp := QuickSort(input, len(input), 0)

	if comp != expComp {
		t.Errorf("%v != %v", expComp, comp)
	}

	if  !reflect.DeepEqual(result, expected) {
		t.Errorf("%v != %v", expected, result)
	}
}

func TestQuickSort_6(t *testing.T) {
	input := []int{1,2,5,3,4}
	expected := []int{1,2,3,4,5}
	expComp := 10

	result, comp := QuickSort(input, len(input), 0)

	if comp != expComp {
		t.Errorf("%v != %v", expComp, comp)
	}

	if  !reflect.DeepEqual(result, expected) {
		t.Errorf("%v != %v", expected, result)
	}
}

func TestQuickSort_7(t *testing.T) {
	input := []int{2,1,3,4,5}
	expected := []int{1,2,3,4,5}
	expComp := 7

	result, comp := QuickSort(input, len(input), 0)

	if comp != expComp {
		t.Errorf("%v != %v", expComp, comp)
	}

	if  !reflect.DeepEqual(result, expected) {
		t.Errorf("%v != %v", expected, result)
	}
}
