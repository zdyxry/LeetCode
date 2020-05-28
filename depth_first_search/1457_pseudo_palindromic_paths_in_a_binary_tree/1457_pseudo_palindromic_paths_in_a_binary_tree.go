package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pseudoPalindromicPaths(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var ans int
	var h func(node *TreeNode, temp int)
	h = func(node *TreeNode, temp int) {
		temp ^= (1 << node.Val)
		if node.Left == nil && node.Right == nil {
			if temp == 0 || temp&(temp-1) == 0 {
				ans++
			}
		}

		if node.Left != nil {
			h(node.Left, temp)
		}

		if node.Right != nil {
			h(node.Right, temp)
		}

	}

	h(root, 0)
	return ans
}
