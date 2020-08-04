package main

import (
	"fmt"

	"github.com/openmind13/data-structures-and-algorithms/binarytree"
)

func main() {
	fmt.Printf("data-structures-and-algorithms\n\n")

	tree := binarytree.New()
	tree.Add(100)
	tree.Add(10)
	tree.Add(200)
	tree.Add(10)
	tree.Add(500)
	tree.Add(150)

	tree.Print()
}
