package queue

// Queue FIFO data structure
type Queue interface {
	// Enqueue adds an item to the end of the queue
	Enqueue(item int)
	// Dequeue gets the item at the start of the queue
	// if empty returns -1 and false
	// if not returns the first item and true
	Dequeue() (int, bool)
	// Len total of items in queue
	Len() int
}
