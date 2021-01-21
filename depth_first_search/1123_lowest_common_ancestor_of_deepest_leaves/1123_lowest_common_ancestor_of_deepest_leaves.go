package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	//if root == nil{return nil}
	_, lca := lcaHelper(root, 0)
	return lca
}

func lcaHelper(root *TreeNode, depth int) (int, *TreeNode) {
	if root == nil {
		return depth, nil
	}
	depth++
	ldepth, left := lcaHelper(root.Left, depth)
	rdepth, right := lcaHelper(root.Right, depth)
	if ldepth > rdepth {
		return ldepth, left
	}
	if rdepth > ldepth {
		return rdepth, right
	}
	// Since both sides are the same depth, this is the lca
	return ldepth, root
}
