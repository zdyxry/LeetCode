package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	head2 := head.Next
	p1 := head
	p2 := head2
	for p1.Next != nil && p2.Next != nil {
		p1.Next = p2.Next
		p1 = p1.Next
		p2.Next = p1.Next
		p2 = p2.Next
	}
	p1.Next = head2
	return head
}
