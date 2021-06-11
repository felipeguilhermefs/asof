package linkedlist

import "testing"

// TestPalindrome
func TestPalindrome(t *testing.T) {
	t.Run("should be palindrome", func(t *testing.T) {
		l1 := newList([]int{1, 2, 1})
		if !palindrome(l1) {
			t.Fatalf("Should be palindrome: %v", l1)
		}

		l2 := newList([]int{5, 5})
		if !palindrome(l2) {
			t.Fatalf("Should be palindrome: %v", l2)
		}

		l3 := newList([]int{4, 2, 3, 1, 3, 2, 4})
		if !palindrome(l3) {
			t.Fatalf("Should be palindrome: %v", l3)
		}

		l4 := newList([]int{9})
		if !palindrome(l4) {
			t.Fatalf("Should be palindrome: %v", l4)
		}
	})

	t.Run("should NOT be palindrome", func(t *testing.T) {
		l1 := newList([]int{1, 2, 2})
		if palindrome(l1) {
			t.Fatalf("Should NOT be palindrome: %v", l1)
		}

		l2 := newList([]int{1, 2, 3, 4, 5})
		if palindrome(l2) {
			t.Fatalf("Should NOT be palindrome: %v", l2)
		}

		l3 := newList([]int{6, 3, 2, 8})
		if palindrome(l3) {
			t.Fatalf("Should NOT be palindrome: %v", l3)
		}

		l4 := newList([]int{})
		if palindrome(l4) {
			t.Fatalf("Should NOT be palindrome: %v", l4)
		}
	})
}
