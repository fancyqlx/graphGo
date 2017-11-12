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
	fmt.Printf("Number %d\n", g.Number)
	fmt.Printf("Diameter: %d\n", g.Diameter)
}

// PrintVertex is used for printing the information of a vertex
func PrintVertex(v *Vertex) {
	fmt.Printf("id: %d\n", v.id)
	fmt.Printf("Neighbors: ")
	for k := range v.neighbors {
		fmt.Printf("%d ", k)
	}
	fmt.Println()
	fmt.Printf("inMsg: ")
	for _, msg := range v.inMsg {
		fmt.Printf("%v ", msg)
	}
	fmt.Println()
	fmt.Printf("msgList: ")
	for k, v := range v.msgList {
		fmt.Printf("%d, %v ", k, v)
	}
	fmt.Println()
}
