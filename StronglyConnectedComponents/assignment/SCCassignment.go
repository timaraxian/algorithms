package main

import (
	"fmt"
	"sort"
	"os"
	"log"
	"bufio"
	"strings"
	"strconv"
)

type Graph struct {
	arcs map[int]map[int]bool
}

type Set struct {
	items map[int]bool
}

func (set *Set) Add(v int) {
	set.items[v] = true
}

func (set Set) Has(v int) bool {
	_, ok := set.items[v]
	return ok
}

type Stack struct {
	items []int
}

func (stack *Stack) Push(v int) {
	stack.items = append(stack.items, v)
}

func (stack *Stack) Pop() (v int, ok bool) {
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
		arcs: make(map[int]map[int]bool),
	}
}

func NewSet() Set {
	return Set{
		items: make(map[int]bool),
	}
}

func NewStack() Stack {
	return Stack{
		items: make([]int, 0),
	}
}

// -----------------------------------------------------------------------------

func (graph *Graph) AddArc(u, v int) {
	if _, ok := graph.arcs[u]; !ok {
		graph.arcs[u] = make(map[int]bool)
	}

	graph.arcs[u][v] = true

	if _, ok := graph.arcs[v]; !ok {
		graph.arcs[v] = make(map[int]bool)
	}
}

func (graph Graph) GetAdjacent(u int) (a []int) {
	for v := range graph.arcs[u] {
		a = append(a, v)
	}

	return a
}

// -----------------------------------------------------------------------------

func DFS(graph Graph) Stack {
	stack := NewStack()
	seen := NewSet()
	order := NewStack()
	for i := len(graph.arcs); i > 0; i-- {
		path := NewStack()
		if !seen.Has(i) {
			stack.Push(i)
		}
		seen.Add(i)
		for len(stack.items) > 0 {
			current, ok := stack.Pop()
			if !ok {
				panic(current)
			}
			path.Push(current)
			adjacent := graph.GetAdjacent(current)
			for j := range adjacent {
				if !seen.Has(adjacent[j]) {
					stack.Push(adjacent[j])
				}
				seen.Add(adjacent[j])
			}
		}
		for range path.items {
			v, ok := path.Pop()
			if !ok {
				panic("pathpop")
			}
			order.Push(v)
		}
	}
	return order
}

func SCC(graph Graph, order Stack) []int {
	stack := NewStack()
	seen := NewSet()
	orderLen := len(order.items)
	sccSize := make([]int, 0)
	for i := 0; i < orderLen; i++ {
		components := NewStack()
		leader, ok := order.Pop();
		if !ok {
			panic(leader)
		}
		if !seen.Has(leader) {
			stack.Push(leader)
		}
		seen.Add(leader)
		for len(stack.items) > 0 {
			current, ok := stack.Pop()
			if !ok {
				panic(current)
			}
			components.Push(current)
			adjacent := graph.GetAdjacent(current)
			for j := range adjacent {
				if !seen.Has(adjacent[j]) {
					stack.Push(adjacent[j])
				}
				seen.Add(adjacent[j])
			}
		}
		sccSize = append(sccSize, len(components.items))
	}
	return sccSize
}

// -----------------------------------------------------------------------------

func main() {
	graph := NewGraph()

	fmt.Println("Opening file...")
	file, err := os.Open("SCC.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Println("Reading file...")
	scanner := bufio.NewReader(file)
	for {
		line, _, err := scanner.ReadLine()
		if err != nil {
			break
		}
		nodes := strings.Split(strings.Trim(string(line), " "), " ")
		if len(nodes) != 2 {
			panic(len(nodes))
		}

		u, err := strconv.Atoi(nodes[0])
		if err != nil {
			panic(nodes[0])
		}
		v, err := strconv.Atoi(nodes[1])
		if err != nil {
			panic(nodes[1])
		}
		graph.AddArc(u, v)
	}

	fmt.Println("Reversing graph...")
	revGraph := NewGraph()
	for u, uMap := range graph.arcs {
		for v := range uMap {
			revGraph.AddArc(v, u)
		}
	}

	fmt.Println("DFS on reversed graph...")
	order := DFS(revGraph)

	fmt.Println("DFS on graph...")
	sizes := SCC(graph, order)

	fmt.Println("Sorting sizes...")
	sort.Ints(sizes)

	fmt.Println(sizes[len(sizes)-5:])
}
