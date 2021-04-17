package stack

import "testing"

// TestLinkedStack run default stack tests to linked list stack
func TestLinkedStack(t *testing.T) {
	testStack(t, &LinkedStack{})
}
