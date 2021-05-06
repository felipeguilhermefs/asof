package linkedlist

import "testing"

// TestSingleLinkedList
func TestSingleLinkedList(t *testing.T) {
	testLinkedList(t, func() LinkedList {
		return &SingleLinkedList{}
	})
}
