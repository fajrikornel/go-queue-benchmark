package queue

import (
	"reflect"
	"testing"
)

func TestArrayQueueEnqueue(t *testing.T) {
	testCases := []struct {
		name     string
		testCase func(*testing.T)
	}{
		{
			"test_enqueue_on_empty_queue", func(t *testing.T) {
				queue := NewArrayQueue[int](1)

				val := 123
				queue.Enqueue(val)

				if queue.front != 0 {
					t.Fatalf("Queue front not 0")
				}
				if queue.back != 0 {
					t.Fatalf("Queue back not 0")
				}
				if queue.queue[0] != val {
					t.Fatalf("Queue value not %d", val)
				}
			},
		},
		{
			"test_enqueue_on_full_queue", func(t *testing.T) {
				defer func() {
					if r := recover(); r == nil {
						t.Fatal("Did not panic!!")
					}
				}()

				queue := NewArrayQueue[int](1)
				queue.front = 0
				queue.back = 0

				val := 123
				queue.Enqueue(val)
			},
		},
		{
			"test_enqueue_on_non_empty_queue", func(t *testing.T) {
				queue := NewArrayQueue[int](2)
				queue.front = 0
				queue.back = 0
				queue.queue[0] = 321

				val := 123
				queue.Enqueue(val)

				if queue.front != 1 {
					t.Fatalf("Queue front not 1")
				}
				if queue.back != 0 {
					t.Fatalf("Queue back not 0")
				}
				if !reflect.DeepEqual(queue.queue, []int{321, 123}) {
					t.Fatalf("Queue value not expected: %v", queue.queue)
				}
			},
		},
		{
			"test_enqueue_on_edge_of_array", func(t *testing.T) {
				queue := NewArrayQueue[int](2)
				queue.front = 1
				queue.back = 1
				queue.queue[1] = 321

				val := 123
				queue.Enqueue(val)

				if queue.front != 0 {
					t.Fatalf("Queue front not 0")
				}
				if queue.back != 1 {
					t.Fatalf("Queue back not 1")
				}
				if !reflect.DeepEqual(queue.queue, []int{123, 321}) {
					t.Fatalf("Queue value not expected: %v", queue.queue)
				}
			},
		},
		{
			"test_enqueue_on_enqueuable_queue_when_back_is_at_edge_of_array", func(t *testing.T) {
				queue := NewArrayQueue[int](3)
				queue.front = 0
				queue.back = 2
				queue.queue[0] = 321
				queue.queue[2] = 321

				val := 123
				queue.Enqueue(val)

				if queue.front != 1 {
					t.Fatalf("Queue front not 1")
				}
				if queue.back != 2 {
					t.Fatalf("Queue back not 2")
				}
				if !reflect.DeepEqual(queue.queue, []int{321, 123, 321}) {
					t.Fatalf("Queue value not expected: %v", queue.queue)
				}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, tc.testCase)
	}
}
func TestArrayQueueDequeue(t *testing.T) {
	testCases := []struct {
		name     string
		testCase func(*testing.T)
	}{
		{
			"test_dequeue_on_empty_queue_never_enqueued", func(t *testing.T) {
				defer func() {
					if r := recover(); r == nil {
						t.Fatalf("Did not panic!!")
					}
				}()

				queue := NewArrayQueue[int](1)

				queue.Dequeue()
			},
		},
		{
			"test_dequeue_when_queue_size_is_one", func(t *testing.T) {
				queue := NewArrayQueue[int](2)
				queue.front = 0
				queue.back = 0
				queue.queue[0] = 123

				actualValue := queue.Dequeue()

				if queue.front != -1 {
					t.Fatalf("Queue front not -1")
				}
				if queue.back != -1 {
					t.Fatalf("Queue back not -1")
				}
				if actualValue != 123 {
					t.Fatalf("Dequeue value not expected: %v", actualValue)
				}
			},
		},
		{
			"test_dequeue_when_queue_size_is_not_one", func(t *testing.T) {
				queue := NewArrayQueue[int](2)
				queue.front = 1
				queue.back = 0
				queue.queue[0] = 123
				queue.queue[1] = 321

				actualValue := queue.Dequeue()

				if queue.front != 1 {
					t.Fatalf("Queue front not 1")
				}
				if queue.back != 1 {
					t.Fatalf("Queue back not 1")
				}
				if actualValue != 123 {
					t.Fatalf("Dequeue value not expected: %v", actualValue)
				}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, tc.testCase)
	}
}

func TestArrayQueueIntegration(t *testing.T) {
	testCases := []struct {
		name     string
		testCase func(*testing.T)
	}{
		{
			"test_queue_size_1", func(t *testing.T) {
				queue := NewArrayQueue[int](1)

				queue.Enqueue(123)
				if !reflect.DeepEqual(queue.queue, []int{123}) {
					t.Fatalf("Not expected!")
				}

				func() {
					defer func() {
						if r := recover(); r == nil {
							t.Fatalf("Did not panic!!")
						}
					}()

					queue.Enqueue(321)
				}()

				val := queue.Dequeue()
				if val != 123 {
					t.Fatalf("Not expected!")
				}

				func() {
					defer func() {
						if r := recover(); r == nil {
							t.Fatalf("Did not panic!!")
						}
					}()

					queue.Dequeue()
				}()

				queue.Enqueue(456)
				if !reflect.DeepEqual(queue.queue, []int{456}) {
					t.Fatalf("Not expected!")
				}

				val = queue.Dequeue()
				if val != 456 {
					t.Fatalf("Not expected!")
				}
			},
		},
		{
			"test_queue_size_3", func(t *testing.T) {
				queue := NewArrayQueue[int](3)

				queue.Enqueue(123)
				queue.Enqueue(456)
				queue.Enqueue(789)
				if !reflect.DeepEqual(queue.queue, []int{123, 456, 789}) {
					t.Fatalf("Not expected!")
				}

				func() {
					defer func() {
						if r := recover(); r == nil {
							t.Fatalf("Did not panic!!")
						}
					}()

					queue.Enqueue(321)
				}()

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
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, tc.testCase)
	}
}
