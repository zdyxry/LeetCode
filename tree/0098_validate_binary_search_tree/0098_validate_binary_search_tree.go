package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	MIN, MAX := -1<<63, 1<<63-1

	return recur(MIN, MAX, root)
}

func recur(min, max int, root *TreeNode) bool {
	if root == nil {
		return true
	}

	return min < root.Val && root.Val < max && recur(min, root.Val, root.Left) && recur(root.Val, max, root.Right)
}

func main() {
	root := &TreeNode{}
	res := isValidBST(root)
	fmt.Println(res)
}
