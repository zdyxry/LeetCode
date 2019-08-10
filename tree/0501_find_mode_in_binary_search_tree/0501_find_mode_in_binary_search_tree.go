package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) []int {
	if root == nil{
		return nil
	}

	prev := -1 << 31
	count := 0
	max := 1
	result := make([]int, 0)
	findBSTMode(root, &prev, &count, &max, &result)

	return result
}

func findBSTMode(root *TreeNode, prev, count,max *int, result *[]int) {
	if root == nil {
		return 
	}
	findBSTMode(root.Left, prev,count,max, result)
	if root.Val == *prev {
		*count++
	} else {
		*count=1
	}
	if *max < *count {
		*result = []int{root.Val}
		*max = *count
	} else if *max == *count{
		*result = append(*result, root.Val)
	}
	*prev = root.Val

	findBSTMode(root.Right, prev, count, max, result)
}

func main() {
	root := &TreeNode{1, nil,nil}
	root.Right = &TreeNode{2,nil,nil}
	root.Right.Right = &TreeNode{2,nil,nil}
	res := findMode(root)
	fmt.Println(res)
}