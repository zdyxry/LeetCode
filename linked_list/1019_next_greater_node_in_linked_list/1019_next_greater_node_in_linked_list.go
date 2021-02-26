package main 

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func nextLargerNodes(head *ListNode) []int {
	var data []int
	for head != nil {
		data = append(data, head.Val)
		head = head.Next
	}
	stack, ans := make([]int, len(data)), make([]int, len(data))

	for index := 0; index < len(data); index++ {
		if len(stack) == 0 {
			stack = append(stack, index)
		} else {
			for len(stack) > 0 && data[index] > data[stack[len(stack)-1]] {
				pop := len(stack) - 1
				ans[stack[pop]] = data[index]
				stack = stack[:pop]
			}
			stack = append(stack, index)
		}
	}
	return ans
}
