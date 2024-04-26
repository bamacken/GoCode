package main

import (
	"GoCode/structure/array"
	linkedlist "GoCode/structure/linkedlist"
	"GoCode/structure/tree"
	"fmt"
	"sort"
)

type Person struct {
	firstName string
	lastName  string
}

func NamePerson(first, last string) {
	p := Person{first, last}
	fmt.Println(p.firstName + " " + p.lastName)
}

func main() {
	fmt.Println("Hello, Linkedlist!")

	ll := linkedlist.NewLinkedList()

	ll.Append(1)
	ll.Append(2)
	ll.Append(3)
	ll.Append(4)
	ll.Append(5)
	ll.Append(6)
	ll.Append(7)
	ll.Append(8)
	ll.Append(9)

	array1 := []float64{10.0, 20.0, 30.0, 40.0, 50.0}
	array2 := []float64{5.0, 15.0, 25.0, 35.0, 45.0}

	merged := append(array1, array2...)
	sort.Float64s(merged)

	fmt.Println("Merged and sorted array:", merged)

	arr1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	print(len((arr1)))

	d := array.NewDynamicArray(0)
	fmt.Print(d)

	tree := tree.NewBinaryTree(3)
	tree.Insert(1)
	tree.Insert(4)
	tree.Insert(2)
	tree.Insert(5)

	slice := tree.PostOrderTraversal()
	fmt.Println(slice)

	tree.Max()

	fmt.Printf("This binary tree is %d nodes tall", tree.Height())

	tree.BredthFirstSearch()

	tree.Remove(3)
}
