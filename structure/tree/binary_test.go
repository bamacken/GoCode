package tree

import (
	"fmt"
	"testing"
)

func TestBinaryTree(t *testing.T) {
	// Test case 1: Insert elements into the binary tree
	tree := NewBinaryTree(5)
	tree.Insert(3)
	tree.Insert(7)
	tree.Insert(2)
	tree.Insert(4)
	tree.Insert(6)
	tree.Insert(8)

	// Test case 2: Check the size of the binary tree
	if tree.Size() != 7 {
		t.Errorf("Expected size of 7, got %d", tree.Size())
	}

	// Test case 2: Check if an element exists in the binary tree
	if !tree.Contains(4) {
		t.Errorf("Expected element 4 to exist in the binary tree")
	}

	// Test case 4: Check if an element does not exist in the binary tree
	if tree.Contains(9) {
		t.Errorf("Expected element 9 to not exist in the binary tree")
	}

	// Test case 5: Traverse the binary tree in-order
	expectedInOrder := []int{2, 3, 4, 5, 6, 7, 8}
	inOrder := tree.InOrderTraversal()
	if !equalSlices(inOrder, expectedInOrder) {
		t.Errorf("Expected in-order traversal %v, got %v", expectedInOrder, inOrder)
	}

	// Test case 6: Traverse the binary tree pre-order
	expectedPreOrder := []int{5, 3, 2, 4, 7, 6, 8}
	preOrder := tree.PreOrderTraversal()
	if !equalSlices(preOrder, expectedPreOrder) {
		t.Errorf("Expected pre-order traversal %v, got %v", expectedPreOrder, preOrder)
	}

	// Test case 7: Traverse the binary tree post-order
	expectedPostOrder := []int{2, 4, 3, 6, 8, 7, 5}
	postOrder := tree.PostOrderTraversal()
	if !equalSlices(postOrder, expectedPostOrder) {
		t.Errorf("Expected post-order traversal %v, got %v", expectedPostOrder, postOrder)
	}

	// Test case 8: Remove an element from the binary tree
	tree.Remove(4)
	if tree.Contains(4) {
		t.Errorf("Expected element 4 to be removed from the binary tree")
	}
}

// Helper function to check if two slices are equal
func equalSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// Example function for NewBinaryTree
func ExampleNewBinaryTree() {
	tree := NewBinaryTree(5)
	fmt.Println(tree.Size())
	// Output: 1
}

// Example function for Insert
func ExampleBinaryTree_Insert() {
	tree := NewBinaryTree(5)
	tree.Insert(3)
	tree.Insert(7)
	fmt.Println(tree.Size())
	// Output: 3
}

// Example function for Size
func ExampleBinaryTree_Size() {
	tree := NewBinaryTree(5)
	tree.Insert(3)
	tree.Insert(7)
	fmt.Println(tree.Size())
	// Output: 3
}

// Example function for Contains
func ExampleBinaryTree_Contains() {
	tree := NewBinaryTree(5)
	tree.Insert(3)
	tree.Insert(7)
	fmt.Println(tree.Contains(3))
	// Output: true
}

// Example function for InOrderTraversal
func ExampleBinaryTree_InOrderTraversal() {
	tree := NewBinaryTree(5)
	tree.Insert(3)
	tree.Insert(7)
	inOrder := tree.InOrderTraversal()
	fmt.Println(inOrder)
	// Output: [3 5 7]
}

// Example function for PreOrderTraversal
func ExampleBinaryTree_PreOrderTraversal() {
	tree := NewBinaryTree(5)
	tree.Insert(3)
	tree.Insert(7)
	preOrder := tree.PreOrderTraversal()
	fmt.Println(preOrder)
	// Output: [5 3 7]
}

// Example function for PostOrderTraversal
func ExampleBinaryTree_PostOrderTraversal() {
	tree := NewBinaryTree(5)
	tree.Insert(3)
	tree.Insert(7)
	postOrder := tree.PostOrderTraversal()
	fmt.Println(postOrder)
	// Output: [3 7 5]
}

// Example function for Remove
func ExampleBinaryTree_Remove() {
	tree := NewBinaryTree(5)
	tree.Insert(3)
	tree.Insert(7)
	tree.Remove(3)
	fmt.Println(tree.Contains(3))
	// Output: false
}

// Benchmark function for NewBinaryTree
func BenchmarkNewBinaryTree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewBinaryTree(5)
	}
}

// Benchmark function for Insert
func BenchmarkBinaryTree_Insert(b *testing.B) {
	tree := NewBinaryTree(5)
	for i := 0; i < b.N; i++ {
		tree.Insert(i)
	}
}

// Benchmark function for Size
func BenchmarkBinaryTree_Size(b *testing.B) {
	tree := NewBinaryTree(5)
	for i := 0; i < b.N; i++ {
		tree.Size()
	}
}

// Benchmark function for Contains
func BenchmarkBinaryTree_Contains(b *testing.B) {
	tree := NewBinaryTree(5)
	for i := 0; i < b.N; i++ {
		tree.Contains(3)
	}
}

// Benchmark function for InOrderTraversal
func BenchmarkBinaryTree_InOrderTraversal(b *testing.B) {
	tree := NewBinaryTree(5)
	for i := 0; i < b.N; i++ {
		tree.InOrderTraversal()
	}
}

// Benchmark function for PreOrderTraversal
func BenchmarkBinaryTree_PreOrderTraversal(b *testing.B) {
	tree := NewBinaryTree(5)
	for i := 0; i < b.N; i++ {
		tree.PreOrderTraversal()
	}
}

// Benchmark function for PostOrderTraversal
func BenchmarkBinaryTree_PostOrderTraversal(b *testing.B) {
	tree := NewBinaryTree(5)
	for i := 0; i < b.N; i++ {
		tree.PostOrderTraversal()
	}
}

// Benchmark function for Remove
func BenchmarkBinaryTree_Remove(b *testing.B) {
	tree := NewBinaryTree(5)
	for i := 0; i < b.N; i++ {
		tree.Remove(3)
	}
}
