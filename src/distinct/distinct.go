package distinct

import (
	"github.com/felipeguilhermefs/asof/datastructures/linkedlist"
	"github.com/felipeguilhermefs/asof/datastructures/set"
)

//DistinctSingleLinkedList returns a new SingleLinkedList with distinct elements
// for a single linked list we can do this by keeping a collection of seen elements
// and checking against that. In this case we used the Set datastructure
func DistinctSingleLinkedList(list *linkedlist.SingleLinkedList) *linkedlist.SingleLinkedList {
	distinct := linkedlist.SingleLinkedList{}
	if list == nil {
		return &distinct
	}

	seen := set.Set{}

	list.Traverse(func(value int) bool {

		if !seen.Contains(value) {
			seen.Add(value)
			distinct.Append(value)
		}

		return true;
	})

	return &distinct
}
