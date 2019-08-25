package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	a1 := [100]int{}
	s1 := a1[:0]
	search(root1, &s1)

	a2 := [100]int{}
	s2 := a2[:0]
	search(root2, &s2)

	return a1 == a2
}

func search(root *TreeNode, sp *[]int) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		*sp = append(*sp, root.Val)
		return
	}

	search(root.Left, sp)
	search(root.Right, sp)
}

func main() {
	root1 := &TreeNode{1, &TreeNode{2,nil,nil}, &TreeNode{3,nil,nil}}
	root2 := &TreeNode{4, &TreeNode{2,nil,nil}, &TreeNode{3,nil,nil}}
	res := leafSimilar(root1, root2)
	fmt.Println(res)
}