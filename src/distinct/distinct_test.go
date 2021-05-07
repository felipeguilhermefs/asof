package distinct

import (
	"testing"

	"github.com/felipeguilhermefs/asof/datastructures/linkedlist"
)

//TestDistinctSingleLinkedList
func TestDistinctSingleLinkedList(t *testing.T) {

	t.Run("should return emtpy when empty", func(t *testing.T) {
		distinct := DistinctSingleLinkedList(&linkedlist.SingleLinkedList{})
		if distinct.Len() != 0 {
			t.Fatalf("Should be empty, but len = %d", distinct.Len())
		}
	})

	t.Run("should return emtpy when nil", func(t *testing.T) {
		distinct := DistinctSingleLinkedList(nil)
		if distinct.Len() != 0 {
			t.Fatalf("Should be empty, but len = %d", distinct.Len())
		}
	})

	t.Run("should return only distinct elements", func(t *testing.T) {
		testList := linkedlist.SingleLinkedList{}
		testList.Append(1)
		testList.Append(1)
		testList.Append(1)
		testList.Append(12)
		testList.Append(2)

		distinct := DistinctSingleLinkedList(&testList)
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
