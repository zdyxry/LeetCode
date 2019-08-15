package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	_, res := helper(root)
	return res
}

func helper(root *TreeNode) (length, diameter int) {
	if root == nil {
		return 
	}

	leftLen, leftDia := helper(root.Left)
	rightLen, rightDia := helper(root.Right)

	length = max(leftLen, rightLen) + 1
	diameter = max(leftLen+rightLen, max(leftDia, rightDia))
	return length, diameter
}

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	root := &TreeNode{1, nil,nil}
	root.Left = &TreeNode{2, nil,nil}
	root.Right = &TreeNode{3,nil,nil}
	root.Left.Left = &TreeNode{4, nil,nil}
	root.Left.Right = &TreeNode{5,nil,nil}
	res := diameterOfBinaryTree(root)
	fmt.Println(res)
}