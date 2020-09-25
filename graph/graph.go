package graph

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Graph struct
type Graph struct {
	countVertexes int
	countArcs     int

	incidenceMatrix [][]int
	adjacencyMatrix [][]int
}

// CountVertexes return vertexes count of current graph
func (graph *Graph) CountVertexes() int {
	return graph.countVertexes
}

// CountArcs return arcs count of current graph
func (graph *Graph) CountArcs() int {
	return graph.countArcs
}

// PrintAdjacencyMatrix ...
func (graph *Graph) PrintAdjacencyMatrix() {
	for i := 0; i < graph.countVertexes; i++ {
		for j := 0; j < graph.countVertexes; j++ {
			fmt.Printf("%v\t", graph.adjacencyMatrix[i][j])
		}
		fmt.Printf("\n\n\n")
	}
}

// NewGraphFromTGFFile ...
func NewGraphFromTGFFile(filename string) (*Graph, error) {
	graph := &Graph{}

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	reader := bufio.NewReader(file)

	var (
		vertexes = 0
		arcs     = 0
	)

	// scan vertexes
	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			return nil, err
		}

		if strings.Contains(str, "#") {
			break
		}

		vertexes++
		fmt.Printf("%s", str)
	}

	// init adjacency matrix
	graph.adjacencyMatrix = make([][]int, vertexes)
	for i := range graph.adjacencyMatrix {
		graph.adjacencyMatrix[i] = make([]int, vertexes)
	}

	// scan arcs
	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				break
			}

			return nil, err
		}

		arcs++
		//splited := strings.Split(str, " ")
		fmt.Printf("%s", str)
	}

	graph.countVertexes = vertexes
	graph.countArcs = arcs / 2

	return graph, nil
}
