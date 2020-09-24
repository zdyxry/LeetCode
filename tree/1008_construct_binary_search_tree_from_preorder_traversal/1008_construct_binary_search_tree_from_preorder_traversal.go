package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstFromPreorder(preorder []int) *TreeNode {
	if devide := len(preorder); devide > 0 {
		for i, val := range preorder {
			if val > preorder[0] {
				devide = i
				break
			}
		}
		return &TreeNode{
			Val:   preorder[0],
			Left:  bstFromPreorder(preorder[1:devide]),
			Right: bstFromPreorder(preorder[devide:]),
		}
	}
	return nil
}
