package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return sum == root.Val
	}

	return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
}

func main() {
	root := &TreeNode{5, nil, nil}
	root.Left = &TreeNode{4, nil, nil}
	root.Right = &TreeNode{8, nil, nil}
	root.Left.Left = &TreeNode{11, &TreeNode{7, nil, nil}, &TreeNode{2, nil, nil}}
	root.Right.Left = &TreeNode{13, nil, nil}
	root.Right.Right = &TreeNode{4, nil, &TreeNode{1, nil, nil}}
	sum := 22
	fmt.Println(hasPathSum(root, sum))

}
