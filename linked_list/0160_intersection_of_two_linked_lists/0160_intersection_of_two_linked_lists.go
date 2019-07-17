package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	a := headA
	b := headB

	//if a & b have different len, then we will stop the loop after second iteration
	for a != b {
		if a != nil {
			fmt.Println("a is", a.Val)
		}
		if b != nil {
			fmt.Println("b is", b.Val)
		}
		// fmt.Println(a.Val, b.Val)
		//for the end of first iteration, we just reset the pointer to the head of another linkedlist
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}

		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
		fmt.Println(a, b)
	}
	return a
}

func main() {
	headA := &ListNode{1, nil}
	headA.Next = &ListNode{2, nil}
	headA.Next.Next = &ListNode{3, nil}

	headB := &ListNode{2, nil}
	headB.Next = &ListNode{3, nil}
	result := getIntersectionNode(headA, headB)
	fmt.Println(result)
}
