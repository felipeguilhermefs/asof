package stack

import "testing"

// TestSliceStack run default stack tests to slice stack
func TestSliceStack(t *testing.T) {
	testStack(t, &SliceStack{})
}
