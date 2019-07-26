package main

import (
	"container/heap"
	"fmt"
)

func lastStoneWeight(stones []int) int {
	pq := PQ(stones)
	heap.Init(&pq)
	for len(pq) >= 2 {
		a, b := heap.Pop(&pq).(int), heap.Pop(&pq).(int)
		if a == b {
			continue
		}
		heap.Push(&pq, a-b)
	}

	if len(pq) == 0 {
		return 0
	}

	return pq[0]
}

type PQ []int

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	return pq[i] > pq[j]
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PQ) Push(x interface{}) {
	temp := x.(int)
	*pq = append(*pq, temp)
}

func (pq *PQ) Pop() interface{} {
	temp := (*pq)[len(*pq)-1]
	*pq = (*pq)[0 : len(*pq)-1]
	return temp

}

func main() {
	stones := []int{2, 7, 4, 1, 8, 1}
	res := lastStoneWeight(stones)
	fmt.Println(res)
}
