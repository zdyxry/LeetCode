package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func goodNodes(root *TreeNode) (goodNodeCount int) {
	getGoodNodes(root, root.Val, &goodNodeCount)
	return
}

func getGoodNodes(root *TreeNode, prev int, count *int) {
	if root == nil {
		return
	}

	if prev <= root.Val {
		*count += 1
		prev = root.Val
	}

	getGoodNodes(root.Left, prev, count)
	getGoodNodes(root.Right, prev, count)
}
