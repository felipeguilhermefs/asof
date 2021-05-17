package linkedlist

import "testing"

// TestLinkedListMergesort
func TestLinkedListMergesort(t *testing.T) {

	t.Run("Mergesort should sort", func(t *testing.T) {
		linkedList := LinkedList{}
		linkedList.PushRight(5)
		linkedList.PushRight(2)
		linkedList.PushRight(4)
		linkedList.PushRight(1)
		linkedList.PushRight(6)
		linkedList.PushRight(3)
		linkedList.PushRight(7)

		linkedList.start = mergesort(linkedList.start)

		linkedList.end = findEnd(linkedList.start)

		assertSameOrder(t, &linkedList, []int{1, 2, 3, 4, 5, 6, 7})
	})

	t.Run("Mergesort should sort reversed list", func(t *testing.T) {
		linkedList := LinkedList{}
		linkedList.PushRight(7)
		linkedList.PushRight(6)
		linkedList.PushRight(5)
		linkedList.PushRight(4)
		linkedList.PushRight(3)
		linkedList.PushRight(2)
		linkedList.PushRight(1)

		linkedList.start = mergesort(linkedList.start)

		linkedList.end = findEnd(linkedList.start)

		assertSameOrder(t, &linkedList, []int{1, 2, 3, 4, 5, 6, 7})
	})

	t.Run("Mergesort should sort already sorted list", func(t *testing.T) {
		linkedList := LinkedList{}
		linkedList.PushRight(1)
		linkedList.PushRight(2)
		linkedList.PushRight(3)
		linkedList.PushRight(4)
		linkedList.PushRight(5)
		linkedList.PushRight(6)
		linkedList.PushRight(7)

		linkedList.start = mergesort(linkedList.start)

		linkedList.end = findEnd(linkedList.start)

		assertSameOrder(t, &linkedList, []int{1, 2, 3, 4, 5, 6, 7})
	})

	t.Run("Mergesort should sort with equal elements", func(t *testing.T) {
		linkedList := LinkedList{}
		linkedList.PushRight(1)
		linkedList.PushRight(1)
		linkedList.PushRight(1)
		linkedList.PushRight(1)

		linkedList.start = mergesort(linkedList.start)

		linkedList.end = findEnd(linkedList.start)

		assertSameOrder(t, &linkedList, []int{1, 1, 1, 1})
	})
}

// mergesort is messing with the pointers so the end pointer in the list
// is messed up after the sort and we need to fix it
func findEnd(start *node) *node {
	current := start
	for current.next != nil {
		current = current.next
	}
	return current
}
