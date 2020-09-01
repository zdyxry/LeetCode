package main

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxLevelSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var (
		curMaxSum      = math.MinInt32
		curMaxSumLevel = 1
		curLevel       = 1
		q              []*TreeNode
	)
	q = append(q, root)
	for len(q) != 0 {
		size := len(q)
		curLevelSum := 0
		for i := 0; i < size; i++ {
			item := q[0]
			curLevelSum += item.Val
			if item.Left != nil {
				q = append(q, item.Left)
			}
			if item.Right != nil {
				q = append(q, item.Right)
			}
			q = q[1:]
		}
		if curLevelSum > curMaxSum {
			curMaxSum = curLevelSum
			curMaxSumLevel = curLevel
		}
		curLevel++
	}
	return curMaxSumLevel
}
