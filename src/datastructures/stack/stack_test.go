package stack

import (
	"math/rand"
	"testing"
	"time"
)

// testStack general stack test
func testStack(t *testing.T, stack Stack) {
	seed := time.Now().UnixNano()
	rand.Seed(seed)

	for name, list := range map[string][]int{
		"random-len10":   rand.Perm(10),
		"random-len1000": rand.Perm(1000),
	} {
		t.Run(name, func(t *testing.T) {

			shouldPushAll(t, stack, list)

			shouldPopAllFromTop(t, stack, list)

			shouldNotifyWhenEmpty(t, stack)

		})
	}
}

func shouldPushAll(t *testing.T, stack Stack, list []int) {
	for _, item := range list {
		stack.Push(item)
	}

	if stack.Len() != len(list) {
		t.Fatalf("Push error: got %d; want %d", stack.Len(), len(list))
	}
}

func shouldPopAllFromTop(t *testing.T, stack Stack, list []int) {
	for i := len(list) - 1; i >= 0; i-- {

		item, ok := stack.Pop()

		if !ok || item != list[i] {
			t.Fatalf("Pop error: ok %v; index %d", ok, i)
		}
	}

	if stack.Len() != 0 {
		t.Fatalf("Pop error, not empty: got %d; want 0", stack.Len())
	}
}

func shouldNotifyWhenEmpty(t *testing.T, stack Stack) {
	_, ok := stack.Pop()
	if ok {
		t.Fatalf("Pop empty error: should be empty")
	}
}
