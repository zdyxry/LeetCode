package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {
	var sum int = 0
	var maxLevel int = 0
	deep(root, 0, &maxLevel, &sum)
	return sum
}

func deep(root *TreeNode, level int, maxLevel *int, sum *int) {
	if root == nil {
		return
	}
	if level > *maxLevel {
		*sum = root.Val
		*maxLevel = level
	} else if level == *maxLevel {
		*sum = *sum + root.Val
	}
	deep(root.Left, level+1, maxLevel, sum)
	deep(root.Right, level+1, maxLevel, sum)
}
