package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil {
		return head
	}

	tail := head
	for i := 0; i < k; i++ {
		if tail.Next == nil {
			return rotateRight(head, k%(i+1))
		}
		tail = tail.Next
	}
	newTail := head
	for tail.Next != nil {
		newTail, tail = newTail.Next, tail.Next
	}

	newHead := newTail.Next
	newTail.Next, tail.Next = nil, head

	return newHead
}

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	res := rotateRight(head, 2)
	fmt.Println(res.Val)
}
