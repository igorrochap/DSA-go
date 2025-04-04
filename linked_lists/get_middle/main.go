package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	fast := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		head = head.Next
	}
	return head
}

func main() {
	testcases := []ListNode{
		{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
		{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}},
	}
	for _, testcase := range testcases {
		middle := middleNode(&testcase)
		fmt.Println(middle.Val)
	}
}
