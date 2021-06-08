package linkedlist

// quicksort inplace a linked list.
// O(nlgn) Time complexity
func quicksort(start, end *node) {
	if start == nil || start == end {
		return
	}

	pivot := partition(start, end)

	quicksort(start, pivot)

	quicksort(pivot.next, end)
}

func partition(start, end *node) *node {
	if start == end || start == nil || end == nil {
		return start
	}

	// prev holds the last element smaller than the pivot
	prev := start
	// current holds the position where smaller elements could be placed
	current := start
	// we are using the end node as pivot, so anything smaller
	// will be grouped at the start.
	for start != end {
		if start.value < end.value {
			prev = current
			start.value, current.value = current.value, start.value
			current = current.next
		}

		start = start.next
	}

	current.value, end.value = end.value, current.value

	// since we are just swapping values we just need to return prev
	// as it points to division point found by the chosen pivot
	return prev
}
