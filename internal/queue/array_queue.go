package queue

import (
	"fmt"
	"strings"
)

type ArrayQueue[T any] struct {
	queue    []T
	front    int
	back     int
	capacity int
}

func NewArrayQueue[T any](capacity int) *ArrayQueue[T] {
	return &ArrayQueue[T]{
		queue:    make([]T, capacity),
		front:    -1,
		back:     -1,
		capacity: capacity,
	}
}

func (q *ArrayQueue[T]) String() string {
	values := make([]string, q.capacity)
	for i := 0; i < q.capacity; i++ {
		values[i] = fmt.Sprintf("%v", q.queue[i])
	}

	return strings.Join(values, ",")
}

func (q *ArrayQueue[T]) Enqueue(val T) {
	if q.front < 0 {
		q.front = q.getNextIndex(q.front)
		q.back = q.getNextIndex(q.back)
		q.queue[q.back] = val
		return
	}

	next := q.getNextIndex(q.back)
	if next == q.front {
		panic("Queue is full!")
	}

	q.back = next
	q.queue[next] = val
}

func (q *ArrayQueue[T]) Dequeue() T {
	if q.front < 0 {
		panic("Queue is empty!")
	}

	val := q.queue[q.front]
	if q.front == q.back {
		q.back = -1
		q.front = -1
	} else {
		q.front = q.getNextIndex(q.front)
	}

	return val
}

func (q *ArrayQueue[T]) getNextIndex(i int) int {
	if i+1 > q.capacity-1 {
		return 0
	}

	return i + 1
}
