package main

import (
	"graphGo/src/graph"
)

func main() {
	g := graph.Preprocess("../graphdata/data.in")
	graph.PrintGraph(g)
}
