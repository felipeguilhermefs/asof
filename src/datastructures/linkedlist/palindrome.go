package linkedlist

// palindrome check if the LinkedList is a palindrome.
// In a doubly linked list it is enough to traverse in both
// directions comparing left and right elements.
// O(n) time complexity.
func palindrome(list *LinkedList) bool {
	if list == nil || list.IsEmpty() {
		return false
	}

	half := list.Len() / 2
	left := list.start
	right := list.end
	// As optimization we traverse only half of the list in the worst case.
	for i := 0; i <= half; i++ {
		if left.value != right.value {
			return false
		}

		left = left.next
		right = right.previous
	}

	return true
}
