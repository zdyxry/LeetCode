package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func checkTree(root *TreeNode) bool {
	if root.Left.Val+root.Right.Val == root.Val {
		return true
	}
	return false

}
