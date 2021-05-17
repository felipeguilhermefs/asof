package linkedlist

import "testing"

//TestDistinct
func TestDistinct(t *testing.T) {

	t.Run("should return emtpy when empty", func(t *testing.T) {
		distinct := distinct(&LinkedList{})
		if distinct.Len() != 0 {
			t.Fatalf("Should be empty, but len = %d", distinct.Len())
		}
	})

	t.Run("should return emtpy when nil", func(t *testing.T) {
		distinct := distinct(nil)
		if distinct.Len() != 0 {
			t.Fatalf("Should be empty, but len = %d", distinct.Len())
		}
	})

	t.Run("should return only distinct elements", func(t *testing.T) {
		testList := LinkedList{}
		testList.PushRight(1)
		testList.PushRight(1)
		testList.PushRight(1)
		testList.PushRight(12)
		testList.PushRight(2)

		assertSameOrder(t, distinct(&testList), []int{1, 12, 2})
	})
}
