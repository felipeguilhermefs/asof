package queue

import "testing"

// TestTwoStackQueue
func TestTwoStackQueue(t *testing.T) {
	testQueue(t, &TwoStackQueue{})
}
