package queue

type Queue[T any] interface {
	Enqueue(val T)
	Dequeue() T
}
