package main

import (
	"testing"

	q "example.com/go-dsa/internal/queue"
)

func BenchmarkArrayQueue_Enqueue(b *testing.B) {
	queue := q.NewArrayQueue[int](1_000_000)

	b.ResetTimer()
	for i := 0; i < 1_000_000; i++ {
		queue.Enqueue(123)
	}
}

func BenchmarkListQueue_Enqueue(b *testing.B) {
	queue := q.NewListQueue[int]()

	b.ResetTimer()
	for i := 0; i < 1_000_000; i++ {
		queue.Enqueue(123)
	}
}

func BenchmarkArrayQueue_Dequeue(b *testing.B) {
	queue := q.NewArrayQueue[int](1_000_000)
	for i := 0; i < 1_000_000; i++ {
		queue.Enqueue(123)
	}

	b.ResetTimer()
	for i := 0; i < 1_000_000; i++ {
		queue.Dequeue()
	}
}

func BenchmarkListQueue_Dequeue(b *testing.B) {
	queue := q.NewListQueue[int]()
	for i := 0; i < 1_000_000; i++ {
		queue.Enqueue(123)
	}

	b.ResetTimer()
	for i := 0; i < 1_000_000; i++ {
		queue.Dequeue()
	}
}
