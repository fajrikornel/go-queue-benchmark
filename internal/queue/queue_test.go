package queue

import (
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
					t.Fatalf("Queue front not 1")
				}
				if queue.back != 0 {
					t.Fatalf("Queue back not 1")
				}
				if queue.queue[0] != val {
					t.Fatalf("Queue value not %d", val)
				}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, tc.testCase)
	}
}
