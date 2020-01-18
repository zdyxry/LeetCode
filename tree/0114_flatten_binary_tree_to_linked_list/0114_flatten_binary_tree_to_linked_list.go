package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	recur(root)
}

func recur(root *TreeNode) *TreeNode {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return root
	}

	if root.Left == nil {
		return recur(root.Right)
	}

	if root.Right == nil {
		root.Right = root.Left
		root.Left = nil
		return recur(root.Right)
	}

	res := recur(root.Right)
	recur(root.Left).Right = root.Right
	root.Right = root.Left
	root.Left = nil

	return res
}
