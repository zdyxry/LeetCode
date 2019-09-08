package main 

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	resPre := &ListNode{}
	cur := resPre
	carry := 0 

	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry = sum / 10

		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
	}

	return resPre.Next
}

func main() {
	l1 := &ListNode{1, &ListNode{2,&ListNode{3, nil}}}
	l2 := &ListNode{4, &ListNode{5,&ListNode{6, nil}}}

	res := addTwoNumbers(l1, l2)
	fmt.Println(res.Val)
}