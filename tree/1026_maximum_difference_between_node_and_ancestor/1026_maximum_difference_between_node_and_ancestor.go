package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var maxDiff int

func maxAncestorDiff(root *TreeNode) int {
	maxDiff = 0
	dfs(root, root.Val, root.Val)
	return maxDiff
}

func dfs(root *TreeNode, min, max int) {
	if nil == root {
		return
	}
	val := root.Val
	if val < min {
		min = val
	}
	if val > max {
		max = val
	}
	if max-min > maxDiff {
		maxDiff = max - min
	}
	if nil != root.Left {
		dfs(root.Left, min, max)
	}
	if nil != root.Right {
		dfs(root.Right, min, max)
	}
	return
}
