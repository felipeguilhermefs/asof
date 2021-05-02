package list

import (
	"math/rand"
	"testing"
	"time"
)

// TestLinkedListAppend
func TestLinkedListAppend(t *testing.T) {
	seed := time.Now().UnixNano()
	rand.Seed(seed)

	for name, list := range map[string][]int{
		"random-len10":   rand.Perm(10),
		"random-len1000": rand.Perm(1000),
	} {
		t.Run(name, func(t *testing.T) {
			linkedList := LinkedList{}

			for _, item := range list {
				linkedList.Append(item)
			}

			if linkedList.Len() != len(list) {
				t.Fatalf("Append length error: got %d; want %d", linkedList.Len(), len(list))
			}

			for i, item := range list {
				if !linkedList.Contains(item) {
					t.Fatalf("Append insertion error: index %d; item %d", i, item)
				}
			}

		})
	}
}

// TestLinkedListDelete
func TestLinkedListDelete(t *testing.T) {

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

		if linkedList.Len() != 3 {
			t.Fatalf("Delete length error: got %d; want 3", linkedList.Len())
		}
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

		if linkedList.Len() != 3 {
			t.Fatalf("Delete length error: got %d; want 3", linkedList.Len())
		}
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

		if linkedList.Len() != 4 {
			t.Fatalf("Delete length error: got %d; want 4", linkedList.Len())
		}
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

		if linkedList.Len() != 4 {
			t.Fatalf("Delete length error: got %d; want 4", linkedList.Len())
		}
	})

	t.Run("should Traverse from head to tail", func(t *testing.T) {
		linkedList := LinkedList{}
		linkedList.Append(2)
		linkedList.Append(5)
		linkedList.Append(1)
		linkedList.Append(4)
		linkedList.Append(1)

		visited := []int{}
		linkedList.Traverse(func(item int) {
			visited = append(visited, item)
		})

		expected := []int{2, 5, 1, 4, 1}
		for i := range expected {
			if visited[i] != expected[i] {
				t.Fatalf("Unexpected: index= %d expected= %v visited= %v", i, expected, visited)
			}
		}
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
}
