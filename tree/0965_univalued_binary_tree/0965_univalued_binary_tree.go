package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if (root.Left != nil && root.Val != root.Left.Val) || (root.Right != nil && root.Val != root.Right.Val) {
		return false
	}

	return isUnivalTree(root.Left) && isUnivalTree(root.Right)
}

func main() {
	root := &TreeNode{1, nil,nil}
	root.Left = &TreeNode{1, &TreeNode{1,nil,nil}, &TreeNode{1,nil,nil}}
	root.Right = &TreeNode{1, nil, &TreeNode{1,nil,nil}}

	res := isUnivalTree(root)
	fmt.Println(res)
}