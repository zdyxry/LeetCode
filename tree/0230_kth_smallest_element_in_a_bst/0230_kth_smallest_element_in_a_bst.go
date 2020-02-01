package main 

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	res, count := 0, 0 
	inorder(root, k, &count, &res)
	return res
}

func inorder(node *TreeNode, k int, count *int, ans *int) {
	if node != nil {
		inorder(node.Left, k, count, ans)
		*count++
		if *count == k {
			*ans = node.Val
			return
		}
		inorder(node.Right, k, count, ans)
	}
}