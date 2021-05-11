package distinct

import (
	"testing"

	"github.com/felipeguilhermefs/asof/datastructures/linkedlist"
)

//TestDistinctLinkedList
func TestDistinctLinkedList(t *testing.T) {

	t.Run("should return emtpy when empty", func(t *testing.T) {
		distinct := DistinctLinkedList(&linkedlist.LinkedList{})
		if distinct.Len() != 0 {
			t.Fatalf("Should be empty, but len = %d", distinct.Len())
		}
	})

	t.Run("should return emtpy when nil", func(t *testing.T) {
		distinct := DistinctLinkedList(nil)
		if distinct.Len() != 0 {
			t.Fatalf("Should be empty, but len = %d", distinct.Len())
		}
	})

	t.Run("should return only distinct elements", func(t *testing.T) {
		testList := linkedlist.LinkedList{}
		testList.PushRight(1)
		testList.PushRight(1)
		testList.PushRight(1)
		testList.PushRight(12)
		testList.PushRight(2)

		distinct := DistinctLinkedList(&testList)
		if distinct.Len() != 3 {
			t.Fatalf("Should have length 3, but len = %d", distinct.Len())
		}

		visited := []int{}
		distinct.Traverse(func(value int) bool {
			visited = append(visited, value)
			return true
		})
		if visited[0] != 1 || visited[1] != 12 || visited[2] != 2 {
			t.Fatalf("Visited elements different from expected, received %v", visited)
		}

	})
}
