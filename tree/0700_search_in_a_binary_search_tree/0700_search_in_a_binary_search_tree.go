package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	switch{
	case root.Val < val:
		return searchBST(root.Right, val)
	case root.Val > val:
		return searchBST(root.Left, val)
	default:
		return root
	}
}

func main() {
	root := &TreeNode{2,nil,nil}
	root.Left = &TreeNode{1,nil,nil}
	root.Right = &TreeNode{3, nil,nil}
	res := searchBST(root, 3)
	fmt.Println(res)
}