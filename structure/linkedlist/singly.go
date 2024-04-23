package structure

type Node struct {
	Next  *Node
	Value int
}

type LinkedList struct {
	Head *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

// and the specified values as a new node at the end of the linkedlist
func (ll *LinkedList) Append(value int) {
	newNode := &Node{Value: value}
	if ll.Head == nil {
		ll.Head = newNode
		return
	}

	lastNode := ll.Head
	for lastNode.Next != nil {
		lastNode = lastNode.Next
	}
	lastNode = newNode

}

/*
Singly Linked List Pseudo Code
A singly linked list consists of nodes, where each node contains a value and a reference to the next node in the list.

Class Node
    Define value
    Define next as null

Class LinkedList
    Define head as null

    Function insert(value):
        newNode = new Node(value)
        newNode.next = head
        head = newNode

    Function search(value):
        current = head
        while current is not null:
            if current.value == value:
                Return "Found"
            current = current.next
        Return "Not found"

    Function delete(value):
        current = head
        previous = null
        while current is not null:
            if current.value == value:
                if previous is not null:
                    previous.next = current.next
                else:
                    head = current.next
                Return "Deleted"
            previous = current
            current = current.next
        Return "Not found"

*/
