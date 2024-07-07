package queue

import (
	"reflect"
	"testing"
)

func TestEnqueue(t *testing.T) {
	testCases := []struct {
		name     string
		testCase func(*testing.T)
	}{
		{
			"test_enqueue_on_empty_queue", func(t *testing.T) {
				queue := New[int](1)

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

				queue := New[int](1)
				queue.front = 0
				queue.back = 0

				val := 123
				queue.Enqueue(val)
			},
		},
		{
			"test_enqueue_on_non_empty_queue", func(t *testing.T) {
				queue := New[int](2)
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
	}

	for _, tc := range testCases {
		t.Run(tc.name, tc.testCase)
	}
}
