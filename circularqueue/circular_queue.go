package queue

import "github.com/golang-ds/linkedlist/circularly"

type CircularQueue[T any] struct {
	data circularly.LinkedList[T]
}

// New constructs and returns an empty circular-queue.
// time-complexity: O(1)
func New[T any]() CircularQueue[T] {
	return CircularQueue[T]{}
}

// Enqueue adds an element to the end of the queue.
// time-complexity: O(1)
func (q *CircularQueue[T]) Enqueue(data T) {
	q.data.AddLast(data)
}

// Dequeue removes and returns the front element of the queue. It returns false if the queue was empty.
// time-complexity: O(1)
func (q *CircularQueue[T]) Dequeue() (val T, ok bool) {
	return q.data.RemoveFirst()
}

// First returns the front element of the queue. It returns false if the queue was empty.
// time-complexity: O(1)
func (q *CircularQueue[T]) First() (val T, ok bool) {
	return q.data.First()
}

// Size returns the number of the elements in the queue.
// time-complexity: O(1)
func (q *CircularQueue[T]) Size() int {
	return q.data.Size
}

// IsEmpty returns true if the queue is empty.
// time-complexity: O(1)
func (q *CircularQueue[T]) IsEmpty() bool {
	return q.data.IsEmpty()
}

// Rotate moves the first element to the end of the queue.
// time-complexity: O(1)
func (q *CircularQueue[T]) Rotate() {
	q.data.Rotate()
}

// String returns the string representation of the queue.
// time-complexity: O(1)
func (q *CircularQueue[T]) String() string {
	return q.data.String()
}
