package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	var head, node *ListNode
	if l1.Val < l2.Val {
		head = l1
		node = l1
		l1 = l1.Next
	} else {
		head = l2
		node = l2
		l2 = l2.Next
	}

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			node.Next = l1
			l1 = l1.Next
		} else {
			node.Next = l2
			l2 = l2.Next
		}

		// 有了这一步，head才是一个完整的链
		node = node.Next
	}

	if l1 != nil {
		// 连上l1剩余的链
		node.Next = l1
	}

	if l2 != nil {
		// 连上l2剩余的链
		node.Next = l2
	}

	return head
}

func main() {
	l1 := &ListNode{1, nil}
	l1.Next = &ListNode{2, nil}
	l1.Next.Next = &ListNode{4, nil}
	l2 := &ListNode{1, nil}
	l2.Next = &ListNode{3, nil}
	l2.Next.Next = &ListNode{4, nil}
	head := mergeTwoLists(l1, l2)
	fmt.Println(head.Next.Next.Val)
}
