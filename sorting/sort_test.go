package sorting_test

import (
	"testing"

	"github.com/openmind13/data-structures-and-algorithms/sorting"
	"github.com/stretchr/testify/assert"
)

func TestSortings(t *testing.T) {
	slice := []int{9, 17, 4, 5, 1, 29, 11, -2}
	sortedSlice := []int{-2, 1, 4, 5, 9, 11, 17, 29}

	assert.Equal(t, sortedSlice, sorting.BubbleSort(slice))

	slice = []int{9, 17, 4, 5, 1, 29, 11, -2}
	assert.Equal(t, sortedSlice, sorting.InsertionSort(slice))

	slice = []int{9, 17, 4, 5, 1, 29, 11, -2}
	assert.Equal(t, sortedSlice, sorting.SelectionSort(slice))

	slice = []int{9, 17, 4, 5, 1, 29, 11, -2}
	assert.Equal(t, sortedSlice, sorting.QuickSort(slice))

	slice = []int{9, 17, 4, 5, 1, 29, 11, -2}
	assert.Equal(t, []int{-2, 11, 29, 1, 5, 4, 17, 9}, sorting.ReverseArray(slice))
}
