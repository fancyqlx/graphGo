package main

import (
	"graphGo/src/graph"
)

func main() {
	g := graph.ReadGraph("../graphdata/data.in")

	g.Diameter = graph.Diameter(g)
	g.Init(g.Number)
	g.BoundedBFS()
	/*for _, v := range g.Vertices {
		graph.PrintVertex(v)
	}*/
	graph.PrintGraph(g)
}
