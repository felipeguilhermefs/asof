package linkedlist

import "testing"

// TestQuicksort
func TestQuicksort(t *testing.T) {

	t.Run("Quicksort should sort", func(t *testing.T) {
		testList := newList([]int{5, 2, 4, 1, 6, 3, 7})

		quicksort(testList.start, testList.end)

		assertSameOrder(t, testList, []int{1, 2, 3, 4, 5, 6, 7})
	})

	t.Run("Quicksort should sort reversed list", func(t *testing.T) {
		testList := newList([]int{7, 6, 5, 4, 3, 2, 1})

		quicksort(testList.start, testList.end)

		assertSameOrder(t, testList, []int{1, 2, 3, 4, 5, 6, 7})
	})

	t.Run("Quicksort should sort already sorted list", func(t *testing.T) {
		testList := newList([]int{1, 2, 3, 4, 5, 6, 7})

		quicksort(testList.start, testList.end)

		assertSameOrder(t, testList, []int{1, 2, 3, 4, 5, 6, 7})
	})

	t.Run("Quicksort should sort with equal elements", func(t *testing.T) {
		testList := newList([]int{1, 1, 1, 1})

		quicksort(testList.start, testList.end)

		assertSameOrder(t, testList, []int{1, 1, 1, 1})
	})
}
