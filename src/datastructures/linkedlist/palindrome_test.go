package linkedlist

import "testing"

// TestPalindrome
func TestPalindrome(t *testing.T) {
	t.Run("should be palindrome", func(t *testing.T) {
		l1 := newTestList([]int{1, 2, 1})
		if !palindrome(l1) {
			t.Fatalf("Should be palindrome: %v", l1)
		}

		l2 := newTestList([]int{5, 5})
		if !palindrome(l2) {
			t.Fatalf("Should be palindrome: %v", l2)
		}

		l3 := newTestList([]int{4, 2, 3, 1, 3, 2, 4})
		if !palindrome(l3) {
			t.Fatalf("Should be palindrome: %v", l3)
		}

		l4 := newTestList([]int{})
		if !palindrome(l4) {
			t.Fatalf("Should be palindrome: %v", l4)
		}

		l5 := newTestList([]int{9})
		if !palindrome(l5) {
			t.Fatalf("Should be palindrome: %v", l5)
		}
	})

	t.Run("should NOT be palindrome", func(t *testing.T) {
		l1 := newTestList([]int{1, 2, 2})
		if palindrome(l1) {
			t.Fatalf("Should NOT be palindrome: %v", l1)
		}

		l2 := newTestList([]int{1, 2, 3, 4, 5})
		if palindrome(l2) {
			t.Fatalf("Should NOT be palindrome: %v", l2)
		}

		l3 := newTestList([]int{6, 3, 2, 8})
		if palindrome(l3) {
			t.Fatalf("Should NOT be palindrome: %v", l3)
		}
	})
}

func newTestList(data []int) *LinkedList {
	list := LinkedList{}

	for _, val := range data {
		list.PushRight(val)
	}

	return &list
}
