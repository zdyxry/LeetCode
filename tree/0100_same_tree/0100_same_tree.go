package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func main() {
	p := &TreeNode{1, nil, nil}
	p.Left = &TreeNode{2, nil, nil}
	p.Right = &TreeNode{3, nil, nil}

	q := &TreeNode{1, nil, nil}
	q.Left = &TreeNode{2, nil, nil}
	q.Right = &TreeNode{3, nil, nil}

	res := isSameTree(p, q)
	fmt.Println(res)
}
