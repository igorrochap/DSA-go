package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func mergeSort(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}
	middle := findMiddle(head)
	afterMiddle := middle.Next
	middle.Next = nil
	left := mergeSort(head)
	right := mergeSort(afterMiddle)
	sortedList := merge(left, right)
	return sortedList
}

func findMiddle(head *Node) *Node {
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func merge(list1, list2 *Node) *Node {
	head := &Node{}
	tail := head
	for list1 != nil && list2 != nil {
		if list1.Value < list2.Value {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}
		tail = tail.Next
	}
	if list1 != nil {
		tail.Next = list1
	}
	if list2 != nil {
		tail.Next = list2
	}
	return head.Next
}

func main() {
	list := &Node{Value: 7, Next: &Node{Value: 1, Next: &Node{Value: 3, Next: &Node{Value: 9}}}}
	sorted := mergeSort(list)
	for sorted != nil {
		fmt.Println(sorted.Value)
		sorted = sorted.Next
	}
}
