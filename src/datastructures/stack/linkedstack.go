package stack

// LinkedStack stack using a linked list
type LinkedStack struct {
	length int
	top    *node
}

type node struct {
	value int
	prev  *node
}

func (s *LinkedStack) Len() int {
	return s.length
}

func (s *LinkedStack) Push(item int) {
	s.top = &node{
		value: item,
		prev:  s.top,
	}

	s.length++
}

func (s *LinkedStack) Pop() (int, bool) {
	if s.length == 0 {
		return -1, false
	}

	topNone := s.top

	s.top = topNone.prev
	s.length--

	return topNone.value, true
}
