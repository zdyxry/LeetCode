package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return helper(root, p.Val, q.Val)
}

func helper(root *TreeNode, p, q int) *TreeNode {
	r := root.Val
	if p < r && q < r {
		return helper(root.Left, p, q)
	} else if r < p && r < q {
		return helper(root.Right, p, q)
	}
	return root
}

func main() {
	p := &TreeNode{2, nil, nil}
	q := &TreeNode{8, nil, nil}
	root := &TreeNode{6,
		p,
		q,
	}
	res := lowestCommonAncestor(root, p, q)
	fmt.Println(res.Val)
}
