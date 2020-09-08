package main

import (
	"fmt"

	"github.com/openmind13/data-structures-and-algorithms/binarytree"
	"github.com/openmind13/data-structures-and-algorithms/linkedlist"
	"github.com/openmind13/data-structures-and-algorithms/queue"
	"github.com/openmind13/data-structures-and-algorithms/stack"
)

func main() {
	fmt.Printf("------------------------------\n")
	fmt.Printf("data-structures-and-algorithms\n")
	// fmt.Printf("------------------------------\n")

	testBinaryTree()
	// testStack()
	// testLinkedList()
	// testQueue()
}

func testBinaryTree() {
	fmt.Printf("testing binary tree\n")
	fmt.Printf("-------------------\n")

	tree := binarytree.New()
	fmt.Println(tree.IsEmpty())
	tree.Add(100)
	tree.Add(10)
	tree.Add(200)
	tree.Add(10)
	tree.Add(500)
	tree.Add(150)
	fmt.Println(tree.Depth())

	tree.Print()
	tree.Clear()
	fmt.Printf("after clear\n")

	tree.Print()

	tree.Add(50)
	tree.Add(80)

	tree.Print()
}

func testStack() {
	fmt.Printf("-------------\n")
	fmt.Printf("testing stack\n")
	fmt.Printf("-------------\n")

	stack := stack.New()

	stack.Push(10)
	stack.Push(500)

	fmt.Println(stack.Top())

	stack.Print()

	fmt.Println(stack.Size())

	stack.Clear()
	stack.Print()
}

func testLinkedList() {
	fmt.Printf("-------------------\n")
	fmt.Printf("testing linked list\n")
	fmt.Printf("-------------------\n")

	list := linkedlist.New()
	//fmt.Println(list.IsEmpty())

	list.AddHead(10)
	list.AddHead(500)
	list.AddTail(300)
	for i := 0; i < 10; i++ {
		list.AddHead(i)
	}
	fmt.Println(list.Size())
	list.Print()

	list.Clear()
	// fmt.Println("after clear")

	// fmt.Println(list.IsEmpty())
	// list.Print()
}

func testQueue() {
	fmt.Printf("-------------\n")
	fmt.Printf("testing queue\n")
	fmt.Printf("-------------\n")

	queue := queue.New()
	// fmt.Println(queue.IsEmpty())
	queue.Add(10)
	queue.Add(20)
	queue.Add(30)
	queue.Print()
	queue.Add(50)
	queue.Print()
	fmt.Println(queue.Get())
	queue.Print()
	fmt.Println(queue.Size())

	slice := queue.Dequeue()
	fmt.Println(slice)
}
