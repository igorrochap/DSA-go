package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var newList *ListNode
	for head != nil {
		nextNode := head.Next
		head.Next = newList
		newList = head
		head = nextNode
	}
	return newList
}

func display(head *ListNode) {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

func main() {
	testcases := []ListNode{
		{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
		{1, &ListNode{2, nil}},
		{},
	}
	for _, testcase := range testcases {
		list := reverseList(&testcase)
		display(list)
	}
}
