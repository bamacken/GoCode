package structure

// An unrolled linked list is a data structure that combines the properties of linked lists and arrays.
// In an unrolled linked list, each node contains multiple elements, allowing for more efficient
// memory usage and cache performance compared to traditional linked lists.
// The main operations on an unrolled linked list include insertion, deletion, and search.
// https://www.geeksforgeeks.org/unrolled-linked-list-set-1-introduction/?ref=lbp

/*
Unrolled Linked List Pseudo Code
An unrolled linked list is a complex data structure that is a variant of a linked list in which each node contains multiple elements and a reference to the next node.

Define Unrolled Linked List Node Structure:

Create a class named UnrolledListNode:
Include an elements array to store multiple values.
Include a next reference to the next UnrolledListNode.
Include an numElements counter to track the number of valid elements in elements.
Initialize the Unrolled Linked List:

Create a head node:
Set head to null initially.
Function to Add an Element:

Traverse to the last node:
If the last node's elements array is not full, add the new element there.
If it is full, create a new node, transfer some elements to balance the load, and add the new element to the appropriate node.
Function to Remove an Element:

Traverse nodes looking for the element.
Remove the element from the node's elements array.
Optionally, rebalance or merge nodes if too few elements remain.
*/
