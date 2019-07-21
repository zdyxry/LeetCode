package main

import "fmt"

// ListNode is linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func main() {
	head := &ListNode{1, nil}
	head.Next = &ListNode{2, nil}
	head.Next.Next = &ListNode{3, nil}
	head.Next.Next.Next = &ListNode{4, nil}
	result := middleNode(head)
	fmt.Println(result.Val)
}
