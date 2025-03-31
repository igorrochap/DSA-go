package main

import "fmt"

type Node struct {
	Value    int
	Next     *Node
	Previous *Node
}

type DoublyLinkedList struct {
	Head *Node
	Tail *Node
}

func (list *DoublyLinkedList) AddToFront(value int) {
	newNode := &Node{Value: value, Next: nil, Previous: nil}
	newNode.Next = list.Head
	if list.Head != nil {
		list.Head.Previous = newNode
	} else {
		list.Tail = newNode
	}
	list.Head = newNode
}

func (list *DoublyLinkedList) AddToEnd(value int) {
	newNode := &Node{Value: value, Next: nil, Previous: nil}
	newNode.Previous = list.Tail
	if list.Tail != nil {
		list.Tail.Next = newNode
	} else {
		list.Head = newNode
	}
	list.Tail = newNode
}

func (list *DoublyLinkedList) RemoveFromFront() (int, bool) {
	if list.Head == nil {
		return -1, false
	}
	removedValue := list.Head.Value
	list.Head = list.Head.Next
	if list.Head != nil {
		list.Head.Previous = nil
	} else {
		list.Tail = nil
	}
	return removedValue, true
}

func (list *DoublyLinkedList) RemoveFromEnd() (int, bool) {
	if list.Tail == nil {
		return -1, false
	}
	removedValue := list.Tail.Value
	list.Tail = list.Tail.Previous
	if list.Tail != nil {
		list.Tail.Previous = nil
	} else {
		list.Head = nil
	}
	return removedValue, true
}

func main() {
	dll := &DoublyLinkedList{}
	dll.AddToFront(3)
	dll.AddToFront(2)
	dll.AddToFront(1)
	dll.AddToEnd(4)
	dll.AddToEnd(5)

	fmt.Println(dll.RemoveFromFront())
	fmt.Println(dll.RemoveFromEnd())
	fmt.Println(dll.RemoveFromFront())
	fmt.Println(dll.RemoveFromEnd())
}
