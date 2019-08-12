package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type state struct {
	minDiff  int
	previous int
}

func getMinimumDifference(root *TreeNode) int {
	st := state{1024, 1024}
	search(root, &st)
	return st.minDiff
}

func search(root *TreeNode, st *state) {
	if root == nil {
		return
	}

	search(root.Left, st)

	newDiff := diff(st.previous, root.Val)
	if st.minDiff > newDiff {
		st.minDiff = newDiff
	}

	st.previous = root.Val
	search(root.Right, st)
}

func diff(a, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}

func main() {
	root := &TreeNode{1, nil, nil}
	root.Right = &TreeNode{3, nil, nil}
	root.Right.Left = &TreeNode{2, nil, nil}
	res := getMinimumDifference(root)
	fmt.Println(res)
}
