package stack_test

import (
	"testing"

	"github.com/openmind13/data-structures-and-algorithms/stack"
	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	intStack := stack.New()
	intStack.Push(10)
	intStack.Push(500)
	intStack.Push(200)

	assert.Equal(t, 3, intStack.Size())

	item, err := intStack.Pop()
	assert.NoError(t, err)
	assert.Equal(t, 200, item)

	assert.Equal(t, 500, intStack.Top())

	assert.Equal(t, 2, intStack.Size())

	intStack.Clear()
	assert.Equal(t, 0, intStack.Size())
}
