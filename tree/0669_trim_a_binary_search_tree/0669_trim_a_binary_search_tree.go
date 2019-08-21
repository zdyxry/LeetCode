package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func trimBST(root *TreeNode, L, R int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val < L {
		return trimBST(root.Right, L, R)
	}

	if R < root.Val {
		return trimBST(root.Left, L, R)
	}

	root.Left = trimBST(root.Left, L, R)
	root.Right = trimBST(root.Right, L, R)
	return root
}

func main() {
	root := &TreeNode{1, &TreeNode{0, nil, nil}, &TreeNode{2, nil, nil}}
	res := trimBST(root, 1, 2)
	fmt.Println(res.Val)
}