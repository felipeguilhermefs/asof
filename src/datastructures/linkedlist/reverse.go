package linkedlist

// reverse the linked list inplace.
// O(n) Time Complextity
func reverse(list *LinkedList) {
	if list == nil {
		return
	}

	current := list.start
	for current != nil {
		next := current.next
		current.next, current.previous = current.previous, current.next
		current = next
	}

	list.end, list.start = list.start, list.end
}
