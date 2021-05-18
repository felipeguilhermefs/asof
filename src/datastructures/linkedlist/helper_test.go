package linkedlist

import "testing"

func assertSameOrder(t *testing.T, linkedList *LinkedList, expected []int) {
	if linkedList.Len() != len(expected) {
		t.Fatalf("Length error: got %d; want %d", linkedList.Len(), len(expected))
	}

	current := linkedList.start
	for i := range expected {
		if current.value != expected[i] {
			t.Fatalf("Unexpected: index= %d expected= %v visited= %d", i, expected, current.value)
		}
		current = current.next
	}
}

func newList(data []int) *LinkedList {
	list := LinkedList{}

	for _, val := range data {
		list.PushRight(val)
	}

	return &list
}
