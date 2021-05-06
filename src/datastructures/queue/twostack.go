package queue

import "github.com/felipeguilhermefs/asof/datastructures/stack"

// TwoStackQueue queue (FIFO) implemented with 2 stacks
// The idea is that:
//  - we enqueue to stack 1
//  - we dequeue from stack 2
//  - when there is nothing in stack 2
//    we pop from stack 1 and push to stack 2
type TwoStackQueue struct {
	stack1 stack.LinkedStack
	stack2 stack.LinkedStack
}

// Len total of items in queue
func (q *TwoStackQueue) Len() int {
	return q.stack1.Len() + q.stack2.Len()
}

// Enqueue insert item at the end of the queue
func (q *TwoStackQueue) Enqueue(item int) {
	// always push to the stack 1
	q.stack1.Push(item)
}

// Dequeue get item at the start of the queue
func (q *TwoStackQueue) Dequeue() (int, bool) {
	// we pop from stack 2 if there is something there
	if item, ok := q.stack2.Pop(); ok {
		return item, true
	}

	// otherwise we get everything from stack 1 and push to stack 2
	// given that they pop LIFO their order will be reversed and
	// we'll get the FIFO behaviour
	item, ok := q.stack1.Pop()
	for ok {
		q.stack2.Push(item)
		item, ok = q.stack1.Pop()
	}

	// now we can just pop from stack 2
	return q.stack2.Pop()
}
