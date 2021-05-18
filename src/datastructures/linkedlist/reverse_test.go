package linkedlist

import "testing"

// TestReverse
func TestReverse(t *testing.T) {

	t.Run("Reverse should consider nil", func(t *testing.T) {
		var testList *LinkedList

		reverse(testList)

		if testList != nil {
			t.Fatal("should be nil")
		}
	})

	t.Run("Reverse empty", func(t *testing.T) {
		testList := &LinkedList{}

		reverse(testList)

		assertSameOrder(t, testList, []int{})

		if testList.start != nil || testList.end != nil {
			t.Fatal("should be empty")
		}
	})

	t.Run("Reverse single item", func(t *testing.T) {
		testList := newList([]int{1})

		reverse(testList)

		assertSameOrder(t, testList, []int{1})

		if testList.start.value != 1 || testList.end.value != 1 {
			t.Fatal("should be 1")
		}
	})

	t.Run("Reverse a list", func(t *testing.T) {
		testList := newList([]int{1, 2, 3, 4, 5, 6, 7})

		reverse(testList)

		assertSameOrder(t, testList, []int{7, 6, 5, 4, 3, 2, 1})

		if testList.start.value != 7 {
			t.Fatal("start should be 7")
		}

		if testList.end.value != 1 {
			t.Fatal("end should be 1")
		}
	})
}
