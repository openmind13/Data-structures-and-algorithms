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

	tgfFileLines []string
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

// PrintIncidenceMatrix ...
func (graph *Graph) PrintIncidenceMatrix() {
	for i := 0; i < graph.countVertexes; i++ {
		for j := 0; j < graph.countArcs; j++ {
			fmt.Printf("%v\t", graph.incidenceMatrix[i][j])
		}
		fmt.Printf("\n\n\n")
	}
}

// PrintTGFFile ...
func (graph *Graph) PrintTGFFile() {
	for _, str := range graph.tgfFileLines {
		fmt.Printf("%s", str)
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

		graph.tgfFileLines = append(graph.tgfFileLines, str)

		if strings.Contains(str, "#") {
			break
		}

		vertexes++

		// parse str
		// fmt.Printf("%s", str)
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

		graph.tgfFileLines = append(graph.tgfFileLines, str)

		arcs++

		// parse str
		//splited := strings.Split(str, " ")
		// fmt.Printf("%s", str)
	}

	graph.countVertexes = vertexes
	graph.countArcs = arcs / 2

	// init adjacency matrix
	graph.adjacencyMatrix = make([][]int, graph.countVertexes)
	for i := range graph.adjacencyMatrix {
		graph.adjacencyMatrix[i] = make([]int, graph.countVertexes)
	}

	// init incidence matrix
	graph.incidenceMatrix = make([][]int, graph.countVertexes)
	for i := range graph.incidenceMatrix {
		graph.incidenceMatrix[i] = make([]int, graph.countArcs)
	}

	return graph, nil
}
