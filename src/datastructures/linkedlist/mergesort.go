package linkedlist

// mergesort inplace a linked list.
// O(nlgn) time complexity
func mergesort(start *node) *node {
	if start == nil || start.next == nil {
		return start
	}

	middle := split(start)

	left := mergesort(start)
	right := mergesort(middle)

	res := merge(left, right)
	return res
}

func split(start *node) *node {
	if start == nil {
		return start
	}

	// to get the middle of a linked list we need 2 pointers
	// a slow pointer moves forward 1 step, and the fast 2 steps
	// when the fast reaches the end, the slow pointer will be
	// roughly in the middle
	slow := start
	fast := start
	for fast.next != nil && fast.next.next != nil {
		slow = slow.next
		fast = fast.next.next
	}

	middle := slow.next

	// to effectively split we should drop the borders pointers
	middle.previous = nil
	slow.next = nil

	return middle
}

func merge(left, right *node) *node {
	if left == nil {
		return right
	}

	if right == nil {
		return left
	}

	// we keep the start to represent the left most/smallest element
	var start *node

	if left.value < right.value {
		start = left
		left = left.next
	} else {
		start = right
		right = right.next
	}

	// current will hold the "current" smallest element
	// so we can interweave left and right nodes
	current := start
	for left != nil && right != nil {
		if left.value < right.value {
			current.next, left.previous = left, current
			current = left
			left = left.next
		} else {
			current.next, right.previous = right, current
			current = right
			right = right.next
		}
	}

	// append the remaining side
	if left != nil {
		current.next, left.previous = left, current
	} else {
		current.next, right.previous = right, current
	}

	return start
}
