package linkedlist

type TraverseFn = func(int) bool

// LinkedList single link data structure
type LinkedList struct {
	head   *node
	tail   *node
	length int
}

type node struct {
	value int
	next  *node
}

// PushRight pushes an element to the end/right of the list. O(1) time complexity
func (ll *LinkedList) PushRight(element int) {
	newNode := &node{element, nil}

	if ll.head == nil {
		ll.head = newNode
	} else {
		ll.tail.next = newNode
	}

	ll.tail = newNode
	ll.length++
}

// PopRight remove and return the element at the end of the list. O(n) time complexity
//     If not empty returns element, true
//     Else 					 -1, false
func (ll *LinkedList) PopRight() (int, bool) {
	if ll.tail == nil {
		return -1, false
	}

	if ll.head == ll.tail {
		element := ll.tail.value

		ll.Clear()

		return element, true
	}

	current := ll.head
	for current.next != nil {
		if current.next == ll.tail {
			element := ll.tail.value

			ll.tail = current
			ll.tail.next = nil
			ll.length--

			return element, true
		}
	}

	// should never get here
	return -1, false
}

// PushLeft add an element to the start/left of the list. O(1) time complexity
func (ll *LinkedList) PushLeft(element int) {
	newNode := &node{element, ll.head}

	if ll.head == nil {
		ll.tail = newNode
	}

	ll.head = newNode
	ll.length++
}

// Contains returns true if element is found. O(n) time complexity
func (ll *LinkedList) Contains(element int) bool {
	current := ll.head
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
	if ll.head == nil {
		return
	}

	if ll.head.value == element {
		ll.head = ll.head.next
		if ll.head == nil {
			ll.tail = nil
		}

		ll.length--
		return
	}

	current := ll.head
	for current.next != nil {
		if current.next.value == element {
			if current.next == ll.tail {
				ll.tail = current
			}

			current.next = current.next.next
			ll.length--
			return
		}

		current = current.next
	}
}

// Traverse call fn for each element in the list. O(n) time complexity
func (ll *LinkedList) Traverse(fn func(int) bool) {
	current := ll.head

	for current != nil {

		if !fn(current.value) {
			return
		}

		current = current.next
	}
}

// Clear empties the list. O(1) time complexity
func (ll *LinkedList) Clear() {
	ll.head = nil
	ll.tail = nil
	ll.length = 0
}

// Len total of elements in this list. O(1) time complexity
func (ll *LinkedList) Len() int {
	return ll.length
}
