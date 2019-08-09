package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) int {
	m := make(map[int]int)
	m[0] = 1
	ans := 0
	helper(root, sum, m, 0, &ans)
	return ans
}

func helper(root *TreeNode, target int, m map[int]int, cur int, ans *int) {
	fmt.Println("##### m is", m)
	if root == nil {
		return
	}
	cur = cur + root.Val
	*ans += m[cur-target]
	m[cur]++
	if root.Left != nil {
		helper(root.Left, target, m, cur, ans)
	}
	if root.Right != nil {
		helper(root.Right, target, m, cur, ans)
	}

	m[cur]--
}

func main() {
	root := &TreeNode{10, nil, nil}
	root.Left = &TreeNode{5, nil, nil}
	root.Right = &TreeNode{-3, nil, nil}
	root.Left.Left = &TreeNode{3, nil, nil}
	root.Left.Right = &TreeNode{2, nil, nil}
	root.Right.Right = &TreeNode{11, nil, nil}
	res := pathSum(root, 8)
	fmt.Println(res)
}
