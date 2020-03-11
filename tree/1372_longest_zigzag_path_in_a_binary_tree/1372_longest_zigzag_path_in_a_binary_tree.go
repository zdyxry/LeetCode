package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestZigZag(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res, _, _ := bst(root)
	return res - 1
}

func bst(node *TreeNode) (int, int, int) {
	if node == nil {
		return 0, 0, 0
	}
	lb, _, lr := bst(node.Left)
	rb, rl, _ := bst(node.Right)
	l := lr + 1
	r := rl + 1
	b := max(max(l, r), max(lb, rb))
	return b, l, r
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
