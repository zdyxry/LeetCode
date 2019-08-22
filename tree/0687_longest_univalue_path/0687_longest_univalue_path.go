package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestUnivaluePath(root *TreeNode) int {
	maxLen := 0
	helper(root, &maxLen)
	return maxLen
}

func helper(n *TreeNode, maxLen *int) int {
	if n == nil {
		return 0
	}
	
	l := helper(n.Left, maxLen)
	r := helper(n.Right, maxLen)
	res := 0

	if n.Left != nil && n.Val == n.Left.Val {
		*maxLen = max(*maxLen, l + 1)
		res = max(res, l+1)
	}

	if n.Right != nil && n.Val == n.Right.Val {
		*maxLen = max(*maxLen, r + 1)
		res = max(res, r + 1)
	}

	if n.Left != nil && n.Val == n.Left.Val && n.Right != nil && n.Val == n.Right.Val {
		*maxLen = max(*maxLen, l + r + 2)
	}

	return res
}

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	root := &TreeNode{1, nil,nil}
	root.Left = &TreeNode{1, nil,nil}
	root.Right = &TreeNode{1, nil,nil}
	res := longestUnivaluePath(root)
	fmt.Println(res)
}