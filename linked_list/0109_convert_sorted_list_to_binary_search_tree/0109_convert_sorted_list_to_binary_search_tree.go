package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	return transMidToRoot(head, nil)
}

func transMidToRoot(begin, end *ListNode) *TreeNode {
	if begin == end {
		return nil
	}

	if begin.Next == end {
		return &TreeNode{Val: begin.Val}
	}

	fast, slow := begin, begin
	for fast != end && fast.Next != end {
		fast = fast.Next.Next
		slow = slow.Next
	}

	mid := slow

	return &TreeNode{
		Val:   mid.Val,
		Left:  transMidToRoot(begin, mid),
		Right: transMidToRoot(mid.Next, end),
	}
}
