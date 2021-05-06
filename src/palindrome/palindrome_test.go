package palindrome

import (
	"testing"

	"github.com/felipeguilhermefs/asof/datastructures/linkedlist"
)

// TestIsPalindrome
func TestIsPalindrome(t *testing.T) {
	t.Run("should be palindrome", func(t *testing.T) {
		l1 := newTestList([]int{1,2,1})
		if !IsPalindrome(l1) {
			t.Fatalf("Should be palindrome: %v", l1)
		}

		l2 := newTestList([]int{5,5})
		if !IsPalindrome(l2) {
			t.Fatalf("Should be palindrome: %v", l2)
		}

		l3 := newTestList([]int{4,2,3,1,3,2,4})
		if !IsPalindrome(l3) {
			t.Fatalf("Should be palindrome: %v", l3)
		}

		l4 := newTestList([]int{})
		if !IsPalindrome(l4) {
			t.Fatalf("Should be palindrome: %v", l4)
		}

		l5 := newTestList([]int{9})
		if !IsPalindrome(l5) {
			t.Fatalf("Should be palindrome: %v", l5)
		}
	})

	t.Run("should NOT be palindrome", func(t *testing.T) {
		l1 := newTestList([]int{1,2,2})
		if IsPalindrome(l1) {
			t.Fatalf("Should NOT be palindrome: %v", l1)
		}

		l2 := newTestList([]int{1,2,3,4,5})
		if IsPalindrome(l2) {
			t.Fatalf("Should NOT be palindrome: %v", l2)
		}

		l3 := newTestList([]int{6,3,2,8})
		if IsPalindrome(l3) {
			t.Fatalf("Should NOT be palindrome: %v", l3)
		}
	})
}

func newTestList(data []int) *linkedlist.SingleLinkedList {
	list := linkedlist.SingleLinkedList{}

	for _, val := range data {
		list.Append(val)
	}

	return &list
}
