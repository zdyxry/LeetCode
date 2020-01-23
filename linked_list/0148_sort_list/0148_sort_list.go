package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	left, right := split(head)
	return merge(sortList(left), sortList(right))
}

func split(head *ListNode) (left, right *ListNode) {
	slow, fast := head, head
	var slowPre *ListNode

	for fast != nil && fast.Next != nil {
		slowPre, slow = slow, slow.Next
		fast = fast.Next.Next
	}

	slowPre.Next = nil

	left, right = head, slow
	return
}

func merge(left, right *ListNode) *ListNode {
	cur := &ListNode{}
	headPre := cur

	for left != nil && right != nil {
		if left.Val < right.Val {
			cur.Next, left = left, left.Next
		} else {
			cur.Next, right = right, right.Next
		}
		cur = cur.Next
	}

	if left == nil {
		cur.Next = right
	} else {
		cur.Next = left
	}

	return headPre.Next
}

func main() {
	head := &ListNode{3, nil}
	head.Next = &ListNode{2, nil}
	head.Next.Next = &ListNode{1, nil}
	res := sortList(head)
	fmt.Println(res)
}
