package linkedlist

import (
	"errors"
	"fmt"
)

// ItemType - type of element
type ItemType interface{}

// List struct contain information about list
type List struct {
	length int
	head   *nodeType
	tail   *nodeType
}

type nodeType struct {
	value ItemType
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
func (list *List) AddItem(index int, value ItemType) error {
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
func (list *List) AddHead(value ItemType) {
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
func (list *List) AddTail(value ItemType) {
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

// RemoveItem - find first value and remove it from list
func (list *List) RemoveItem(item ItemType) error {
	node := list.head

	for node != nil {
		if node.value.(int) == item.(int) {
			// if only one element in list
			if node.next == nil && node.prev == nil {
				list.head = nil
				list.tail = nil
			}

			// in the middle
			if node.prev != nil && node.next != nil {
				next := node.next
				prev := node.prev

				next.prev = prev
				prev.next = next
			}

			// if element is firts
			if node.prev == nil && node.next != nil {
				next := node.next
				next.prev = nil

				list.head = next
			}

			// if element if last
			if node.next == nil && node.prev != nil {
				prev := node.prev
				prev.next = nil

				list.tail = prev
			}

			node = nil
			list.length--
			return nil
		}

		node = node.next
	}

	return errors.New("Item not found")
}

// Clear all list
func (list *List) Clear() {
	// GC will delete all unattainable nodes
	list.head = nil
	list.tail = nil

	list.length = 0
}
