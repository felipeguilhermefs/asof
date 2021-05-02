package list

// LinkedList single link data structure
// please prefer import "container/list"
type LinkedList struct {
	head   *node
	tail   *node
	length int
}

type node struct {
	value int
	next  *node
}

// Append pushes an item to the end of the list
func (ll *LinkedList) Append(item int) {
	newNode := &node{item, nil}

	if ll.head == nil {
		ll.head = newNode
	} else {
		ll.tail.next = newNode
	}

	ll.tail = newNode
	ll.length++
}

// Prepend add an item to the start of the list
func (ll *LinkedList) Prepend(item int) {
	newNode := &node{item, ll.head}

	if ll.head == nil {
		ll.tail = newNode
	}

	ll.head = newNode
	ll.length++
}

// Contains returns true if item is found
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

// Delete removes the item requested from the list
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

// Traverse call fn for each item in the list
func (ll *LinkedList) Traverse(fn func(int)) {
	current := ll.head
	for current != nil {
		fn(current.value)
		current = current.next
	}
}

// Len total of items in this list
func (ll *LinkedList) Len() int {
	return ll.length
}
