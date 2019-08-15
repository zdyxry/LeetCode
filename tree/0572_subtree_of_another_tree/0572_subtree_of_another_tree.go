package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if isSame(s, t) {
		return true
	}
	if s == nil {
		return false
	}
	return isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func isSame(s *TreeNode, t *TreeNode) bool {
	if s == nil {return t == nil}
	if t == nil {return false}
	if s.Val != t.Val {
		return false
	}

	return isSame(s.Left, t.Left) && isSame(s.Right, t.Right)
}

func main() {
	s := &TreeNode{3, nil,nil}
	s.Left = &TreeNode{4, 
		&TreeNode{1,nil,nil},
		&TreeNode{2, nil,nil},
	}
	s.Right = &TreeNode{5, nil,nil}

	t := &TreeNode{4, 
		&TreeNode{1,nil,nil},
		&TreeNode{2, nil,nil},
	}
	res := isSubtree(s, t)
	fmt.Println(res)
}