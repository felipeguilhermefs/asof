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
				linkedItem, ok := linkedList.Index(i)
				if linkedItem != item || !ok {
					t.Fatalf("Append order error: got %d; want %d", linkedItem, item)
				}
			}

		})
	}
}

// TestLinkedListPrepend
func TestLinkedListPrepend(t *testing.T) {
	seed := time.Now().UnixNano()
	rand.Seed(seed)

	for name, list := range map[string][]int{
		"random-len10":   rand.Perm(10),
		"random-len1000": rand.Perm(1000),
	} {
		t.Run(name, func(t *testing.T) {
			linkedList := LinkedList{}

			for _, item := range list {
				linkedList.Prepend(item)
			}

			if linkedList.Len() != len(list) {
				t.Fatalf("Prepend length error: got %d; want %d", linkedList.Len(), len(list))
			}

			for i, item := range list {
				reverseIndex := len(list) - 1 - i
				linkedItem, ok := linkedList.Index(reverseIndex)
				if linkedItem != item || !ok {
					t.Fatalf("Prepend order error: got %d; want %d", linkedItem, item)
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

		if item, _ := linkedList.Index(0); item == 1 {
			t.Fatalf("Delete error: found item")
		}

		if linkedList.Len() != 3 {
			t.Fatalf("Delete length error: got %d; want 3", linkedList.Len())
		}

		if item, _ := linkedList.Index(0); item != 2 {
			t.Fatalf("Delete order error: should have kept order")
		}
	})

	t.Run("should remove from end", func(t *testing.T) {
		linkedList := LinkedList{}
		linkedList.Append(1)
		linkedList.Append(2)
		linkedList.Append(3)
		linkedList.Append(4)

		linkedList.Delete(4)

		if _, ok := linkedList.Index(3); ok {
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

		if item, _ := linkedList.Index(2); item == 3 {
			t.Fatalf("Delete error: found item")
		}

		if linkedList.Len() != 4 {
			t.Fatalf("Delete length error: got %d; want 4", linkedList.Len())
		}

		if item, _ := linkedList.Index(2); item != 4 {
			t.Fatalf("Delete order error: should have kept order")
		}
	})
}
