package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumRootToLeaf(root *TreeNode) int {
	ans := 0
	var f func(*TreeNode, int)
	f = func(r *TreeNode, s int) {
		if r != nil {
			s = s*2 + r.Val
			if r.Left == nil && r.Right == nil {
				ans += s
			} else {
				f(r.Left, s)
				f(r.Right, s)
			}
		}
	}
	f(root, 0)
	return ans
}
