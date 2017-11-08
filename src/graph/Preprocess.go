package graph

import (
	"fmt"
	"os"
)

// ReadGraph is used for reading graph data from files
// Then construct Graph struct and return a pointer to it
func ReadGraph(path string) *Graph {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	g := NewGraph()
	var i, j, w int
	for {
		_, err := fmt.Fscanln(file, &i, &j, &w)
		if err != nil {
			break
		}
		g.AddItem(i, j, w)
	}
	return g
}

// PrintGraph is used to print the information of struct Graph
func PrintGraph(g *Graph) {
	fmt.Printf("Rounds: %d\n", g.Round)
	fmt.Printf("Girth: %d\n", g.Girth)
	fmt.Printf("The number of the graph is %d\n", g.Number)
	fmt.Printf("Diameter: %d\n", g.Diameter)
}
