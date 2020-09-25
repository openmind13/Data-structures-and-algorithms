package graph_test

import (
	"testing"

	"github.com/openmind13/data-structures-and-algorithms/graph"
	"github.com/stretchr/testify/assert"
)

func TestGraph(t *testing.T) {
	g1, err := graph.NewGraphFromTGFFile("./test_graphs/g1.tgf")
	assert.NoError(t, err)
	assert.Equal(t, 7, g1.CountVertexes())
}
