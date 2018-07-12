package BFS

type Node struct {
	Value string
}

type Graph struct {
	Edges map[Node]map[Node]bool
	Found []Node
}

type Queue struct {
	List []Node
}

// -----------------------------------------------------------------------------

func NewNode(value string) Node {
	return Node{
		value: value,
	}
}

func NewGraph() Graph {
	return Graph{
		edges: make(map[Node]map[Node]bool),
	}
}

func NewQueue() Queue {
	return Queue{
		list: []Node{},
	}
}

// -----------------------------------------------------------------------------

func (graph *Graph) AddNode(n Node) {
	if _, ok := graph.edges[n]; !ok {
		graph.edges[n] = make(map[Node]bool)
	}
}

func (graph *Graph) AddEdge(n1, n2 Node) {
	if _, ok := graph.edges[n1]; !ok {
		graph.edges[n1] = make(map[Node]bool)
	}

	if _, ok := graph.edges[n1][n2]; !ok {
		graph.edges[n1][n2] = true
	}

	if _, ok := graph.edges[n2]; !ok {
		graph.edges[n2] = make(map[Node]bool)
	}

	if _, ok := graph.edges[n2][n1]; !ok {
		graph.edges[n2][n1] = true
	}
}

func (graph *Graph) GetNodes() []Node {
	var nodes []Node

	for key := range graph.edges {
		nodes = append(nodes, key)
	}

	return nodes
}

func (graph *Graph) GetNeighbours(n Node) []Node {
	var nodes []Node

	for value := range graph.edges[n] {
		nodes = append(nodes, value)
	}

	return nodes
}

// -----------------------------------------------------------------------------

func (queue *Queue) Pop() Node{
	n := queue.list[0]
	queue.list = queue.list[1:]

	return n
}

func (queue *Queue) Push(n Node) {
	queue.list = append(queue.list, n)
}

// -----------------------------------------------------------------------------

func BFS(graph Graph, s Node) bool {
	// add node s to the queue
	q := NewQueue()
	q.Push(s)
	for len(q.List) > 0 {

	}
	// get one of the neighbours that isn't itself and then add that to the queue

	// repeat step 2 until you reach the desired node. and return true

}

// -----------------------------------------------------------------------------

func main() {
	g := NewGraph()

	s := NewNode("s")
	a := NewNode("a")
	b := NewNode("b")
	c := NewNode("c")
	d := NewNode("d")
	e := NewNode("e")

	g.AddEdge(s,a)
	g.AddEdge(s,b)
	g.AddEdge(a,c)
	g.AddEdge(b,c)
	g.AddEdge(b,d)
	g.AddEdge(c,d)
	g.AddEdge(c,e)
	g.AddEdge(d,e)


}

