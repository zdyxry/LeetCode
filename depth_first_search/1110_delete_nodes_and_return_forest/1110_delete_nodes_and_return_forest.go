package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	m := make(map[int]bool)
	for i := 0; i < len(to_delete); i++ {
		m[to_delete[i]] = true
	}
	return delNode(root, m)
}

func delNode(root *TreeNode, m map[int]bool) []*TreeNode {
	root, ret := del(root, m)
	if root == nil {
		return ret
	}
	return append(ret, root)
}

func del(root *TreeNode, m map[int]bool) (*TreeNode, []*TreeNode) {
	ret := []*TreeNode{}
	if root == nil {
		return nil, nil
	}
	if m[root.Val] {
		return nil, append(delNode(root.Left, m), delNode(root.Right, m)...)
	}
	left, l := del(root.Left, m)
	root.Left = left
	ret = append(ret, l...)

	right, r := del(root.Right, m)
	root.Right = right
	ret = append(ret, r...)

	return root, ret
}
