package stack

// Stack LIFO data structure
type Stack interface {
	// Push adds an item to top of the stack
	Push(item int)
	// Pop gets the item at the top of the stack
	// if empty returns -1 and false
	// if not returns the top item and true
	Pop() (int, bool)
	// Peek similar to Pop, but doesn't remove the item from the stack
	Peek() (int, bool)
	// Len total of items in stack
	Len() int
}

// NewStack creates a new stack instance
func NewStack() Stack {
	//LinkedStack is a better implementation at the moment
	return &LinkedStack{}
}
