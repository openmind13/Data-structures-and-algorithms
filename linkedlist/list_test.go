package linkedlist_test

import (
	"testing"

	"github.com/openmind13/data-structures-and-algorithms/linkedlist"
	"github.com/stretchr/testify/assert"
)

func TestLinkedList(t *testing.T) {
	mylist := linkedlist.New()
	mylist.AddHead(10)
	mylist.AddHead(500)
	mylist.AddHead(300)
	assert.Equal(t, 3, mylist.Size())

	mylist.AddTail(200)
	assert.Equal(t, 4, mylist.Size())
}
