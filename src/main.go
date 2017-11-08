package main

import (
	"graphGo/src/graph"
)

func main() {
	g := graph.ReadGraph("../graphdata/data.in")
	graph.PrintGraph(g)
}
