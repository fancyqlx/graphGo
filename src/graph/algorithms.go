package graph

import "container/heap"

// BellmanFord is the method for simulating Bellman-Ford algorithm
func BellmanFord(g *Graph) {
	// Initailize msgList from outMsg
	for _, v := range g.Vertices {
		item := heap.Pop(&v.outMsg).(*Item)
		v.inMsg = append(v.inMsg, item.value)
	}
	// Bellman-Ford loop
	for i := 0; i < g.Number; i++ {
		for _, v := range g.Vertices {
			send(g, v)
		}
	}
}

func send(g *Graph, v *Vertex) {
	for _, msg := range v.inMsg {
		if value, ok := v.msgList[msg.src]; ok {
			if msg.hop >= value.hop {
				continue
			}
		}
		v.msgList[msg.src] = msg
		msgforsend := NewMessage(msg.hop+1, msg.src, msg.dist, msg.pre)
		for k := range v.neighbors {
			g.Vertices[k].inMsg = append(g.Vertices[k].inMsg, msgforsend)
		}
	}
	v.inMsg = nil
}

// Diameter is the method for computing diameter of g
func Diameter(g *Graph) int {
	BellmanFord(g)
	var D int
	for _, v := range g.Vertices {
		for _, msg := range v.msgList {
			if msg.hop > D {
				D = msg.hop
			}
		}
	}
	return D
}
