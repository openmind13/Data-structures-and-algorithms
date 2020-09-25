package sorting_test

import (
	"testing"

	"github.com/openmind13/data-structures-and-algorithms/sorting"
	"github.com/stretchr/testify/assert"
)

func TestSortings(t *testing.T) {
	slice := []int{9, 17, 4, 5, 1, 29, 11, -2}
	expectedSlice := []int{-2, 1, 4, 5, 9, 11, 17, 29}

	assert.Equal(t, expectedSlice, sorting.BubbleSort(slice))
	assert.Equal(t, expectedSlice, sorting.InsertionSort(slice))
	assert.Equal(t, expectedSlice, sorting.SelectionSort(slice))
}
