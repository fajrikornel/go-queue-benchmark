package queue

import (
	"fmt"
	"strings"
)

type Queue[T any] struct {
	queue    []T
	front    int
	back     int
	capacity int
}

func New[T any](capacity int) *Queue[T] {
	return &Queue[T]{
		queue:    make([]T, capacity),
		front:    -1,
		back:     -1,
		capacity: capacity,
	}
}

func (q *Queue[T]) String() string {
	values := make([]string, q.capacity)
	for i := 0; i < q.capacity; i++ {
		values[i] = fmt.Sprintf("%v", q.queue[i])
	}

	return strings.Join(values, ",")
}

func (q *Queue[T]) Enqueue(val T) {
	if q.front < 0 {
		q.front = q.getNextIndex(q.front)
		q.back = q.getNextIndex(q.back)
		q.queue[q.front] = val
		return
	}
}

func (q *Queue[T]) getNextIndex(i int) int {
	if i+1 > q.capacity-1 {
		return 0
	}

	return i + 1
}
