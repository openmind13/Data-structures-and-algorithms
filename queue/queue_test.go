package queue_test

import (
	"testing"

	"github.com/openmind13/data-structures-and-algorithms/queue"
	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	intQueue := queue.New()
	intQueue.Add(10)
	intQueue.Add(20)
	intQueue.Add(30)

	item, err := intQueue.Get()
	assert.NoError(t, err)
	assert.Equal(t, 10, item)
}
