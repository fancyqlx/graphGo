package graph

import "container/heap"
import "math"

// Message is a triple for sending among vertices
type Message struct {
	hop, src, dist int
	pre            int
}

// Vertex is a struct for describing a vertex
// id is an integer for identifying a vertex
// neighbors is a slice storing ids of neighbors
// inMsg is a slice storing Messages from its neighbors
// outMsg is a map storing Messages to its neighbors, the false value means the message has been sent
// edges is a map storing values of edges between its neighbors and itself
// Pre is a map for storing predecessors
type Vertex struct {
	id        int
	neighbors map[int]bool
	edges     map[int]int
	inMsg     []*Message
	msgList   map[int]*Message
	outMsg    PriorityQueue
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
	B                int
	W                int
	Next             bool
}

// NewMessage is a factory method for Message
func NewMessage(hop, src, dist, pre int) *Message {
	return &Message{hop, src, dist, pre}
}

// Update is a method for updating data of a message
func (msg *Message) Update(h, s, d, p int) {
	msg.hop = h
	msg.src = s
	msg.dist = d
	msg.pre = p
}

// NewVertex is a factory method for Vertex
func NewVertex(i int) *Vertex {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	return &Vertex{i, make(map[int]bool), make(map[int]int), make([]*Message, 0), make(map[int]*Message), pq}
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
	v.msgList = make(map[int]*Message)
	v.outMsg = nil
}

// Sendto is a method for sending messages to its neighbors
// This method will select i valid messages from outMsg, then
// send them to vertex's neighbors
func (v *Vertex) Sendto(g *Graph, i int) {
	for ; i > -1 && len(v.outMsg) > 0; i-- {
		item := heap.Pop(&v.outMsg).(*Item)
		msg := item.value
		for k := range v.neighbors {
			if msg.pre != k && msg.dist+v.edges[k] <= g.T {
				msgforsend := NewMessage(msg.hop, msg.src, msg.dist, v.id)
				g.Vertices[k].inMsg = append(g.Vertices[k].inMsg, msgforsend)
			}
		}
		if _, ok := v.msgList[msg.src]; ok {
			g.Next = true
		} else {
			v.msgList[msg.src] = msg
		}
	}
}

// ProcessMsg is to process messages it received
func (v *Vertex) ProcessMsg() {
	for _, msg := range v.inMsg {
		msg.hop++
		msg.dist += v.edges[msg.pre]
		heap.Push(&v.outMsg, &Item{value: msg})
	}
	v.inMsg = nil
}

// NewGraph is a factory method
func NewGraph() *Graph {
	return &Graph{Vertices: make(map[int]*Vertex)}
}

// Init is a method used in the beginning of the algorithm
// It initializes parameters of the graph
func (g *Graph) Init(b int) {
	g.Alpha = 3
	g.Beta = g.Number * g.W
	g.T = 1
	g.B = b
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
	if w > g.W {
		g.W = w
	}
}

// Clear is a method to clean all the cached data of vertices
func (g *Graph) Clear() {
	for _, v := range g.Vertices {
		v.Clear()
	}
	g.Next = false
}

// ComputeG is to compute g_v
func (g *Graph) ComputeG() int {
	gv := math.MaxInt32
	for _, v := range g.Vertices {
		for k := range v.neighbors {
			for s := range v.msgList {
				if msg, ok := g.Vertices[k].msgList[s]; ok && v.id != msg.pre && k != v.msgList[s].pre {
					value := v.msgList[s].dist + msg.dist + v.edges[k]
					if value < gv {
						gv = value
					}
				}
			}
		}
	}
	return gv
}

// Superstep is one round of excuting the graph algorithm
func (g *Graph) Superstep() {
	num := 0
	condi1 := false
	for _, v := range g.Vertices {
		if len(v.outMsg) == 0 {
			num++
			continue
		}
		v.Sendto(g, g.B)
		if g.Next {
			condi1 = true
			break
		}
	}
	if num == g.Number {
		g.Next = true
	}
	for _, v := range g.Vertices {
		v.ProcessMsg()
	}
	if g.Next {
		gv := g.ComputeG()
		if condi1 {
			g.Beta = min(2*g.T, gv)
		} else {
			if gv > 2*g.T {
				g.Alpha = 2 * g.T
			} else {
				g.Beta = min(g.Beta, gv)
			}
		}
	}
}

// BoundedBFS is the main algorithm for computing girth
func (g *Graph) BoundedBFS() {
	for g.Beta-g.Alpha > 2 {
		for _, v := range g.Vertices {
			msg := NewMessage(0, v.id, 0, -1)
			heap.Push(&v.outMsg, &Item{value: msg})
		}

		for !g.Next {
			g.Superstep()
			g.Round++
		}
		if g.T == (g.Alpha+g.Beta)/4 {
			g.T++
		} else {
			g.T = (g.Alpha + g.Beta) / 4
		}
		g.Clear()
		println(g.T, g.Alpha, g.Beta)
	}
	g.Girth = g.Beta
}
