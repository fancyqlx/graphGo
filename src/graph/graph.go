package graph

// Message is a triple for sending among vertices
type Message struct {
	hop, src, dist int
}

// Vertex is a struct for describing a vertex
// id is an integer for identifying a vertex
// neighbors is a slice storing ids of neighbors
// inMsg is a slice storing Messages from its neighbors
// outMsg is a map storing Messages to its neighbors, the false value means the message has been sent
// edges is a map storing values of edges between its neighbors and itself
type Vertex struct {
	id        int
	neighbors map[int]bool
	edges     map[int]int
	inMsg     []*Message
	outMsg    map[*Message]bool
}

// Graph is a struct for schedulering the algorithm,
// collecting necessory information and storing all
// the information of this graph
type Graph struct {
	Round            int
	Alpha, Beta, T   int
	Girth            int
	Number, Diameter int
	Vertices         map[int]*Vertex
}

// NewMessage is a factory method for Message
func NewMessage(hop, src, dist int) *Message {
	return &Message{hop, src, dist}
}

// Update is a method for updating data of a message
func (msg *Message) Update(h, s, d int) {
	msg.hop = h
	msg.src = s
	msg.dist = d
}

// NewVertex is a factory method for Vertex
func NewVertex(i int) *Vertex {
	return &Vertex{i, make(map[int]bool), make(map[int]int), make([]*Message, 0), make(map[*Message]bool)}
}

// AddNeighbor is a method for adding neighbor's id and weight value between them
func (v *Vertex) AddNeighbor(i int, w int) {
	if _, ok := v.neighbors[i]; ok {
		return
	}
	v.neighbors[i] = true
	v.edges[i] = w
}

// Clear is a method for clear inMsg and outMsg of a vertex
func (v *Vertex) Clear() {
	v.inMsg = nil
	v.outMsg = make(map[*Message]bool)
}

// Sendto is a method for sending messages to its neighbors
// This method will select i valid messages from outMsg, then
// send them to vertex's neighbors
func (v *Vertex) Sendto(i int) {

}

// NewGraph is a factory method
func NewGraph() *Graph {
	return &Graph{Vertices: make(map[int]*Vertex)}
}

// AddItem is used for adding a formated line from the graphdata
// Item consists of source id, destination id and the edge weight between them
func (g *Graph) AddItem(i, j, w int) {
	_, oku := g.Vertices[i]
	_, okv := g.Vertices[j]
	if oku && okv {

	} else if oku {
		g.Vertices[j] = NewVertex(j)
		g.Number++
	} else if okv {
		g.Vertices[i] = NewVertex(i)
		g.Number++
	} else {
		g.Vertices[i] = NewVertex(i)
		g.Vertices[j] = NewVertex(j)
		g.Number += 2
	}
	g.Vertices[i].AddNeighbor(j, w)
	g.Vertices[j].AddNeighbor(i, w)
}

// Clear is a method to clean all the cached data of vertices
func (g *Graph) Clear() {
	for _, v := range g.Vertices {
		v.Clear()
	}
}
