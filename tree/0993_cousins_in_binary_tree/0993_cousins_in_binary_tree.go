package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCousins(root *TreeNode, x, y int) bool {
	root = &TreeNode{Left: root}
	px, dx := dfs(root, x)
	py, dy := dfs(root, y)
	return dx == dy && px != py
}

func dfs(root *TreeNode, x int) (*TreeNode, int) {
	if root == nil {
		return nil, 0
	}

	if (root.Left != nil && root.Left.Val == x) || (root.Right != nil && root.Right.Val == x) {
		return root, 1
	}

	if parent, depth := dfs(root.Left, x); depth > 0 {
		return parent, depth + 1
	}

	if parent, depth := dfs(root.Right, x); depth > 0 {
		return parent, depth + 1
	}

	return nil, 0
}

func main() {
	root := &TreeNode{1,nil,nil}
	root.Left = &TreeNode{2,nil,nil}
	root.Right = &TreeNode{3,nil,nil}
	root.Left.Left = &TreeNode{4,nil,nil}
	res := isCousins(root, 3, 4)
	fmt.Println(res)
}