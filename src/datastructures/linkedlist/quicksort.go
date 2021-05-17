package linkedlist

func partition(start, end *node) *node {
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

//quicksort inplace a linked list. O(nlgn) time complexity
func quicksort(start, end *node) {
	if start == nil || start == end {
		return
	}

	pivot := partition(start, end)

	quicksort(start, pivot)

	quicksort(pivot.next, end)
}
