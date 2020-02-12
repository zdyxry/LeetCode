package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	s1 := make([]int, 0, 128)
	for l1 != nil {
		s1 = append(s1, l1.Val)
		l1 = l1.Next
	}

	s2 := make([]int, 0, 128)
	for l2 != nil {
		s2 = append(s2, l2.Val)
		l2 = l2.Next
	}

	sum := 0
	head := &ListNode{Val: 0}

	for len(s1) > 0 || len(s2) > 0 {
		if len(s1) > 0 {
			sum += s1[len(s1)-1]
			s1 = s1[:len(s1)-1]
		}

		if len(s2) > 0 {
			sum += s2[len(s2)-1]
			s2 = s2[:len(s2)-1]
		}

		head.Val = sum % 10
		ln := &ListNode{Val: sum / 10}
		ln.Next = head
		head = ln

		sum /= 10
	}

	if head.Val == 0 {
		return head.Next
	}

	return head

}

func main() {
	l1 := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	l2 := &ListNode{4, &ListNode{5, &ListNode{6, nil}}}

	res := addTwoNumbers(l1, l2)
	fmt.Println(res.Next.Val)
}
