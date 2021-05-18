package linkedlist

import "testing"

// TestMergesort
func TestMergesort(t *testing.T) {

	t.Run("Mergesort should sort", func(t *testing.T) {
		testList := newList([]int{5, 2, 4, 1, 6, 3, 7})

		testList.start = mergesort(testList.start)

		testList.end = findEnd(testList.start)

		assertSameOrder(t, testList, []int{1, 2, 3, 4, 5, 6, 7})
	})

	t.Run("Mergesort should sort reversed list", func(t *testing.T) {
		testList := newList([]int{7, 6, 5, 4, 3, 2, 1})

		testList.start = mergesort(testList.start)

		testList.end = findEnd(testList.start)

		assertSameOrder(t, testList, []int{1, 2, 3, 4, 5, 6, 7})
	})

	t.Run("Mergesort should sort already sorted list", func(t *testing.T) {
		testList := newList([]int{1, 2, 3, 4, 5, 6, 7})

		testList.start = mergesort(testList.start)

		testList.end = findEnd(testList.start)

		assertSameOrder(t, testList, []int{1, 2, 3, 4, 5, 6, 7})
	})

	t.Run("Mergesort should sort with equal elements", func(t *testing.T) {
		testList := newList([]int{1, 1, 1, 1})

		testList.start = mergesort(testList.start)

		testList.end = findEnd(testList.start)

		assertSameOrder(t, testList, []int{1, 1, 1, 1})
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
