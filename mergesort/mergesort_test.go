package algorithms

import (
	"testing"
	"reflect"
	"math/rand"
	"time"
	"sort"
)

func TestMergeSort_baseCase(t *testing.T) {
	input := make([]int, 0)
	expected := make([]int, 0)

	if result := MergeSort(input); !reflect.DeepEqual(result, expected) {
		t.Errorf("%v != %v", expected, result)
	}
}

func TestMergeSort_n1(t *testing.T) {
	input := []int{8}
	expected := []int{8}

	if result := MergeSort(input); !reflect.DeepEqual(result, expected) {
		t.Errorf("%v != %v", expected, result)
	}
}

func TestMergeSort_n2(t *testing.T) {
	input := []int{8, 6}
	expected := []int{6, 8}

	if result := MergeSort(input); !reflect.DeepEqual(result, expected) {
		t.Errorf("%v != %v", expected, result)
	}

	input = []int{6, 8}
	expected = []int{6, 8}

	if result := MergeSort(input); !reflect.DeepEqual(result, expected) {
		t.Errorf("%v != %v", expected, result)
	}
}

func TestMergeSort_n3(t *testing.T) {

	input := []int{10, 6, 8}
	expected := []int{6, 8, 10}

	if result := MergeSort(input); !reflect.DeepEqual(result, expected) {
		t.Errorf("%v != %v", expected, result)
	}

	input = []int{8, 6, 10}
	expected = []int{6, 8, 10}

	if result := MergeSort(input); !reflect.DeepEqual(result, expected) {
		t.Errorf("%v != %v", expected, result)
	}
}

func TestMergeSort_ndick(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	sorted := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		sorted[i] = i
	}

	shuffled := make([]int, len(sorted))
	perm := rand.Perm(len(sorted))
	for i, v := range perm {
		shuffled[v] = sorted[i]
	}
	if result := MergeSort(shuffled); !reflect.DeepEqual(result, sorted) {
		t.Errorf("%v != %v", sorted, result)
	}
}

func TestMergeSort_ndickbubble(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	sorted := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		sorted[i] = i
	}

	shuffled := make([]int, len(sorted))
	perm := rand.Perm(len(sorted))
	for i, v := range perm {
		shuffled[v] = sorted[i]
	}
	if result := bubble(shuffled); !reflect.DeepEqual(result, sorted) {
		t.Errorf("%v != %v", sorted, result)
	}
}

func BenchmarkMergeSort(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		b.StopTimer()
		r := randSlice(10000)
		b.StartTimer()
		MergeSort(r)
	}
}

func BenchmarkMergeSort2(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		b.StopTimer()
		r := randSlice(10000)
		b.StartTimer()
		sort.Ints(r)
	}
}

func BenchmarkMergeSort3(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		b.StopTimer()
		r := randSlice(10000)
		b.StartTimer()
		bubble(r)
	}
}


func randSlice(n int) []int {
	return  rand.Perm(n)
}

func bubble(y []int) []int {
	x := make([]int, len(y))
	copy(x, y)
	end := len(x) - 1

	for {

		if end == 0 {
			break
		}

		for i := 0; i < len(x)-1; i++ {

			if x[i] > x[i+1] {
				x[i], x[i+1] = x[i+1], x[i]
			}

		}

		end -= 1

	}
	return x
}
