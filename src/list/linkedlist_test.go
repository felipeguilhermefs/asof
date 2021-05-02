package list

import "testing"

// TestLinkedList
func TestLinkedList(t *testing.T) {

	t.Run("should append to the end of the list", func(t *testing.T) {
		linkedList := LinkedList{}
		linkedList.Append(1)
		linkedList.Append(2)
		linkedList.Append(3)
		linkedList.Append(4)

		assertSameOrder(t, &linkedList, []int{1, 2, 3, 4})
	})

	t.Run("should prepend to the start of the list", func(t *testing.T) {
		linkedList := LinkedList{}
		linkedList.Prepend(1)
		linkedList.Prepend(2)
		linkedList.Prepend(3)
		linkedList.Prepend(4)

		assertSameOrder(t, &linkedList, []int{4, 3, 2, 1})
	})

	t.Run("should remove from begining", func(t *testing.T) {
		linkedList := LinkedList{}
		linkedList.Append(1)
		linkedList.Append(2)
		linkedList.Append(3)
		linkedList.Append(4)

		linkedList.Delete(1)

		if linkedList.Contains(1) {
			t.Fatalf("Delete error: found item")
		}

		assertSameOrder(t, &linkedList, []int{2, 3, 4})
	})

	t.Run("should remove from end", func(t *testing.T) {
		linkedList := LinkedList{}
		linkedList.Append(1)
		linkedList.Append(2)
		linkedList.Append(3)
		linkedList.Append(4)

		linkedList.Delete(4)

		if linkedList.Contains(4) {
			t.Fatalf("Delete error: found item")
		}

		assertSameOrder(t, &linkedList, []int{1, 2, 3})
	})

	t.Run("should remove from middle", func(t *testing.T) {
		linkedList := LinkedList{}
		linkedList.Append(1)
		linkedList.Append(2)
		linkedList.Append(3)
		linkedList.Append(4)
		linkedList.Append(5)

		linkedList.Delete(3)

		if linkedList.Contains(3) {
			t.Fatalf("Delete error: found item")
		}

		assertSameOrder(t, &linkedList, []int{1, 2, 4, 5})
	})

	t.Run("should not remove if nothing is found", func(t *testing.T) {
		linkedList := LinkedList{}
		linkedList.Append(1)
		linkedList.Append(2)
		linkedList.Append(3)
		linkedList.Append(4)

		linkedList.Delete(5)

		if linkedList.Contains(5) {
			t.Fatalf("Delete error: found item")
		}

		assertSameOrder(t, &linkedList, []int{1, 2, 3, 4})
	})

	t.Run("should keep pointers consistent after delete last item", func(t *testing.T) {
		linkedList := LinkedList{}
		linkedList.Append(1)

		linkedList.Append(2)
		linkedList.Delete(2)

		linkedList.Append(3)

		assertSameOrder(t, &linkedList, []int{1, 3})
	})

	t.Run("should not visit any item in Traverse, if list is empty", func(t *testing.T) {
		linkedList := LinkedList{}

		visited := []int{}
		linkedList.Traverse(func(item int) {
			visited = append(visited, item)
		})

		if len(visited) > 0 {
			t.Fatalf("Found something: len= %d visited= %v", len(visited), visited)
		}
	})

	t.Run("should empty the list", func(t *testing.T) {
		linkedList := LinkedList{}
		linkedList.Append(1)
		linkedList.Append(2)
		linkedList.Append(3)

		linkedList.Clear()

		assertSameOrder(t, &linkedList, []int{})
	})
}

func assertSameOrder(t *testing.T, linkedList *LinkedList, expected []int) {
	if linkedList.Len() != len(expected) {
		t.Fatalf("Length error: got %d; want %d", linkedList.Len(), len(expected))
	}

	visited := []int{}
	linkedList.Traverse(func(item int) {
		visited = append(visited, item)
	})

	for i := range expected {
		if visited[i] != expected[i] {
			t.Fatalf("Unexpected: index= %d expected= %v visited= %v", i, expected, visited)
		}
	}
}
