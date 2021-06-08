package linkedlist

// LinkedList liner collection of elements more info: https://en.wikipedia.org/wiki/Linked_list
// This implementation uses a singly linked list approach
type LinkedList struct {
	start  *node
	end    *node
	length int
}

type node struct {
	value int
	next  *node
}

// PushRight pushes an element to the end/right of the list.
// O(1) time complexity
func (ll *LinkedList) PushRight(element int) {
	newNode := &node{element, nil}

	if ll.IsEmpty() {
		ll.start = newNode
	} else {
		// when not empty the current end node should point
		// to the new one before it is updated, so we don't get
		// a broken link
		ll.end.next = newNode
	}

	ll.end = newNode
	ll.length++
}

// PopRight remove and return the element at the end of the list.
//     If not empty returns element, true
//     Else 					 -1, false
// O(n) time complexity
func (ll *LinkedList) PopRight() (int, bool) {
	if ll.IsEmpty() {
		return -1, false
	}

	// single element
	if ll.start == ll.end {
		element := ll.end.value

		ll.Clear()

		return element, true
	}

	// To pop the right most element the second to last needs to be found
	// so it can become the last one. Since it is a singly linked list
	// we need to traverse the list to find it
	current := ll.start
	for current.next != nil {
		if current.next == ll.end {
			element := ll.end.value

			ll.end = current
			ll.end.next = nil
			ll.length--

			return element, true
		}
	}

	// should never get here
	return -1, false
}

// PushLeft add an element to the start/left of the list.
// O(1) time complexity
func (ll *LinkedList) PushLeft(element int) {
	newNode := &node{element, ll.start}

	if ll.IsEmpty() {
		ll.end = newNode
	}

	ll.start = newNode
	ll.length++
}

// PopLeft remove and return the element at the start of the list.
//     If not empty returns element, true
//     Else 					 -1, false
// O(1) time complexity
func (ll *LinkedList) PopLeft() (int, bool) {
	if ll.IsEmpty() {
		return -1, false
	}

	element := ll.start.value

	// single element
	if ll.start == ll.end {
		ll.Clear()
	} else {
		ll.start = ll.start.next
		ll.length--
	}

	return element, true
}

// Contains returns true if element is found.
// O(n) time complexity
func (ll *LinkedList) Contains(element int) bool {
	current := ll.start
	for current != nil {
		if current.value == element {
			return true
		}
		current = current.next
	}
	return false
}

// Delete removes the element requested from the list.
// O(n) time complexity
func (ll *LinkedList) Delete(element int) {
	if ll.IsEmpty() {
		return
	}

	if ll.start.value == element {
		ll.PopLeft()

		return
	}

	// We keep the current node and operate over next pointer so we
	// can adjust pointers when the element is found
	current := ll.start
	for current.next != nil {
		if current.next.value == element {
			if current.next == ll.end {
				ll.end = current
			}

			current.next = current.next.next
			ll.length--
			return
		}

		current = current.next
	}
}

// Clear empties the list.
// O(1) time complexity
func (ll *LinkedList) Clear() {
	ll.start = nil
	ll.end = nil
	ll.length = 0
}

// Len total of elements in this list.
// O(1) time complexity
func (ll *LinkedList) Len() int {
	return ll.length
}

// IsEmpty returns true if there is no elements.
// O(1) time complexity
func (ll *LinkedList) IsEmpty() bool {
	return ll.start == nil
}
