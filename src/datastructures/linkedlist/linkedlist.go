package linkedlist

// LinkedList liner collection of elements more info: https://en.wikipedia.org/wiki/Linked_list
// This implementation uses a doubly linked list approach
type LinkedList struct {
	start  *node
	end    *node
	length int
}

type node struct {
	value    int
	previous *node
	next     *node
}

// PushRight pushes an element to the end/right of the list.
// O(1) time complexity
func (ll *LinkedList) PushRight(element int) {
	newNode := &node{
		value:    element,
		previous: ll.end,
		next:     nil,
	}

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
// O(1) time complexity
func (ll *LinkedList) PopRight() (int, bool) {
	if ll.IsEmpty() {
		return -1, false
	}

	element := ll.end.value

	// single element
	if ll.start == ll.end {
		ll.Clear()
	} else {
		ll.end = ll.end.previous
		ll.end.next = nil
		ll.length--
	}

	return element, true
}

// PushLeft add an element to the start/left of the list.
// O(1) time complexity
func (ll *LinkedList) PushLeft(element int) {
	newNode := &node{
		value:    element,
		previous: nil,
		next:     ll.start,
	}

	if ll.IsEmpty() {
		ll.end = newNode
	} else {
		// when not empty the current start node should point
		// to the new one before it is updated, so we don't get
		// a broken link
		ll.start.previous = newNode
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
		ll.start.previous = nil
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

	if ll.end.value == element {
		ll.PopRight()

		return
	}

	current := ll.start.next
	for current != nil {
		if current.value == element {
			current.previous.next = current.next

			if current.next != nil {
				current.next.previous = current.previous
			}

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
