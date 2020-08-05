package main

import (
	"fmt"

	"github.com/openmind13/data-structures-and-algorithms/binarytree"
	"github.com/openmind13/data-structures-and-algorithms/stack"
)

func main() {
	fmt.Printf("------------------------------\n")
	fmt.Printf("data-structures-and-algorithms\n")
	fmt.Printf("------------------------------\n")

	// testBinaryTree()
	testLinkedList()
}

func testBinaryTree() {
	fmt.Printf("testing binary tree\n")
	fmt.Printf("-------------------\n")

	tree := binarytree.New()
	tree.Add(100)
	tree.Add(10)
	tree.Add(200)
	tree.Add(10)
	tree.Add(500)
	tree.Add(150)

	tree.Print()
	tree.Clear()
	fmt.Printf("after clear\n")

	tree.Print()

	tree.Add(50)
	tree.Add(80)

	tree.Print()
}

func testLinkedList() {
	fmt.Printf("testing stack\n")
	fmt.Printf("-------------\n")

	stack := stack.New()
	stack.Push(10)
	stack.Push(500)

	fmt.Println(stack.Top())

	stack.Print()

	fmt.Println(stack.Size())
}
