package main 

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func buildTree(in []int, post []int) *TreeNode {
	if len(in) == 0 {
		return nil
	}

	res := &TreeNode{
		Val: post[len(post)-1],
	}

	if len(in) == 1 {
		return res
	}

	idx := indexOf(res.Val, in)
	res.Left = buildTree(in[:idx], post[:idx])
	res.Right = buildTree(in[idx+1:], post[idx:len(post)-1])

	return res
}

func indexOf(val int, nums []int) int {
	for i, v := range nums {
		if v == val {
			return i
		}
	}
	return 0
}

func main() {
	in := []int{9,3,15,20,7}
	post := []int{9,15,7,20,3}
	res := buildTree(in, post)
	fmt.Println(res.Val)
}