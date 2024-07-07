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
		cur = cur.next
	}

	return strings.Join(values, ",")
}

func (q *ListQueue[T]) Enqueue(val T) {
	newNode := &ListNode[T]{
		val:  val,
		next: nil,
	}

	if q.back == nil && q.front == nil {
		q.front = newNode
		q.back = newNode
		return
	}

	q.back.next = newNode
	q.back = newNode
}

func (q *ListQueue[T]) Dequeue() T {
	if q.front == nil {
		panic("Queue is empty!")
	}

	val := q.front.val
	q.front = q.front.next
	return val
}
