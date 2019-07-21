package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

func main() {
	node := &ListNode{1, nil}
	node.Next = &ListNode{2, nil}
	node.Next.Next = &ListNode{3, nil}
	deleteNode(node.Next)
	fmt.Println(node.Next.Val)
}
