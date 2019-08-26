package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


func rangeSumBST(root *TreeNode, L int, R int) int {
    if root == nil {
		return 0
	}

	sum := 0

	switch {
	case root.Val < L:
		sum = rangeSumBST(root.Right, L, R)
	case R < root.Val:
		sum = rangeSumBST(root.Left, L, R)
	default:
		sum = root.Val
		sum += rangeSumBST(root.Left, L, R)
		sum += rangeSumBST(root.Right, L, R)
	}

	return sum
}

func main() {
	root := &TreeNode{2,&TreeNode{1,nil,nil}, &TreeNode{3,nil,nil}}
	res := rangeSumBST(root, 1,2)
	fmt.Println(res)
}