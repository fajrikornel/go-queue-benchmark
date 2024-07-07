package queue

import (
	"fmt"
	"strings"
)

type ListQueue[T any] struct {
	front *ListNode[T]
	back  *ListNode[T]
}

type ListNode[T any] struct {
	val  T
	next *ListNode[T]
}

func NewListQueue[T any]() *ListQueue[T] {
	return &ListQueue[T]{
		front: nil,
		back:  nil,
	}
}

func (q *ListQueue[T]) String() string {
	values := []string{}
	cur := q.front
	for cur != nil {
		values = append(values, fmt.Sprintf("%v", cur.val))
	}

	return strings.Join(values, ",")
}

func (q *ListQueue[T]) Enqueue(val T) {

}

func (q *ListQueue[T]) Dequeue() T {
	return q.front.val
}
