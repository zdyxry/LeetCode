package main 

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func buildTree(pre []int, in []int) *TreeNode {
	if len(in) == 0 {
		return nil
	}

	res := &TreeNode{
		Val: pre[0],
	}

	if len(in) == 1 {
		return res
	}

	idx := indexOf(res.Val, in)

	res.Left = buildTree(pre[1:idx+1], in[:idx])
	res.Right = buildTree(pre[idx+1:], in[idx+1:])

	return res
}

func indexOf(val int, nums[]int) int {
	for i, v := range nums{
		if v == val{
			return i
		}
	}

	return 0
}

func main() {
	pre := []int{3,9,20,15,7}
	in := []int{9,3,15,20,7}

	res := buildTree(pre, in)
	fmt.Println(res)
}