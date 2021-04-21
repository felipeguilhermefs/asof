package list

// LinkedList single link data structure
type LinkedList struct {
	head *node
	length  int
}

type node struct {
	value int
	next  *node
}

// Append pushes a item to the end of the list
func (ll *LinkedList) Append(item int) {
	ll.length++

	newNode := &node{item, nil}

	if ll.head == nil {
		ll.head = newNode
		return
	}

	current := ll.head
	for current.next != nil {
		current = current.next
	}

	current.next = newNode
}

// Prepend pushes a item to the begining of the list
func (ll *LinkedList) Prepend(item int) {
	newHead := &node{item, ll.head}
	ll.head = newHead
	ll.length++
}

// Delete removes the item requested from the list
func (ll *LinkedList) Delete(item int) {
	if ll.head == nil {
		return
	}

	if ll.head.value == item {
		ll.head = ll.head.next
		ll.length--
		return
	}

	current := ll.head
	for current.next != nil {
		if current.next.value == item {
			current.next = current.next.next
			ll.length--
			return
		}

		current = current.next
	}
}

// Index get the item at index
func (ll *LinkedList) Index(i int) (int, bool) {

	current := ll.head
	for j := 0; j <= i; j++ {
		if current == nil {
			return 0, false
		}

		if j == i {
			return current.value, true
		}

		current = current.next
	}

	return 0, false
}

// Len total of items in this list
func (ll *LinkedList) Len() int {
	return ll.length
}
