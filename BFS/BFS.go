package BFS

type Node struct {
	Value string
	Dist  int
}

type Graph struct {
	Edges map[Node]map[Node]bool
	Found []string
}

type Queue struct {
	List []Node
}

// -----------------------------------------------------------------------------

func NewNode(value string) Node {
	return Node{
		Value: value,
		Dist:  0,
	}
}

func NewGraph() Graph {
	return Graph{
		Edges: make(map[Node]map[Node]bool),
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

func (graph *Graph) inFound(value string) bool {
	for f := range graph.Found {
		if graph.Found[f] == value {
			return true
		}
	}
	graph.Found = append(graph.Found, value)

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
	graph.Found = append(graph.Found, s.Value)
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
			if ! graph.inFound(neighbs[i].Value) {
				q.Push(neighbs[i])
			}
		}
	}
	return false
}

// -----------------------------------------------------------------------------

func BFSdist(graph Graph, s, search Node) (bool, int) {
	graph.Found = append(graph.Found, s.Value)
	q := NewQueue()
	q.Push(s)
	//base case s = search
	if s == search {
		return true, s.Dist
	}

	for len(q.List) > 0 {
		n1 := q.Pop()
		neighbs := graph.GetNeighbours(n1)

		for i := range neighbs {
			neighbs[i].Dist = n1.Dist + 1
			if search.Value == neighbs[i].Value {
				return true, neighbs[i].Dist
			}
			if ! graph.inFound(neighbs[i].Value) {
				q.Push(neighbs[i])
			}
		}
	}
	return false, 0
}
