package binarytree_test

import (
	"testing"

	"github.com/openmind13/data-structures-and-algorithms/binarytree"
	"github.com/stretchr/testify/assert"
)

func TestBinaryTree(t *testing.T) {
	tree := binarytree.New()
	assert.Equal(t, true, tree.IsEmpty())
	tree.Add(100)
	tree.Add(10)
	tree.Add(200)
	tree.Add(500)
	tree.Add(150)
	assert.Equal(t, 3, tree.Depth())

	// tree.RemoveItem(500)
	// tree.RemoveItem(150)
	// assert.Equal(t, 2, tree.Depth())

	tree.Clear()
	assert.Equal(t, 0, tree.Depth())

}
