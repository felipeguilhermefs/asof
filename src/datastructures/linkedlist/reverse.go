package linkedlist

//reverse the linked list inplace
func reverse(list *LinkedList) {
	if list == nil {
		return
	}

	list.end = list.start

	var prev *node
	current := list.start
	for current != nil {
		next := current.next
		current.next = prev
		prev = current
		current = next
	}

	list.start = prev
}
