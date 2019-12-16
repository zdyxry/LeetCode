package main 

import (
	"container/heap"
	"fmt"
)

type minHeap [][]int

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Less(i,j int) bool {
	return h[i][0] + h[i][1] < h[j][0] + h[j][1]
}

func (h minHeap) Swap(i,j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))

}

func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0:n-1]
	return x
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int{
	result , h := [][]int{}, &minHeap{}

	if len(nums1) == 0 || len(nums2) == 0 || k == 0 {
		return result
	}

	if len(nums1) * len(nums2) < k {
		k = len(nums1) * len(nums2)
	}
	heap.Init(h)
	for _, num := range nums1 {
		heap.Push(h, []int{num, nums2[0], 0})
	}
	for len(result) < k {
		min := heap.Pop(h).([]int)
		result = append(result, min[:2])
		if min[2] < len(nums2) - 1 { 
			heap.Push(h, []int{min[0], nums2[min[2]+1], min[2]+1})
		}
	}
	return result
}

func main() {
	nums1 := []int{1,7,11}
	nums2 := []int{2,4,6}
	k := 3
	res := kSmallestPairs(nums1, nums2, k)
	fmt.Println(res)
}