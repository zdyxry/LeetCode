package main 

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func findTilt(root *TreeNode) int {
	if root == nil{return 0}
	if root.Left == nil && root.Right == nil{return 0}
	ans := 0
	value:=sum(root.Left) - sum(root.Right)
	if value < 0{value = -value}
	ans += value
	return ans + findTilt(root.Left) + findTilt(root.Right)
}
func sum(node *TreeNode)int{
	if node == nil {return 0}
	return node.Val + sum(node.Left) + sum(node.Right)
}
