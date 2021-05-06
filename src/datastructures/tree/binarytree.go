package tree

// Binary Tree https://en.wikipedia.org/wiki/Binary_search_tree
type BinaryTree struct {
	root *node
}

type node struct {
	value int
	left  *node
	right *node
}

func (bt *BinaryTree) Insert(value int) {
	newNode := &node{value, nil, nil}
	if bt.root == nil {
		bt.root = newNode
	} else {
		bt.insert(bt.root, newNode)
	}
}

func (bt *BinaryTree) insert(n *node, newNode *node) {
	if n.value < newNode.value {
		if n.right == nil {
			n.right = newNode
		} else {
			bt.insert(n.right, newNode)
		}
	} else {
		if n.left == nil {
			n.left = newNode
		} else {
			bt.insert(n.left, newNode)
		}
	}
}

func (bt *BinaryTree) Contains(value int) bool {
	return bt.contains(bt.root, value)
}

func (bt *BinaryTree) contains(n *node, value int) bool {
	if n == nil {
		return false
	}

	if n.value == value {
		return true
	}

	if n.value < value {
		return bt.contains(n.right, value)
	} else {
		return bt.contains(n.left, value)
	}
}

const (
	InOrder   TraverseOrder = 0
	PreOrder  TraverseOrder = 1
	PostOrder TraverseOrder = 2
)

type TraverseOrder = int

type TraverseConsumer = func(int)

func (bt *BinaryTree) Traverse(order TraverseOrder, consumer TraverseConsumer) {
	if order == InOrder {
		bt.inOrderTraverse(bt.root, consumer)
		return
	}

	if order == PreOrder {
		bt.preOrderTraverse(bt.root, consumer)
		return
	}

	bt.postOrderTraverse(bt.root, consumer)
}

func (bt *BinaryTree) inOrderTraverse(n *node, consumer TraverseConsumer) {
	if n == nil {
		return
	}

	bt.inOrderTraverse(n.left, consumer)
	consumer(n.value)
	bt.inOrderTraverse(n.right, consumer)
}

func (bt *BinaryTree) preOrderTraverse(n *node, consumer TraverseConsumer) {
	if n == nil {
		return
	}

	consumer(n.value)
	bt.preOrderTraverse(n.left, consumer)
	bt.preOrderTraverse(n.right, consumer)
}

func (bt *BinaryTree) postOrderTraverse(n *node, consumer TraverseConsumer) {
	if n == nil {
		return
	}

	bt.postOrderTraverse(n.left, consumer)
	bt.postOrderTraverse(n.right, consumer)
	consumer(n.value)
}
