package linkedlist

import "testing"

//TestDistinct
func TestDistinct(t *testing.T) {

	t.Run("should return emtpy when empty", func(t *testing.T) {
		testList := &LinkedList{}
		distinct(testList)
		if testList.Len() != 0 {
			t.Fatalf("Should be empty, but len = %d", testList.Len())
		}
	})

	t.Run("should return emtpy when nil", func(t *testing.T) {
		var testList *LinkedList
		distinct(testList)
		if testList != nil {
			t.Fatalf("Should be nil, but received = %v", testList)
		}
	})

	t.Run("should return only distinct elements", func(t *testing.T) {
		testList := newList([]int{1, 1, 1, 12, 1, 2})

		distinct(testList)

		assertSameOrder(t, testList, []int{1, 12, 2})
	})
}
