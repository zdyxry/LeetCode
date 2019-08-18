package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil{
		return nil
	}

	if t1 == nil {
		return t2
	}

	if t2 == nil {
		return t1
	}

	root := &TreeNode{Val: t1.Val + t2.Val}

	root.Left = mergeTrees(t1.Left, t2.Left)
	root.Right = mergeTrees(t1.Right, t2.Right)

	return root
}

func main() {
	t1 := &TreeNode{1, nil, nil}
	t1.Left = &TreeNode{3, &TreeNode{5, nil,nil}, nil}
	t1.Right = &TreeNode{2,nil,nil}

	t2 := &TreeNode{2, nil, nil}
	t2.Left = &TreeNode{1, nil, &TreeNode{4,nil,nil}}
	t2.Right = &TreeNode{3, nil, &TreeNode{7,nil,nil}}

	res := mergeTrees(t1, t2)
	fmt.Println(res)
}