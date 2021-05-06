package linkedlist

// SingleLinkedList single link data structure
type SingleLinkedList struct {
	head   *singleLinkNode
	tail   *singleLinkNode
	length int
}

type singleLinkNode struct {
	value int
	next  *singleLinkNode
}

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

func (ll *SingleLinkedList) Prepend(item int) {
	newNode := &singleLinkNode{item, ll.head}

	if ll.head == nil {
		ll.tail = newNode
	}

	ll.head = newNode
	ll.length++
}

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

func (ll *SingleLinkedList) Traverse(fn func(int) bool) {
	current := ll.head

	for current != nil {

		if !fn(current.value) {
			return
		}

		current = current.next
	}
}

func (ll *SingleLinkedList) Clear() {
	ll.head = nil
	ll.tail = nil
	ll.length = 0
}

func (ll *SingleLinkedList) Len() int {
	return ll.length
}
