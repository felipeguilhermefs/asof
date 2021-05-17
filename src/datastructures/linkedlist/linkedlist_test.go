package linkedlist

import "testing"

// TestLinkedList
func TestLinkedList(t *testing.T) {

	t.Run("PushRight should add to the end/right of the list", func(t *testing.T) {
		linkedList := LinkedList{}
		linkedList.PushRight(1)
		linkedList.PushRight(2)
		linkedList.PushRight(3)
		linkedList.PushRight(4)

		assertSameOrder(t, &linkedList, []int{1, 2, 3, 4})
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
		linkedList := LinkedList{}
		linkedList.PushLeft(1)
		linkedList.PushLeft(2)
		linkedList.PushLeft(3)
		linkedList.PushLeft(4)

		assertSameOrder(t, &linkedList, []int{4, 3, 2, 1})
	})

	t.Run("PopLeft should pop element at the start/left of the list", func(t *testing.T) {
		linkedList := LinkedList{}
		linkedList.PushRight(1)
		linkedList.PushRight(2)

		element, ok := linkedList.PopLeft()
		if !ok || element != 1 || linkedList.Len() != 1 {
			t.Fatalf("PopLeft error: expected 1, got %d, %v, %d", element, ok, linkedList.Len())
		}

		element, ok = linkedList.PopLeft()
		if !ok || element != 2 || linkedList.Len() != 0 {
			t.Fatalf("PopLeft error: expected 2, got %d, %v, %d", element, ok, linkedList.Len())
		}

		element, ok = linkedList.PopLeft()
		if ok {
			t.Fatalf("PopLeft error: expected empty list, got %d, %v", element, ok)
		}
	})

	t.Run("should remove from begining", func(t *testing.T) {
		linkedList := LinkedList{}
		linkedList.PushRight(1)
		linkedList.PushRight(2)
		linkedList.PushRight(3)
		linkedList.PushRight(4)

		linkedList.Delete(1)

		if linkedList.Contains(1) {
			t.Fatalf("Delete error: found item")
		}

		assertSameOrder(t, &linkedList, []int{2, 3, 4})
	})

	t.Run("should remove from end", func(t *testing.T) {
		linkedList := LinkedList{}
		linkedList.PushRight(1)
		linkedList.PushRight(2)
		linkedList.PushRight(3)
		linkedList.PushRight(4)

		linkedList.Delete(4)

		if linkedList.Contains(4) {
			t.Fatalf("Delete error: found item")
		}

		assertSameOrder(t, &linkedList, []int{1, 2, 3})
	})

	t.Run("should remove from middle", func(t *testing.T) {
		linkedList := LinkedList{}
		linkedList.PushRight(1)
		linkedList.PushRight(2)
		linkedList.PushRight(3)
		linkedList.PushRight(4)
		linkedList.PushRight(5)

		linkedList.Delete(3)

		if linkedList.Contains(3) {
			t.Fatalf("Delete error: found item")
		}

		assertSameOrder(t, &linkedList, []int{1, 2, 4, 5})
	})

	t.Run("should not remove if nothing is found", func(t *testing.T) {
		linkedList := LinkedList{}
		linkedList.PushRight(1)
		linkedList.PushRight(2)
		linkedList.PushRight(3)
		linkedList.PushRight(4)

		linkedList.Delete(5)

		if linkedList.Contains(5) {
			t.Fatalf("Delete error: found item")
		}

		assertSameOrder(t, &linkedList, []int{1, 2, 3, 4})
	})

	t.Run("should keep pointers consistent after delete last item", func(t *testing.T) {
		linkedList := LinkedList{}
		linkedList.PushRight(1)

		linkedList.PushRight(2)
		linkedList.Delete(2)

		linkedList.PushRight(3)

		assertSameOrder(t, &linkedList, []int{1, 3})
	})

	t.Run("should empty the list", func(t *testing.T) {
		linkedList := LinkedList{}
		linkedList.PushRight(1)
		linkedList.PushRight(2)
		linkedList.PushRight(3)

		linkedList.Clear()

		assertSameOrder(t, &linkedList, []int{})
	})
}

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
