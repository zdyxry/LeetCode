package main 

import (
	"container/heap"
	"fmt"
)

func findKthLargest(nums []int, k int) int {
	temp := highHeap(nums)
	h := &temp
	heap.Init(h)

	if k == 1 {
		return (*h)[0]
	}

	for i := 1; i < k; i++ {
		heap.Remove(h, 0)
	}

	return (*h)[0]
}

type highHeap []int

func (h highHeap) Len() int {
	return len(h)
}

func (h highHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h highHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *highHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *highHeap) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[0 : len(*h)-1]
	return res
}

func main() {
	nums := []int{3,2,1,5,6,4}
	k := 2
	res := findKthLargest(nums, k)
	fmt.Println(res)
}