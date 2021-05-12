package linkedlist

func (ll *LinkedList) partition(start, end *node) *node {
	if start == end || start == nil || end == nil {
		return start
	}

	prev := start
	current := start

	for start != end {
		if start.value < end.value {
			prev = current
			start.value, current.value = current.value, start.value
			current = current.next
		}

		start = start.next
	}

	current.value, end.value = end.value, current.value

	return prev
}

func (ll *LinkedList) sort(start, end *node) {
	if start == nil || start == end {
		return
	}

	pivot := ll.partition(start, end)

	// sort first half
	ll.sort(start, pivot)

	// sort last half
	ll.sort(pivot.next, end)
}
