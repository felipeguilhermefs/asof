package linkedlist

import "testing"

// TestQuicksort
func TestQuicksort(t *testing.T) {

	t.Run("Quicksort should sort", func(t *testing.T) {
		linkedList := LinkedList{}
		linkedList.PushRight(5)
		linkedList.PushRight(2)
		linkedList.PushRight(4)
		linkedList.PushRight(1)
		linkedList.PushRight(6)
		linkedList.PushRight(3)
		linkedList.PushRight(7)

		quicksort(linkedList.start, linkedList.end)

		assertSameOrder(t, &linkedList, []int{1, 2, 3, 4, 5, 6, 7})
	})

	t.Run("Quicksort should sort reversed list", func(t *testing.T) {
		linkedList := LinkedList{}
		linkedList.PushRight(7)
		linkedList.PushRight(6)
		linkedList.PushRight(5)
		linkedList.PushRight(4)
		linkedList.PushRight(3)
		linkedList.PushRight(2)
		linkedList.PushRight(1)

		quicksort(linkedList.start, linkedList.end)

		assertSameOrder(t, &linkedList, []int{1, 2, 3, 4, 5, 6, 7})
	})

	t.Run("Quicksort should sort already sorted list", func(t *testing.T) {
		linkedList := LinkedList{}
		linkedList.PushRight(1)
		linkedList.PushRight(2)
		linkedList.PushRight(3)
		linkedList.PushRight(4)
		linkedList.PushRight(5)
		linkedList.PushRight(6)
		linkedList.PushRight(7)

		quicksort(linkedList.start, linkedList.end)

		assertSameOrder(t, &linkedList, []int{1, 2, 3, 4, 5, 6, 7})
	})

	t.Run("Quicksort should sort with equal elements", func(t *testing.T) {
		linkedList := LinkedList{}
		linkedList.PushRight(1)
		linkedList.PushRight(1)
		linkedList.PushRight(1)
		linkedList.PushRight(1)

		quicksort(linkedList.start, linkedList.end)

		assertSameOrder(t, &linkedList, []int{1, 1, 1, 1})
	})
}
