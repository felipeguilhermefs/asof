package stack

import (
	"math/rand"
	"testing"
	"time"
)

// TestStack
func TestStack(t *testing.T) {
	seed := time.Now().UnixNano()
	rand.Seed(seed)

	for name, list := range map[string][]int{
		"random-len10":   rand.Perm(10),
		"random-len1000": rand.Perm(1000),
	} {
		t.Run(name, func(t *testing.T) {
			stack := Stack{}

			for _, item := range list {
				stack.Push(item)
			}

			if stack.Len() != len(list) {
				t.Fatalf("Push error: got %d; want %d", stack.Len(), len(list))
			}

			for i := len(list) - 1; i >= 0; i-- {
				item, ok := stack.Pop()
				if !ok || item != list[i] {
					t.Fatalf("Pop error: ok %v; index %d", ok, i)
				}
			}

			_, ok := stack.Pop()
			if ok {
				t.Fatalf("Pop empty error: should be empty")
			}

		})
	}
}
