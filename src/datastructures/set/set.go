package set

//Set collection of unique items
type Set struct {
	items map[int]bool
}

// Add a item to the set
func (s *Set) Add(value int) {
	if s.items == nil {
		s.items = make(map[int]bool)
	}

	s.items[value] = true
}

// Contains check if item is in the set
func (s *Set) Contains(value int) bool {
	if s.items == nil {
		return false
	}

	_, ok := s.items[value]

	return ok
}

// Len return set cardinality
func (s *Set) Len() int {
	return len(s.items)
}
