package linkedlist

// LinkedList base linked list data structure
type LinkedList interface {
	// Append pushes an item to the end of the list
	Append(item int)
	// Prepend add an item to the start of the list
	Prepend(item int)
	// Contains returns true if item is found
	Contains(item int) bool
	// Delete removes the item requested from the list
	Delete(item int)
	// Traverse call fn for each item in the list
	Traverse(fn TraverseFn)
	// Clear empties the list
	Clear()
	// Len total of items in this list
	Len() int
}

type TraverseFn = func(int) bool
