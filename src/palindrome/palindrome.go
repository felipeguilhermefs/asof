package palindrome

import (
	"github.com/felipeguilhermefs/asof/datastructures/linkedlist"
	"github.com/felipeguilhermefs/asof/datastructures/stack"
)

// IsPalindrome check if the SingleLinkedList is a palindrome
func IsPalindrome(list *linkedlist.SingleLinkedList) bool {
	s := stack.NewStack()

	list.Traverse(func(value int) bool {
		s.Push(value)

		return true
	})

	check := true
	list.Traverse(func(value int) bool {

		if sValue, ok := s.Pop(); !ok || sValue != value {
			check = false
		}

		return check
	})

	return check
}
