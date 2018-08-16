package main

import (
	"sort"
	"fmt"
	"os"
	"strings"
	"strconv"
	"bufio"
)

type Graph struct {
	arcs map[int]map[int]int
}

// -----------------------------------------------------------------------------

type Set struct {
	items map[int]bool
}

func (set *Set) Add(i int) {
	set.items[i] = true
}

func (set Set) Has(i int) bool {
	_, ok := set.items[i]
	return ok
}

// -----------------------------------------------------------------------------

type ArcStack struct {
	items [][]int
}

func (stack *ArcStack) Push(v []int) {
	stack.items = append(stack.items, v)
}

func (stack *ArcStack) Pop() (v []int, ok bool) {
	if len(stack.items) == 0 {
		return v, false
	}
	v = stack.items[len(stack.items)-1]
	stack.items = stack.items[:len(stack.items)-1]
	return v, true
}

// -----------------------------------------------------------------------------

func NewGraph() Graph {
	return Graph{
		arcs: make(map[int]map[int]int),
	}
}

func NewSet() Set {
	return Set{
		items: make(map[int]bool),
	}
}

func NewArcStack() ArcStack {
	return ArcStack{
		items: make([][]int, 0),
	}
}

// -----------------------------------------------------------------------------

func (graph *Graph) AddArc(u, v, l int) {
	if _, ok := graph.arcs[u]; !ok {
		graph.arcs[u] = make(map[int]int)
	}

	graph.arcs[u][v] = l

	if _, ok := graph.arcs[v]; !ok {
		graph.arcs[v] = make(map[int]int)
	}
}

func (graph Graph) GetArcs(u int) (a [][]int) {
	for v, d := range graph.arcs[u] {
		a = append(a, []int{u, v, d})
	}

	return a
}

// -----------------------------------------------------------------------------

func Dijkstra(graph Graph, start int) (A map[int]int, B map[int][]int) {
	A = make(map[int]int)
	B = make(map[int][]int)

	outArcs := [][]int{{start, start, 0}}
	seen := NewSet()
	for {
		if len(outArcs) == 0 {
			break
		}
		current := outArcs[0]
		seen.Add(current[1])

		A[current[1]] = A[current[0]] + current[2]

		B[current[1]] = make([]int, len(B[current[0]]))
		copy(B[current[1]], B[current[0]])
		B[current[1]] = append(B[current[1]], current[1])

		adj := graph.GetArcs(current[1])

		for _, arc := range adj {
			outArcs = append(outArcs, arc)
		}

		tempSlice := make([][]int, 0)
		for _, outArc := range outArcs {
			if !seen.Has(outArc[1]) {
				tempSlice = append(tempSlice, outArc)
			}
		}

		sort.Slice(tempSlice, func(i, j int) bool {
			if A[tempSlice[i][0]]+tempSlice[i][2] < A[tempSlice[j][0]]+tempSlice[j][2] {
				return true
			}
			return false
		})
		outArcs = tempSlice
	}

	return A, B
}

// -----------------------------------------------------------------------------

func main() {
	graph := NewGraph()

	fmt.Println("opening file...")
	file, err := os.Open("dijkstra.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Println("Reading file...")
	reader := bufio.NewReader(file)
	fmt.Println("Building graph...")
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}

		vertex := strings.Fields(strings.Trim(string(line), " "))
		if len(vertex) == 0 {
			panic(vertex)
		}
		u, err := strconv.Atoi(vertex[0])
		if err != nil {
			panic(u)
		}

		for _, tuple := range vertex[1:] {
			vDist := strings.Split(tuple, ",")
			if len(vDist) != 2 {
				panic(vDist)
			}
			v, err := strconv.Atoi(vDist[0])
			if err != nil {
				panic(v)
			}
			d, err := strconv.Atoi(vDist[1])
			if err != nil {
				panic(v)
			}
			graph.AddArc(u, v, d)
		}
	}

	fmt.Println("Running algorithm...")
	A, _ := Dijkstra(graph, 1)

	requiredPaths := []int{A[7], A[37], A[59], A[82], A[99], A[115], A[133], A[165], A[188], A[197]}
	fmt.Println(requiredPaths)
}
