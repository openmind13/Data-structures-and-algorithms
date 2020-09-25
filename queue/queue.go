package queue

import (
	"errors"
	"fmt"
)

// ItemType - type of items
type ItemType interface{}

// Queue struct
type Queue struct {
	buffer []ItemType
}

// New - create new queue
func New() *Queue {
	return &Queue{
		buffer: nil,
	}
}

// Size return queue size
func (queue *Queue) Size() int {
	return len(queue.buffer)
}

// IsEmpty return bool value
func (queue *Queue) IsEmpty() bool {
	return queue.buffer == nil && len(queue.buffer) == 0
}

// Print all queue
func (queue *Queue) Print() {
	for _, v := range queue.buffer {
		fmt.Printf("<- %v ", v)
	}
	fmt.Printf("\n")
}

// Add ...
func (queue *Queue) Add(value ItemType) {
	queue.buffer = append(queue.buffer, value)
}

// Get ...
func (queue *Queue) Get() (ItemType, error) {
	if queue.IsEmpty() {
		return 0, errors.New("Queue is empty")
	}

	value := queue.buffer[0]
	queue.buffer = queue.buffer[1:]
	return value, nil
}

// Dequeue - return slice of elements
func (queue *Queue) Dequeue() []ItemType {
	return queue.buffer
}
