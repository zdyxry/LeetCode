package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func increasingBST(root *TreeNode) *TreeNode {
	min, _ := helper(root)
	return min
}

func helper(node *TreeNode) (*TreeNode, *TreeNode) {
	min, max := node, node
	if node.Left != nil {
		lMin, lMax := helper(node.Left)
		min = lMin
		lMax.Right = node
	}
	if node.Right != nil {
		rMin, rMax := helper(node.Right)
		max = rMax
		node.Right = rMin
	}
	node.Left = nil
	return min, max
}
