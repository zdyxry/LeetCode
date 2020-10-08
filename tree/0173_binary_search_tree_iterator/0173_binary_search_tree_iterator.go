package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	nums []int
	root *TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	nums := make([]int, 0)
	inorder(root, &nums)
	return BSTIterator{root: root, nums: nums}
}

func inorder(root *TreeNode, nums *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left, nums)
	*nums = append(*nums, root.Val)
	inorder(root.Right, nums)
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	res := this.nums[0]
	this.nums = this.nums[1:]
	return res
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	if len(this.nums) > 0 {
		return true
	}
	return false
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
