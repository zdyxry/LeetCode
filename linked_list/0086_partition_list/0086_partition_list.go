package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	lessHead := &ListNode{}
	noLessHead := &ListNode{}

	lessEnd := lessHead
	noLessEnd := noLessHead

	for head != nil {
		if head.Val < x {
			lessEnd.Next = head
			lessEnd = lessEnd.Next
		} else {
			noLessEnd.Next = head
			noLessEnd = noLessEnd.Next
		}
		head = head.Next
	}

	lessEnd.Next = noLessEnd.Next
	noLessEnd.Next = nil

	head = lessHead.Next
	return head
}
