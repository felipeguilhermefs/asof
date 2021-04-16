package stack

type Stack struct {
	items []int
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() (int, bool) {
	if len(s.items) == 0 {
		return -1, false
	}

	last := len(s.items) -1

	item := s.items[last]
	s.items = s.items[:last]

	return item, true
}
