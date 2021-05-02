package list

// SingleLinkedList single link data structure
// please prefer import "container/list"
type SingleLinkedList struct {
	head   *singleLinkNode
	tail   *singleLinkNode
	length int
}

type singleLinkNode struct {
	value int
	next  *singleLinkNode
}

// Append pushes an item to the end of the list
func (ll *SingleLinkedList) Append(item int) {
	newNode := &singleLinkNode{item, nil}

	if ll.head == nil {
		ll.head = newNode
	} else {
		ll.tail.next = newNode
	}

	ll.tail = newNode
	ll.length++
}

// Prepend add an item to the start of the list
func (ll *SingleLinkedList) Prepend(item int) {
	newNode := &singleLinkNode{item, ll.head}

	if ll.head == nil {
		ll.tail = newNode
	}

	ll.head = newNode
	ll.length++
}

// Contains returns true if item is found
func (ll *SingleLinkedList) Contains(item int) bool {
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
func (ll *SingleLinkedList) Delete(item int) {
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
func (ll *SingleLinkedList) Traverse(fn func(int)) {
	current := ll.head
	for current != nil {
		fn(current.value)
		current = current.next
	}
}

// Clear removes references hold by the list
func (ll *SingleLinkedList) Clear() {
	ll.head = nil
	ll.tail = nil
	ll.length = 0
}

// Len total of items in this list
func (ll *SingleLinkedList) Len() int {
	return ll.length
}
