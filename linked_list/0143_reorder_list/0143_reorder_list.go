package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil {
		return
	}

	runner := head
	size := 0
	stack := []*ListNode{}
	for runner != nil {
		stack = append(stack, runner)
		runner = runner.Next
		size++
	}
	curr := head
	for i := 0; i < size/2; i++ {
		stEl := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		next := curr.Next
		curr.Next = stEl
		stEl.Next = next
		curr = next
	}
	curr.Next = nil
}

func main() {
	head := &ListNode{1, nil}
	head.Next = &ListNode{2, nil}
	head.Next.Next = &ListNode{3, nil}
	reorderList(head)
	fmt.Println(head.Next.Val)
}
