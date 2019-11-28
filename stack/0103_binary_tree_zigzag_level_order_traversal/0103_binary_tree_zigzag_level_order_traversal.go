
package main 

import (
	"fmt"
)

// Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		if level >= len(res) {
			res = append(res, []int{})
		}
		if level %2 == 0 {
			res[level] = append(res[level], root.Val)
		} else {
			temp := make([]int, len(res[level]) + 1)
			temp[0] = root.Val
			copy(temp[1:], res[level])
			res[level] = temp
		}
		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}

	dfs(root, 0 )
	return res
}

func main() {
	// [3,9,20,null,null,15,7]
	root := &TreeNode{3, nil, nil}
	root.Left = &TreeNode{9, nil,nil}
	root.Right = &TreeNode{20, &TreeNode{15, nil,nil}, &TreeNode{7, nil,nil}}
	res := zigzagLevelOrder(root)
	fmt.Println(res)
}