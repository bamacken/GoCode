package tree

import (
	"fmt"
)

type treeNode struct {
	value int
	left  *treeNode
	right *treeNode
}

// Create a new binary tree
func NewBinaryTree(value int) *treeNode {
	return &treeNode{value: value}
}

// Insert a node into the tree
func (b *treeNode) Insert(value int) {
	newNode := &treeNode{value: value}

	if b == nil {
		b = newNode
	} else {
		b.insertHelper(newNode)
	}
}

// helper function to insert a node into the tree
func (t *treeNode) insertHelper(node *treeNode) {
	if node.value < t.value {
		if t.left == nil {
			t.left = node
		} else {
			t.left.insertHelper(node)
		}
	} else {
		if t.right == nil {
			t.right = node
		} else {
			t.right.insertHelper(node)
		}
	}
}

func (t *treeNode) Contains(value int) bool {
	if t == nil {
		return false
	}

	if t.value == value {
		return true
	}

	if t.value < value {
		return t.right.Contains(value)
	} else {
		return t.left.Contains(value)
	}
}

func (t *treeNode) Height() (Height int) {
	if t == nil {
		return 0
	}

	left := t.left.Height()
	right := t.right.Height()

	return 1 + Max(left, right)
}

// Breadth First Search
func (t *treeNode) BredthFirstSearch() {
	if t == nil {
		fmt.Println("The tree is empty")
		return
	}

	// Slice based queue
	queue := []*treeNode{t}

	for len(queue) > 0 {
		current := queue[0]
		fmt.Println(current.value)
		queue = queue[1:]

		if current.left != nil {
			queue = append(queue, current.left)
		}
		if current.right != nil {
			queue = append(queue, current.right)
		}
	}
}

// Remove a node from the tree
func (t *treeNode) Remove(value int) {
	if t == nil {
		return
	}
	t.removalHelper(value)
}

// Helper function to remove a node from the tree
func (t *treeNode) removalHelper(value int) *treeNode {

	if value < t.value {
		//Search the left tree
		t.left = t.left.removalHelper(value)
	} else if value > t.value {
		// Search the right tree
		t.right = t.right.removalHelper(value)
	} else {

		// Case 1: Node has zero children
		if t.left == nil && t.right == nil {
			return nil
		}

		// Case 2: Node has a single child
		if t.left == nil {
			return t.right
		}

		if t.right == nil {
			return t.left
		}

		// Case 3: Node has two children
		temp := t.right.minValueInSubTree()
		t.value = temp.value
		t.right = t.right.removalHelper(temp.value)
	}
	return t
}

// Helper function to find the minimum value in the subtree
func (t *treeNode) minValueInSubTree() *treeNode {
	current := t

	for current.left != nil {
		current = current.left
	}
	return current
}

// Depth First Search - PreOrder
func (t *treeNode) PreOrderTraversal() []int {
	result := []int{}
	t.preOrderTraversalHelper(&result)
	return result
}

func (t *treeNode) preOrderTraversalHelper(result *[]int) {
	if t == nil {
		return
	}

	// Start with the current node value
	*result = append(*result, t.value)

	// Recursively append left subtree results
	if t.left != nil {
		t.left.preOrderTraversalHelper(result)
	}

	// Recursively append right subtree results
	if t.right != nil {
		t.right.preOrderTraversalHelper(result)
	}
}

// Depth First Search - InOrder
func (t *treeNode) InOrderTraversal() []int {
	result := []int{}
	t.inOrderTraversalHelper(&result)
	return result
}

func (t *treeNode) inOrderTraversalHelper(result *[]int) {
	if t == nil {
		return
	}

	// Recursively append left subtree results
	if t.left != nil {
		t.left.inOrderTraversalHelper(result)
	}

	// Start with the current node value
	*result = append(*result, t.value)

	// Recursively append right subtree results
	if t.right != nil {
		t.right.inOrderTraversalHelper(result)
	}
}

// Depth First Search - PostOrder
func (t *treeNode) PostOrderTraversal() []int {
	result := []int{}
	t.postOrderTraversalHelper(&result)
	return result
}

func (t *treeNode) postOrderTraversalHelper(result *[]int) {
	if t == nil {
		return
	}

	// Recursively append left subtree results
	if t.left != nil {
		t.left.postOrderTraversalHelper(result)
	}

	// Recursively append right subtree results
	if t.right != nil {
		t.right.postOrderTraversalHelper(result)
	}

	// Start with the current node value
	*result = append(*result, t.value)
}

// Find the minimum value in the tree
func (t *treeNode) Min() (Min int) {
	if t == nil {
		fmt.Println("Tree is empty")
		return
	}
	currentNode := t
	for currentNode.left != nil {
		currentNode = currentNode.left
	}

	return currentNode.value
}

// Find the maximum value in the tree
func (t *treeNode) Max() (Max int) {
	if t == nil {
		fmt.Println("Tree is empty")
		return
	}
	currentNode := t
	for currentNode.right != nil {
		currentNode = currentNode.right
	}

	return currentNode.value
}

// Max returns the larger of x or y.
func Max(x, y int) (Max int) {
	if x > y {
		return x
	} else {
		return y
	}
}

func (t *treeNode) Size() int {
	if t == nil {
		return 0
	}
	return 1 + t.left.Size() + t.right.Size()
}
