package queue

import (
	"log"
	"testing"
)

func TestListQueueEnqueue(t *testing.T) {
	testCases := []struct {
		name     string
		testCase func(*testing.T)
	}{
		{
			"test_enqueue_on_empty_queue", func(t *testing.T) {
				queue := NewListQueue[int]()

				queue.Enqueue(123)

				expectedQueue := newListQueueFromValues(123)
				if !isQueueEqual(expectedQueue, queue) {
					t.Fatalf("Queue %v not equal to expected %v", queue, expectedQueue)
				}
			},
		},
		{
			"test_enqueue_on_nonempty_queue", func(t *testing.T) {
				queue := newListQueueFromValues(123, 456)

				queue.Enqueue(789)

				expectedQueue := newListQueueFromValues(123, 456, 789)
				if !isQueueEqual(expectedQueue, queue) {
					t.Fatalf("Queue %v not equal to expected %v", queue, expectedQueue)
				}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, tc.testCase)
	}
}

func TestListQueueDnqueue(t *testing.T) {
	testCases := []struct {
		name     string
		testCase func(*testing.T)
	}{
		{
			"test_dequeue_on_empty_queue", func(t *testing.T) {
				queue := NewListQueue[int]()

				defer func() {
					if r := recover(); r == nil {
						log.Fatalf("Did not panic!!")
					}
				}()
				queue.Dequeue()
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, tc.testCase)
	}
}

func newListQueueFromValues[T any](args ...T) *ListQueue[T] {
	queue := NewListQueue[T]()
	for i, arg := range args {
		listNode := &ListNode[T]{
			val:  arg,
			next: nil,
		}

		if i == 0 {
			queue.front = listNode
			queue.back = listNode
			continue
		}

		queue.back.next = listNode
	}

	return queue
}

func isQueueEqual[T comparable](expected, actual *ListQueue[T]) bool {
	curExpected := expected.front
	curActual := actual.front
	for curExpected != nil && curActual != nil {
		if curExpected.val != curActual.val {
			return false
		}

		curExpected = curExpected.next
		curActual = curActual.next
	}

	return curExpected == nil && curActual == nil
}
