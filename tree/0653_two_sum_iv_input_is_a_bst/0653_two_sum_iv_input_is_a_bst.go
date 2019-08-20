package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	return helper(root, root, k)
}

func helper(root, searchRoot *TreeNode, k int) bool {
	if root == nil {
		return false
	}
	return (root.Val * 2 != k && findNode(searchRoot, k-root.Val)) || helper(root.Left, searchRoot, k) || helper(root.Right, searchRoot, k)

}

func findNode(root *TreeNode, target int) bool {
	if root == nil {
		return false
	}
	if root.Val == target {
		return true
	}

	if root.Val < target {
		return findNode(root.Right, target)
	}
	return findNode(root.Left, target)
}

func main() {
	root := &TreeNode{2, nil, nil}
	root.Left = &TreeNode{1,nil,nil}
	root.Right = &TreeNode{3,nil,nil}
	res := findTarget(root, 4)
	fmt.Println(res)
}