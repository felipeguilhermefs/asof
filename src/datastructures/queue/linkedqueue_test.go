package queue

import "testing"

// TestLinkedQueue
func TestLinkedQueue(t *testing.T) {
	testQueue(t, &LinkedQueue{})
}
