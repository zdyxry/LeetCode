package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func balanceBST(root *TreeNode) *TreeNode {
	sortedArray := getSortedArray(root)
	return buildBalanceBySortedArray(sortedArray)
}

func buildBalanceBySortedArray(array []int) *TreeNode {
	if len(array) == 0 {
		return nil
	}
	mid := len(array) / 2
	root := &TreeNode{Val: array[mid]}
	root.Left = buildBalanceBySortedArray(array[:mid])
	root.Right = buildBalanceBySortedArray(array[mid+1:])
	return root
}

func getSortedArray(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	res = append(res, getSortedArray(root.Left)...)
	res = append(res, root.Val)
	res = append(res, getSortedArray(root.Right)...)
	return res
}
