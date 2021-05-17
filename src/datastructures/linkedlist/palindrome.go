package linkedlist

import "github.com/felipeguilhermefs/asof/datastructures/stack"

// palindrome check if the LinkedList is a palindrome
func palindrome(list *LinkedList) bool {
	s := stack.NewStack()

	current := list.start
	for current != nil {
		s.Push(current.value)
		current = current.next
	}

	current = list.start
	for current != nil {
		sValue, ok := s.Pop()
		if !ok || sValue != current.value {
			return false
		}
		current = current.next
	}

	return true
}
