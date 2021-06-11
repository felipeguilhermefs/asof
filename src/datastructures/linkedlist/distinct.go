package linkedlist

import "github.com/felipeguilhermefs/asof/datastructures/set"

// distinct removes all duplicates from the linked list.
// we can do this by keeping a collection of seen elements
// and checking against that. In this case we used the Set datastructure
// O(n) Time complexity. O(n) space complexity.
func distinct(list *LinkedList) {
	if list == nil {
		return
	}

	seen := set.Set{}

	current := list.start
	for current != nil {
		if seen.Contains(current.value) {
			list.deleteNode(current)
		} else {
			seen.Add(current.value)
		}

		current = current.next
	}
}
