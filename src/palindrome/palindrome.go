package palindrome

import (
	"github.com/felipeguilhermefs/asof/datastructures/linkedlist"
	"github.com/felipeguilhermefs/asof/datastructures/stack"
)

// IsPalindrome check if the SingleLinkedList is a palindrome
func IsPalindrome(list *linkedlist.SingleLinkedList) bool {
	s := stack.LinkedStack{}

	list.Traverse(func(value int) {
		s.Push(value)
	})

	check := true
	list.Traverse(func(value int) {
		if sValue, ok := s.Pop(); !ok || sValue != value {
			check = false
		}
	})

	return check
}
