package dsa

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
