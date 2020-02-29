package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}

	max := nums[0]
	posMax := 0
	for i := 1; i < len(nums); i++ {
		if max < nums[i] {
			max = nums[i]
			posMax = i
		}
	}

	leaf := TreeNode{Val: nums[posMax]}
	if posMax >= 1 {
		leaf.Left = constructMaximumBinaryTree(nums[:posMax])
	}
	if posMax+1 < len(nums) {
		leaf.Right = constructMaximumBinaryTree(nums[posMax+1:])
	}
	return &leaf
}
