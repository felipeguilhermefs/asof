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

// PushRight pushes an item to the end/right of the list. O(1) time complexity
func (ll *LinkedList) PushRight(item int) {
	newNode := &node{item, nil}

	if ll.head == nil {
		ll.head = newNode
	} else {
		ll.tail.next = newNode
	}

	ll.tail = newNode
	ll.length++
}

// PushLeft add an item to the start/left of the list. O(1) time complexity
func (ll *LinkedList) PushLeft(item int) {
	newNode := &node{item, ll.head}

	if ll.head == nil {
		ll.tail = newNode
	}

	ll.head = newNode
	ll.length++
}

// Contains returns true if item is found. O(n) time complexity
func (ll *LinkedList) Contains(item int) bool {
	current := ll.head
	for current != nil {
		if current.value == item {
			return true
		}
		current = current.next
	}
	return false
}

// Delete removes the item requested from the list. O(n) time complexity
func (ll *LinkedList) Delete(item int) {
	if ll.head == nil {
		return
	}

	if ll.head.value == item {
		ll.head = ll.head.next
		if ll.head == nil {
			ll.tail = nil
		}

		ll.length--
		return
	}

	current := ll.head
	for current.next != nil {
		if current.next.value == item {
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

// Traverse call fn for each item in the list. O(n) time complexity
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

// Len total of items in this list. O(1) time complexity
func (ll *LinkedList) Len() int {
	return ll.length
}
