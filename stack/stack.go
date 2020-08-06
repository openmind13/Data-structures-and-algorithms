package stack

import (
	"errors"
	"fmt"
)

// ItemType - type of items
type ItemType int

// Stack struct
type Stack struct {
	buffer []ItemType
}

// New - create new stack
func New() Stack {
	return Stack{
		buffer: nil,
	}
}

// Push item on top of stack
func (stack *Stack) Push(value ItemType) {
	stack.buffer = append(stack.buffer, value)
}

// Pop item from top
func (stack *Stack) Pop() (ItemType, error) {
	if stack.IsEmpty() {
		return 0, errors.New("Stack is empty")
	}

	last := stack.buffer[len(stack.buffer)-1]
	stack.buffer = stack.buffer[:len(stack.buffer)-1]
	return last, nil
}

// Print stack
func (stack *Stack) Print() {
	fmt.Printf("_______________\n")

	for i := len(stack.buffer) - 1; i >= 0; i-- {
		if i == len(stack.buffer)-1 {
			fmt.Printf("%v  <-- top\n", stack.buffer[i])
		} else if i == 0 {
			fmt.Printf("%v  <-- bottom\n", stack.buffer[i])
		} else {
			fmt.Printf("%v\n", stack.buffer[i])
		}
	}

	fmt.Printf("^^^^^^^^^^^^^^^\n")
}

// Clear - delete all items from stack
func (stack *Stack) Clear() {
	stack.buffer = nil
}

// Size - return count of items
func (stack *Stack) Size() int {
	return len(stack.buffer)
}

// Top - return top value
func (stack *Stack) Top() ItemType {
	return stack.buffer[len(stack.buffer)-1]
}

// IsEmpty return bool value
func (stack *Stack) IsEmpty() bool {
	return stack.buffer == nil
}
