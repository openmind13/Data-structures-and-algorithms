package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/openmind13/data-structures-and-algorithms/sorting"
)

func main() {
	fmt.Printf("data-structures-and-algorithms\n")
}

func quickSort() {
	slice := generateSlice(10)
	fmt.Println(slice)
	sortedSlice := sorting.QuickSort(slice)
	fmt.Println(sortedSlice)
}

func generateSlice(size int) []int {
	slice := make([]int, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

func testLinkedList() {

}
