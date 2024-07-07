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

func TestListQueueDequeue(t *testing.T) {
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
		{
			"test_dequeue_on_nonempty_queue", func(t *testing.T) {
				queue := newListQueueFromValues(123, 456, 789)

				val := queue.Dequeue()
				if val != 123 {
					log.Fatalf("Expected value not 123!")
				}

				expectedQueue := newListQueueFromValues(456, 789)
				if !isQueueEqual(expectedQueue, queue) {
					log.Fatalf("Queue %v not equal to expected queue %v", queue, expectedQueue)
				}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, tc.testCase)
	}
}

func TestListQueueIntegration(t *testing.T) {
	testCases := []struct {
		name     string
		testCase func(*testing.T)
	}{
		{
			"test_list_queue", func(t *testing.T) {
				queue := NewListQueue[int]()

				queue.Enqueue(123)
				queue.Enqueue(456)
				queue.Enqueue(789)
				if !isQueueEqual(newListQueueFromValues(123, 456, 789), queue) {
					t.Fatalf("Not expected!")
				}

				val := queue.Dequeue()
				if val != 123 {
					t.Fatalf("Not expected!")
				}

				queue.Enqueue(321)

				val = queue.Dequeue()
				if val != 456 {
					t.Fatalf("Not expected!")
				}

				val = queue.Dequeue()
				if val != 789 {
					t.Fatalf("Not expected!")
				}

				val = queue.Dequeue()
				if val != 321 {
					t.Fatalf("Not expected!")
				}

				if !isQueueEqual(NewListQueue[int](), queue) {
					t.Fatalf("Not expected!")
				}
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
		queue.back = queue.back.next
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
