package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapNodes(head *ListNode, k int) *ListNode {
	slow, fast := head, head
	for i := 0; i < k-1; i++ {
		fast = fast.Next
	}
	left := fast
	for fast.Next != nil {
		slow, fast = slow.Next, fast.Next
	}
	slow.Val, left.Val = left.Val, slow.Val
	return head
}
