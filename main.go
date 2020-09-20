package main

import (
	"fmt"

	"github.com/openmind13/data-structures-and-algorithms/binarytree"
	"github.com/openmind13/data-structures-and-algorithms/linkedlist"
	"github.com/openmind13/data-structures-and-algorithms/queue"
	"github.com/openmind13/data-structures-and-algorithms/sorting"
	"github.com/openmind13/data-structures-and-algorithms/stack"
)

func main() {
	fmt.Printf("------------------------------\n")
	fmt.Printf("data-structures-and-algorithms\n\n")
	// fmt.Printf("------------------------------\n")

	//testBubbleSort()
	//testQuickSort()
	//testInsertionSort()
	//testSelectionSort()

	//testStack()
	//testQueue()

	testLinkedList()

	// testBinaryTree()

	fmt.Printf("------------------------------\n")
}

func testBubbleSort() {
	slice := []int{10, 50, 12, 7, 4, 9, 33, 37}
	fmt.Println(slice)
	sortSlice := sorting.BubbleSort(slice)
	fmt.Println(sortSlice)
}

func testInsertionSort() {
	slice := []int{3, 2, 5, 7, 1, 20, -4, 2, 9, 0}
	fmt.Println(slice)
	sortedSlice := sorting.InsertionSort(slice)
	fmt.Println(sortedSlice)
}

func testSelectionSort() {
	slice := []int{6, 7, 10, 15, 1, 4, -2, 5}
	fmt.Println(slice)
	sortedSlice := sorting.SelectionSort(slice)
	fmt.Println(sortedSlice)
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
	fmt.Println("testing int stack")
	intStack := stack.New()
	intStack.Push(10)
	intStack.Push(500)
	fmt.Println(intStack.Top())
	intStack.Print()
	fmt.Println(intStack.Size())
	intStack.Clear()
	intStack.Print()

	fmt.Println("testing string stack")
	stringStack := stack.New()
	stringStack.Push("hello")
	stringStack.Push("world")
	fmt.Println(stringStack.Top())
	stringStack.Print()
	fmt.Println(stringStack.Size())
	stringStack.Clear()
	stringStack.Print()
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
	fmt.Println("testing int queue")
	intQueue := queue.New()
	// fmt.Println(queue.IsEmpty())
	intQueue.Add(10)
	intQueue.Add(20)
	intQueue.Add(30)
	intQueue.Print()
	intQueue.Add(50)
	intQueue.Print()
	fmt.Println(intQueue.Get())
	intQueue.Print()
	fmt.Println(intQueue.Size())
	intSlice := intQueue.Dequeue()
	fmt.Println(intSlice)

	fmt.Println("testing string queue")
	stringQueue := queue.New()
	stringQueue.Add("the")
	stringQueue.Add("go")
	stringQueue.Add("programming")
	stringQueue.Add("language")
	stringQueue.Print()
	fmt.Println(stringQueue.Get())
	stringQueue.Print()
	fmt.Println(stringQueue.Size())
	stringSlice := stringQueue.Dequeue()
	fmt.Println(stringSlice)
}
