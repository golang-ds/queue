package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnqueue(t *testing.T) {
	queue := New[int]()

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	str := queue.String()
	assert.Equal(t, "[ 1 2 3 ]", str)
}

func TestDequeue(t *testing.T) {
	queue := New[int]()

	f, ok := queue.Dequeue()
	assert.False(t, ok)
	assert.Equal(t, 0, f)

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	f, ok = queue.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, 1, f)

	str := queue.String()
	assert.Equal(t, "[ 2 3 ]", str)
}

func TestFirst(t *testing.T) {
	queue := New[int]()

	f, ok := queue.First()
	assert.False(t, ok)
	assert.Equal(t, 0, f)

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	f, ok = queue.First()
	assert.True(t, ok)
	assert.Equal(t, 1, f)

	str := queue.String()
	assert.Equal(t, "[ 1 2 3 ]", str)
}

func TestSize(t *testing.T) {
	queue := New[int]()

	size := queue.Size()
	assert.Equal(t, 0, size)

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	size = queue.Size()
	assert.Equal(t, 3, size)
}

func TestIsEmpty(t *testing.T) {
	queue := New[int]()

	assert.True(t, queue.IsEmpty())

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	assert.False(t, queue.IsEmpty())
}

func TestRotate(t *testing.T) {
	queue := New[int]()

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	queue.Rotate()

	f, ok := queue.First()
	assert.True(t, ok)
	assert.Equal(t, 2, f)

}

func TestString(t *testing.T) {
	queue := New[int]()

	str := queue.String()
	assert.Equal(t, "[ ]", str)

	queue.Enqueue(1)

	str = queue.String()
	assert.Equal(t, "[ 1 ]", str)

	queue.Enqueue(2)
	queue.Enqueue(3)

	str = queue.String()
	assert.Equal(t, "[ 1 2 3 ]", str)
}
