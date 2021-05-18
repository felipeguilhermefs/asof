package linkedlist

import "testing"

// TestLinkedList
func TestLinkedList(t *testing.T) {

	t.Run("PushRight should add to the end/right of the list", func(t *testing.T) {
		testList := LinkedList{}
		testList.PushRight(1)
		testList.PushRight(2)
		testList.PushRight(3)
		testList.PushRight(4)

		assertSameOrder(t, &testList, []int{1, 2, 3, 4})
	})

	t.Run("PopRight should pop element at the end/right of the list", func(t *testing.T) {
		linkedList := LinkedList{}
		linkedList.PushRight(1)
		linkedList.PushRight(2)

		element, ok := linkedList.PopRight()
		if !ok || element != 2 || linkedList.Len() != 1 {
			t.Fatalf("PopRight error: expected 2, got %d, %v, %d", element, ok, linkedList.Len())
		}

		element, ok = linkedList.PopRight()
		if !ok || element != 1 || linkedList.Len() != 0 {
			t.Fatalf("PopRight error: expected 1, got %d, %v, %d", element, ok, linkedList.Len())
		}

		element, ok = linkedList.PopRight()
		if ok {
			t.Fatalf("PopRight error: expected empty list, got %d, %v", element, ok)
		}
	})

	t.Run("PushLeft should add to the start/left of the list", func(t *testing.T) {
		testList := LinkedList{}
		testList.PushLeft(1)
		testList.PushLeft(2)
		testList.PushLeft(3)
		testList.PushLeft(4)

		assertSameOrder(t, &testList, []int{4, 3, 2, 1})
	})

	t.Run("PopLeft should pop element at the start/left of the list", func(t *testing.T) {
		testList := LinkedList{}
		testList.PushRight(1)
		testList.PushRight(2)

		element, ok := testList.PopLeft()
		if !ok || element != 1 || testList.Len() != 1 {
			t.Fatalf("PopLeft error: expected 1, got %d, %v, %d", element, ok, testList.Len())
		}

		element, ok = testList.PopLeft()
		if !ok || element != 2 || testList.Len() != 0 {
			t.Fatalf("PopLeft error: expected 2, got %d, %v, %d", element, ok, testList.Len())
		}

		element, ok = testList.PopLeft()
		if ok {
			t.Fatalf("PopLeft error: expected empty list, got %d, %v", element, ok)
		}
	})

	t.Run("Delete should remove from begining", func(t *testing.T) {
		testList := newList([]int{1, 2, 3, 4})

		testList.Delete(1)

		if testList.Contains(1) {
			t.Fatalf("Delete error: found item")
		}

		assertSameOrder(t, testList, []int{2, 3, 4})
	})

	t.Run("Delete should remove from end", func(t *testing.T) {
		testList := newList([]int{1, 2, 3, 4})

		testList.Delete(4)

		if testList.Contains(4) {
			t.Fatalf("Delete error: found item")
		}

		assertSameOrder(t, testList, []int{1, 2, 3})
	})

	t.Run("Delete should remove from middle", func(t *testing.T) {
		testList := newList([]int{1, 2, 3, 4, 5})

		testList.Delete(3)

		if testList.Contains(3) {
			t.Fatalf("Delete error: found item")
		}

		assertSameOrder(t, testList, []int{1, 2, 4, 5})
	})

	t.Run("Delete should not remove if nothing is found", func(t *testing.T) {
		testList := newList([]int{1, 2, 3, 4})

		testList.Delete(5)

		if testList.Contains(5) {
			t.Fatalf("Delete error: found item")
		}

		assertSameOrder(t, testList, []int{1, 2, 3, 4})
	})

	t.Run("Delete should keep pointers consistent after delete last item", func(t *testing.T) {
		testList := LinkedList{}
		testList.PushRight(1)

		testList.PushRight(2)
		testList.Delete(2)

		testList.PushRight(3)

		assertSameOrder(t, &testList, []int{1, 3})
	})

	t.Run("Clear should empty the list", func(t *testing.T) {
		testList := newList([]int{1, 2, 3})

		testList.Clear()

		assertSameOrder(t, testList, []int{})
	})
}
