package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func dfs(n *TreeNode, d int) (map[int]int, int) {
	if n == nil {
		return map[int]int{}, 0
	}
	if n.Left == nil && n.Right == nil {
		return map[int]int{0: 1}, 0
	}

	lp, lc := dfs(n.Left, d)
	rp, rc := dfs(n.Right, d)

	m := map[int]int{}
	c := lc + rc
	for kr, vr := range rp {
		m[kr+1] += vr
	}

	for kl, vl := range lp {
		m[kl+1] += vl
		for kr, vr := range rp {
			if kl+kr+2 <= d {
				c += vl * vr
			}
		}
	}

	//fmt.Println(n.Val, m, c)

	return m, c
}

func countPairs(root *TreeNode, distance int) int {
	_, lc := dfs(root, distance)
	return lc
}
