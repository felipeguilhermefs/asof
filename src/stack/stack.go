package stack

// Stack is a simple LIFO data structure implementation
type Stack struct {
	items []int
}

func (s *Stack) Len() int {
	return len(s.items)
}

// Push adds an item to top of the stack
func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

// Pop gets the item at the top of the stack
// if empty returns -1 and false
// if not returns the top item and true
func (s *Stack) Pop() (int, bool) {
	if len(s.items) == 0 {
		return -1, false
	}

	last := len(s.items) -1

	item := s.items[last]
	s.items = s.items[:last]

	return item, true
}
