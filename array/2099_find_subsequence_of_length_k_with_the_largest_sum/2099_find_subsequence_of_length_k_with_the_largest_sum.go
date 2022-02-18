package main

type node struct {
	id  int
	val int
}

func maxSubsequence(nums []int, k int) []int {
	var nodes []node
	for i := range nums {
		nodes = append(nodes, node{
			id:  i,
			val: nums[i],
		})
	}
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].val > nodes[j].val
	})
	var r []int
	for i := 0; i < k; i++ {
		r = append(r, nodes[i].id)
	}
	sort.Ints(r)
	for i := range r {
		r[i] = nums[r[i]]
	}
	return r
}
