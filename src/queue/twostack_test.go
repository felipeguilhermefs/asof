package queue

import (
	"math/rand"
	"testing"
	"time"
)

// TestTwoStackQueue
func TestTwoStackQueue(t *testing.T) {
	seed := time.Now().UnixNano()
	rand.Seed(seed)

	for name, list := range map[string][]int{
		"random-len10":   rand.Perm(10),
		"random-len1000": rand.Perm(1000),
	} {
		t.Run(name, func(t *testing.T) {
			queue := TwoStackQueue{}

			for _, item := range list {
				queue.Enqueue(item)
			}

			if queue.Len() != len(list) {
				t.Fatalf("Enqueue error: got %d; want %d", queue.Len(), len(list))
			}

			for i := range list {
				item, ok := queue.Dequeue()
				if !ok || item != list[i] {
					t.Fatalf("Dequeue error: ok %v; index %d", ok, i)
				}
			}

			_, ok := queue.Dequeue()
			if ok || queue.Len() != 0 {
				t.Fatalf("Dequeue empty error: should be empty")
			}

		})
	}
}
