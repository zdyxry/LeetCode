package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDiffInBST(root *TreeNode) int {
	res := 1<<63-1
	pre := 1>>63
	null := pre

	var helper func(*TreeNode)
	helper = func(root *TreeNode) {
		if root.Left != nil {
			helper(root.Left)
		}

		if pre != null {
			res = min(res, root.Val - pre)
		}

		pre = root.Val

		if root.Right != nil {
			helper(root.Right)
		}
	}

	helper(root)

	return res
}

func min(a,b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	root := &TreeNode{4, nil,&TreeNode{6, nil,nil}}
	root.Left = &TreeNode{2, &TreeNode{1,nil,nil}, &TreeNode{3,nil,nil}}
	res := minDiffInBST(root)
	fmt.Println(res)
}