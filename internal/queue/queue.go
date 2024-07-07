package queue

import (
	"fmt"
	"strings"
)

type Queue[T any] interface {
	Enqueue(val T)
	Dequeue() T
}

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
		q.queue[q.front] = val
		return
	}

	next := q.getNextIndex(q.front)
	if next == q.back {
		panic("Queue is full!")
	}

	q.front = next
	q.queue[next] = val
}

func (q *ArrayQueue[T]) Dequeue() T {
	if q.front < 0 {
		panic("Queue is empty!")
	}

	val := q.queue[q.back]
	if q.back == q.front {
		q.front = -1
		q.back = -1
	} else {
		q.back = q.getNextIndex(q.back)
	}

	return val
}

func (q *ArrayQueue[T]) getNextIndex(i int) int {
	if i+1 > q.capacity-1 {
		return 0
	}

	return i + 1
}
