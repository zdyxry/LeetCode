package main

import (
	"container/heap"
	"fmt"
)

// KthLargest define
type KthLargest struct {
	k    int
	heap intHeap
}

//Constructor define
func Constructor(k int, nums []int) KthLargest {
	h := intHeap(nums)
	heap.Init(&h)

	for len(h) > k {
		heap.Pop(&h)
	}

	return KthLargest{
		k:    k,
		heap: h,
	}
}

func (kl *KthLargest) Add(val int) int {
	heap.Push(&kl.heap, val)

	if len(kl.heap) > kl.k {
		heap.Pop(&kl.heap)
	}

	return kl.heap[0]
}

type intHeap []int

func (h intHeap) Len() int {
	return len(h)
}

func (h intHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h intHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *intHeap) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}

func main() {
	nums := []int{4, 5, 8, 2}
	kl := Constructor(3, nums)
	fmt.Println(kl.Add(3))
	fmt.Println(kl.Add(5))
}
