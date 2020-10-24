package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func btreeGameWinningMove(root *TreeNode, n int, x int) bool {

	var (
		target     *TreeNode
		count      func(root *TreeNode) int
		findTarget func(root *TreeNode)
	)

	count = func(root *TreeNode) int {
		if root == nil || root == target {
			return 0
		}
		return 1 + count(root.Left) + count(root.Right)
	}

	findTarget = func(root *TreeNode) {
		if root != nil {
			if root.Val == x {
				target = root
				return
			}
			findTarget(root.Left)
			findTarget(root.Right)
		}
	}

	findTarget(root)

	return count(root) > n/2 || count(target.Left) > n/2 || count(target.Right) > n/2
}
