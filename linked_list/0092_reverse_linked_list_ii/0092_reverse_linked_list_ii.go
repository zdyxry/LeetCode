package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || m >= n {
		return head
	}

	newHead := &ListNode{Val: 0, Next: head}
	pre := newHead
	for count := 0; pre.Next != nil && count < m-1; count++ {
		pre = pre.Next
	}
	if pre.Next == nil {
		return head
	}
	cur := pre.Next
	for i := 0; i < n-m; i++ {
		tmp := pre.Next
		pre.Next = cur.Next
		cur.Next = cur.Next.Next
		pre.Next.Next = tmp
	}

	return newHead.Next
}

// 1,2,3,4,5
// pre = 1, cur = 2
// loop 1:
// tmp = 2
// 1.next = 3
// 2.next = 2.next.next =4
// 1.next.next = 2

// 1,3,2,4,5
// pre=1, cur = 2
// loop2:
// tmp = 1.next = 3
// 1.next = 2.next = 4
// 2.next = 2.next.next =5
// 1.next.next = 3

// 1,4,3,2,5

