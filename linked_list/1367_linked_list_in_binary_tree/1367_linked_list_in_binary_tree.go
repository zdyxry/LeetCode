package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubPath(head *ListNode, root *TreeNode) bool {
	if head == nil {
		return true
	}
	if root == nil {
		return false
	}
	return downwards(head, root) || isSubPath(head, root.Right) || isSubPath(head, root.Left)
}

func downwards(head *ListNode, node *TreeNode) bool {
	if head == nil {
		return true
	}
	if node == nil {
		return false
	}
	return head.Val == node.Val && (downwards(head.Next, node.Left) || downwards(head.Next, node.Right))
}
