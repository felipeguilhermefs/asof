package queue

// LinkedQueue queue (FIFO) implemented using node pointers
type LinkedQueue struct {
	length int
	head   *node
	tail   *node
}

type node struct {
	value int
	next  *node
	prev  *node
}

func (q *LinkedQueue) Len() int {
	return q.length
}

func (q *LinkedQueue) Enqueue(item int) {
	q.length++
	if q.tail == nil {
		newNode := &node{item, nil, nil}
		q.tail = newNode
		q.head = newNode
	} else {
		newNode := &node{item, q.tail, nil}
		q.tail.prev = newNode
		q.tail = newNode
	}
}

func (q *LinkedQueue) Dequeue() (int, bool) {
	if q.head == nil {
		return -1, false
	}

	lastHead := q.head

	q.head = q.head.prev
	if q.head == nil {
		q.tail = nil
	}

	q.length--

	return lastHead.value, true
}
