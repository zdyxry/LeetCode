package main

import "sort"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	res := make([]int, 0, 10)
	getList(root1, &res)
	getList(root2, &res)
	sort.Ints(res)
	return res
}
func getList(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	getList(root.Left, res)
	getList(root.Right, res)
}
