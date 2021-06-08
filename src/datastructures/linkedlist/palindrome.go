package linkedlist

import "github.com/felipeguilhermefs/asof/datastructures/stack"

// palindrome check if the LinkedList is a palindrome.
// A simple way to check if the singly linked list is a
// palindrome is using a stack, but we need to traverse it 2 times.
// O(n) time complexity. O(n) space complexity.
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
