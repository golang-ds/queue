package queue

import "fmt"

type SliceQueue[T any] struct {
	data []T
}

// New constructs and returns an empty slice-queue.
// time-complexity: O(1)
func New[T any]() SliceQueue[T] {
	return SliceQueue[T]{data: []T{}}
}

// Enqueue adds an element to the end of the queue.
// time-complexity: amortized O(1)
func (q *SliceQueue[T]) Enqueue(data T) {
	q.data = append(q.data, data)
}

// Dequeue removes and returns the front element of the queue. It returns false if the queue was empty.
// time-complexity: O(1)
func (q *SliceQueue[T]) Dequeue() (val T, ok bool) {
	if q.IsEmpty() {
		return
	}

	val = q.data[0]

	q.data = q.data[1:]

	return val, true
}

// First returns the front element of the queue. It returns false if the queue was empty.
// time-complexity: O(1)
func (q *SliceQueue[T]) First() (val T, ok bool) {
	if q.IsEmpty() {
		return
	}

	return q.data[0], true
}

// Size returns the number of the elements in the queue.
// time-complexity: O(1)
func (q *SliceQueue[T]) Size() int {
	return len(q.data)
}

// IsEmpty returns true if the queue is empty.
// time-complexity: O(1)
func (q *SliceQueue[T]) IsEmpty() bool {
	return q.Size() == 0
}

// String returns the string representation of the queue.
// time-complexity: O(n)
func (q *SliceQueue[T]) String() string {
	return fmt.Sprint(q.data)
}
