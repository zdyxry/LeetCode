package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeZeroSumSublists(head *ListNode) *ListNode {
	type NodeSum struct {
		Node *ListNode
		Sum  int
	}
	acc, sum, dummy := 0, make(map[int]NodeSum), &ListNode{Next: head}
	sum[0] = NodeSum{dummy, 0}
	for curr := head; curr != nil; curr = curr.Next {
		acc += curr.Val
		if p, ok := sum[acc]; ok {
			display(p.Node)
			fmt.Println(sum, acc, curr.Val)
			for node, subSum := p.Node.Next, p.Sum; node != curr; node = node.Next {
				subSum += node.Val
				delete(sum, subSum)
			}
			p.Node.Next = curr.Next
			display(p.Node)
		} else {
			sum[acc] = NodeSum{curr, acc}
		}
	}
	return dummy.Next
}

func display(head *ListNode) {
	list := head
	for list != nil {
		fmt.Printf("%+v ->", list.Val)
		list = list.Next
	}
	fmt.Println()
}

func main() {
	head := &ListNode{Val: 1, Next: nil}
	head.Next = &ListNode{Val: 2, Next: nil}
	head.Next.Next = &ListNode{Val: -3, Next: nil}
	head.Next.Next.Next = &ListNode{Val: 3, Next: &ListNode{Val: 1, Next: nil}}
	res := removeZeroSumSublists(head)
	display(res)
}
