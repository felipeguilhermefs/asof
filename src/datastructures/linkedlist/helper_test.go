package linkedlist

import "testing"

func assertSameOrder(t *testing.T, linkedList *LinkedList, expected []int) {
	if linkedList.Len() != len(expected) {
		t.Fatalf("Length error: got %d; want %d", linkedList.Len(), len(expected))
	}

	// verify left to right
	current := linkedList.start
	for i := range expected {
		if current.value != expected[i] {
			t.Fatalf("L2R Unexpected: index= %d expected= %v visited= %d", i, expected, current.value)
		}
		current = current.next
	}

	// verify right to left
	current = linkedList.end
	for i := len(expected) - 1; i >= 0; i-- {
		if current.value != expected[i] {
			t.Fatalf("R2L Unexpected: index= %d expected= %v visited= %d", i, expected, current.value)
		}
		current = current.previous
	}
}

func newList(data []int) *LinkedList {
	list := LinkedList{}

	for _, val := range data {
		list.PushRight(val)
	}

	return &list
}
