package stack

// SliceStack stack using slices
type SliceStack struct {
	items []int
}

func (s *SliceStack) Len() int {
	return len(s.items)
}

func (s *SliceStack) Push(item int) {
	s.items = append(s.items, item)
}

func (s *SliceStack) Pop() (int, bool) {
	if len(s.items) == 0 {
		return -1, false
	}

	last := len(s.items) - 1

	item := s.items[last]
	s.items = s.items[:last]

	return item, true
}

func (s *SliceStack) Peek() (int, bool) {
	if len(s.items) == 0 {
		return -1, false
	}

	return s.items[len(s.items)-1], true
}
