package linkedlist

import (
	"errors"
	"fmt"
)

// List struct contain information about list
type List struct {
	length int
	head   *nodeType
	tail   *nodeType
}

type nodeType struct {
	value int
	next  *nodeType
	prev  *nodeType
}

// New create new list
func New() List {
	return List{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

// Size return list size
func (list *List) Size() int {
	return list.length
}

// IsEmpty return empty list or not
func (list *List) IsEmpty() bool {
	return list.head == nil && list.tail == nil && list.length == 0
}

// Print all list
func (list *List) Print() {
	printedItems := 0
	node := list.tail

	for node != nil {
		printedItems++
		if printedItems < list.length {
			fmt.Printf("%v <-> ", node.value)
		} else {
			fmt.Printf("%v\n", node.value)
		}
		node = node.next
	}
}

// AddItem add item on index
func (list *List) AddItem(index, value int) error {
	if index >= list.length {
		fmt.Printf("Get out of range. Length of list = %d. Valid index = [0 - %d]\n",
			list.length, list.length-1)
		return errors.New("Get out of range")
	}

	if list.head == nil && list.tail == nil && list.length == 0 {
		fmt.Printf("List is empty\n")
		return errors.New("List is empty")
	}

	list.length++
	return nil
}

// AddHead add item in back of list
func (list *List) AddHead(value int) {
	if list.head != nil {
		list.head.next = &nodeType{
			value: value,
			next:  nil,
			prev:  list.head,
		}
		list.head = list.head.next
	} else {
		// if first item in list
		list.head = &nodeType{
			value: value,
			next:  nil,
			prev:  nil,
		}
		list.tail = list.head
	}

	list.length++
}

// AddTail add item in front of list
func (list *List) AddTail(value int) {
	if list.tail != nil {
		list.tail.prev = &nodeType{
			value: value,
			next:  list.tail,
			prev:  nil,
		}
		list.tail = list.tail.prev
	} else {
		// if first item in list
		list.tail = &nodeType{
			value: value,
			next:  nil,
			prev:  nil,
		}
		list.head = list.tail
	}

	list.length++
}

// Clear all list
func (list *List) Clear() {
	// GC will delete all unattainable nodes
	list.head = nil
	list.tail = nil

	list.length = 0
}
