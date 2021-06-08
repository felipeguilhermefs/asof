package linkedlist

import "github.com/felipeguilhermefs/asof/datastructures/set"

// distinct returns a new LinkedList with distinct elements
// for a linked list we can do this by keeping a collection of seen elements
// and checking against that. In this case we used the Set datastructure
// O(n) Time complexity. O(n) space complexity.
func distinct(list *LinkedList) *LinkedList {
	distinct := LinkedList{}
	if list == nil {
		return &distinct
	}

	seen := set.Set{}

	current := list.start
	for current != nil {
		if !seen.Contains(current.value) {
			seen.Add(current.value)
			distinct.PushRight(current.value)
		}

		current = current.next
	}

	return &distinct
}
