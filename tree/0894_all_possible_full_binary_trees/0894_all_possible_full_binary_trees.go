/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func allPossibleFBT(N int) []*TreeNode {
	if N%2 == 0 {
		return nil
	}
	if N == 1 {
		return []*TreeNode{new(TreeNode)}
	}

	trees := make([]*TreeNode, 0)
	for i := 1; i <= N-2; i = i + 2 { //注意循环的起点和循环条件
		l := allPossibleFBT(i)
		r := allPossibleFBT(N - 1 - i)
		for _, left := range l {
			for _, right := range r {
				node := new(TreeNode)
				node.Left = left
				node.Right = right
				trees = append(trees, node)
			}
		}
	}
	return trees
}
