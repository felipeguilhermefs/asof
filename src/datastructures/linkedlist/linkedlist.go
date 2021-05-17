package linkedlist

// LinkedList single link data structure
type LinkedList struct {
	start  *node
	end    *node
	length int
}

type node struct {
	value int
	next  *node
}

// PushRight pushes an element to the end/right of the list. O(1) time complexity
func (ll *LinkedList) PushRight(element int) {
	newNode := &node{element, nil}

	if ll.start == nil {
		ll.start = newNode
	} else {
		ll.end.next = newNode
	}

	ll.end = newNode
	ll.length++
}

// PopRight remove and return the element at the end of the list. O(n) time complexity
//     If not empty returns element, true
//     Else 					 -1, false
func (ll *LinkedList) PopRight() (int, bool) {
	if ll.end == nil {
		return -1, false
	}

	if ll.start == ll.end {
		element := ll.end.value

		ll.Clear()

		return element, true
	}

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

// PushLeft add an element to the start/left of the list. O(1) time complexity
func (ll *LinkedList) PushLeft(element int) {
	newNode := &node{element, ll.start}

	if ll.start == nil {
		ll.end = newNode
	}

	ll.start = newNode
	ll.length++
}

// PopLeft remove and return the element at the start of the list. O(1) time complexity
//     If not empty returns element, true
//     Else 					 -1, false
func (ll *LinkedList) PopLeft() (int, bool) {
	if ll.start == nil {
		return -1, false
	}

	if ll.start == ll.end {
		element := ll.end.value

		ll.Clear()

		return element, true
	}

	element := ll.start.value

	ll.start = ll.start.next
	ll.length--

	return element, true
}

// Contains returns true if element is found. O(n) time complexity
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

// Delete removes the element requested from the list. O(n) time complexity
func (ll *LinkedList) Delete(element int) {
	if ll.start == nil {
		return
	}

	if ll.start.value == element {
		ll.start = ll.start.next
		if ll.start == nil {
			ll.end = nil
		}

		ll.length--
		return
	}

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

// Clear empties the list. O(1) time complexity
func (ll *LinkedList) Clear() {
	ll.start = nil
	ll.end = nil
	ll.length = 0
}

// Len total of elements in this list. O(1) time complexity
func (ll *LinkedList) Len() int {
	return ll.length
}
