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
	neighbors []int
	edges     map[int]int
	inMsg     []*Message
	outMsg    map[Message]bool
}

// Graph is a struct for schedulering the algorithm,
// collecting necessory information and storing all
// the information of this graph
type Graph struct {
	Round            int
	Alpha, Beta, T   int
	Girth            int
	Number, Diameter int
	Vertices         []*Vertex
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
	return &Vertex{i, make([]int, 0), make(map[int]int), make([]*Message, 0), make(map[Message]bool)}
}

// Clean is a method for clear inMsg and outMsg of a vertex
func (v *Vertex) Clean() {
	v.inMsg = nil
	v.outMsg = make(map[Message]bool)
}

// Sendto is a method for sending messages to its neighbors
// This method will select i valid messages from outMsg, then
// send them to vertex's neighbors
func (v *Vertex) Sendto(i int) {

}

// NewGraph is a factory method
func NewGraph() *Graph {
	return &Graph{}
}

