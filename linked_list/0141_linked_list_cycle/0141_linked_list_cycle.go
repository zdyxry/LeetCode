package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head != nil {
		fast, slow := head, head

		for fast.Next != nil && fast.Next.Next != nil {
			fast = fast.Next.Next
			slow = slow.Next
			if slow == fast {
				return true
			}
		}
	}

	return false
}

func main() {
	head := &ListNode{1, nil}
	head.Next = &ListNode{2, nil}
	head.Next.Next = &ListNode{3, nil}
	// head.Next.Next.Next = head
	fmt.Println(hasCycle(head))
}
