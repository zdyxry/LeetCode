package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertBst(root *TreeNode) *TreeNode {
	sum := 0
	travel(root, &sum)
	return root
}

func travel(root *TreeNode, sum *int) {
	if root == nil {
		return
	}
	travel(root.Right, sum)
	*sum += root.Val
	root.Val = *sum
	travel(root.Left, sum)
}

func main() {
	root := &TreeNode{5, nil, nil}
	root.Left = &TreeNode{2, nil, nil}
	root.Right = &TreeNode{13, nil, nil}
	res := convertBst(root)
	fmt.Println(res.Left.Val)
}
