package distinct

import (
	"github.com/felipeguilhermefs/asof/datastructures/linkedlist"
	"github.com/felipeguilhermefs/asof/datastructures/set"
)

//DistinctLinkedList returns a new LinkedList with distinct elements
// for a linked list we can do this by keeping a collection of seen elements
// and checking against that. In this case we used the Set datastructure
func DistinctLinkedList(list *linkedlist.LinkedList) *linkedlist.LinkedList {
	distinct := linkedlist.LinkedList{}
	if list == nil {
		return &distinct
	}

	seen := set.Set{}

	list.Traverse(func(value int) bool {

		if !seen.Contains(value) {
			seen.Add(value)
			distinct.PushRight(value)
		}

		return true
	})

	return &distinct
}
