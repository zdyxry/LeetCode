package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isEvenOddTree(root *TreeNode) bool {

	queue := []*TreeNode{root}
	level := 0
	for len(queue) > 0 {
		sign := 1
		if level&1 == 1 {
			sign = -1
		}

		length := len(queue)
		for i := 0; i < length; i++ {
			cur := queue[i]

			if (cur.Val&1) == (level&1) || i > 0 && queue[i-1].Val*sign >= queue[i].Val*sign {
				return false
			}

			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}

		level++
		queue = queue[length:]
	}

	return true
}
