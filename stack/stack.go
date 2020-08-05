package stack

import "fmt"

// Stack ...
type Stack struct {
	buffer []int
}

// New ...
func New() Stack {
	return Stack{
		buffer: nil,
	}
}

// Push ...
func (stack *Stack) Push(value int) {
	stack.buffer = append(stack.buffer, value)
}

// Pop ...
func (stack *Stack) Pop() int {
	last := stack.buffer[len(stack.buffer)-1]
	stack.buffer = stack.buffer[:len(stack.buffer)-1]
	return last
}

// Print ...
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

// Size ...
func (stack *Stack) Size() int {
	return len(stack.buffer)
}

// Top ...
func (stack *Stack) Top() int {
	return stack.buffer[len(stack.buffer)-1]
}
