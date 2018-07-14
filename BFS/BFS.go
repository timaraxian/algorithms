package BFS

type Node struct {
	Value string
}

type Graph struct {
	Edges    map[Node]map[Node]bool
	Found    []Node
	Explored map[Node]int
}

type Queue struct {
	List []Node
}

// -----------------------------------------------------------------------------

func NewNode(value string) Node {
	return Node{
		Value: value,
	}
}

func NewGraph() Graph {
	return Graph{
		Edges:    make(map[Node]map[Node]bool),
		Explored: make(map[Node]int),
	}
}

func NewQueue() Queue {
	return Queue{
		List: []Node{},
	}
}

// -----------------------------------------------------------------------------

func (graph *Graph) AddNode(n Node) {
	if _, ok := graph.Edges[n]; !ok {
		graph.Edges[n] = make(map[Node]bool)
	}
}

func (graph *Graph) AddEdge(n1, n2 Node) {
	if _, ok := graph.Edges[n1]; !ok {
		graph.Edges[n1] = make(map[Node]bool)
	}

	if _, ok := graph.Edges[n1][n2]; !ok {
		graph.Edges[n1][n2] = true
	}

	if _, ok := graph.Edges[n2]; !ok {
		graph.Edges[n2] = make(map[Node]bool)
	}

	if _, ok := graph.Edges[n2][n1]; !ok {
		graph.Edges[n2][n1] = true
	}
}

func (graph *Graph) GetNodes() []Node {
	var nodes []Node

	for key := range graph.Edges {
		nodes = append(nodes, key)
	}

	return nodes
}

func (graph *Graph) GetNeighbours(n Node) []Node {
	var nodes []Node

	for value := range graph.Edges[n] {
		nodes = append(nodes, value)
	}

	return nodes
}

func (graph *Graph) inFound(n Node) bool {
	for f := range graph.Found {
		if graph.Found[f] == n {
			return true
		}
	}
	graph.Found = append(graph.Found, n)

	return false
}

func (graph *Graph) addExplored(n Node, dist int) {
	if _, ok := graph.Explored[n]; !ok {
		graph.Explored[n] = dist
	}
}

func (graph *Graph) inExplored(prev, n Node) bool {
	for e := range graph.Explored {
		if e == n {
			return true
		}
	}
	dist := graph.Explored[prev] + 1
	graph.addExplored(n, dist)

	return false
}

// -----------------------------------------------------------------------------

func (queue *Queue) Pop() Node {
	n := queue.List[0]
	queue.List = queue.List[1:]

	return n
}

func (queue *Queue) Push(n Node) {
	queue.List = append(queue.List, n)
}

// -----------------------------------------------------------------------------

func BFS(graph Graph, s, search Node) bool {
	// add node s to the queue and mark as found
	graph.Found = append(graph.Found, s)
	q := NewQueue()
	q.Push(s)

	//while the q isn't empty
	for len(q.List) > 0 {
		// get the first element in the queue
		n1 := q.Pop()
		neighbs := graph.GetNeighbours(n1)
		// set all the unfound nodes as found and onto the queue
		for i := range neighbs {
			if search == neighbs[i] {
				return true
			}
			if ! graph.inFound(neighbs[i]) {
				q.Push(neighbs[i])
			}
		}
	}
	return false
}

// -----------------------------------------------------------------------------

func BFSdist(graph Graph, s, search Node) (bool, int) {
	graph.addExplored(s, 0)
	q := NewQueue()
	q.Push(s)

	//base case s = search
	if s == search {
		return true, graph.Explored[s]
	}

	for len(q.List) > 0 {
		n1 := q.Pop()
		neighbs := graph.GetNeighbours(n1)

		for i := range neighbs {
			if ! graph.inExplored(n1, neighbs[i]) {
				q.Push(neighbs[i])
			}
			if search.Value == neighbs[i].Value {
				return true, graph.Explored[neighbs[i]]
			}
		}
	}
	return false, 0
}
