package main 

import (
	"fmt"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	var rightStack []*TreeNode
	var res []int
	for cur := root; cur != nil; {
		res = append(res, cur.Val)

		if cur.Left != nil {
			if cur.Right != nil {
				rightStack = append(rightStack, cur.Right)
			}
			cur = cur.Left
		} else {
			if cur.Right != nil {
				cur = cur.Right
			} else {
				if len(rightStack) == 0 {
					break
				}
				cur = rightStack[len(rightStack)-1]
				rightStack = rightStack[:len(rightStack)-1]
			}
		}
	}
	return res
}

func main() {
	root := &TreeNode{1, nil, &TreeNode{2, &TreeNode{3, nil,nil}, nil}} 
	res := preorderTraversal(root)
	fmt.Println(res)
}