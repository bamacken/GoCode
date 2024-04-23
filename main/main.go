package main

import (
	linkedlist "GoCode/structure/linkedlist"
	"fmt"
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

}
