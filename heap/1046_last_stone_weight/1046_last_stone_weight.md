1046. Last Stone Weight


Easy


We have a collection of rocks, each rock has a positive integer weight.

Each turn, we choose the two heaviest rocks and smash them together.  Suppose the stones have weights x and y with x <= y.  The result of this smash is:

If x == y, both stones are totally destroyed;
If x != y, the stone of weight x is totally destroyed, and the stone of weight y has new weight y-x.
At the end, there is at most 1 stone left.  Return the weight of this stone (or 0 if there are no stones left.)

 

Example 1:

```
Input: [2,7,4,1,8,1]
Output: 1
Explanation: 
We combine 7 and 8 to get 1 so the array converts to [2,4,1,1,1] then,
we combine 2 and 4 to get 2 so the array converts to [2,1,1,1] then,
we combine 2 and 1 to get 1 so the array converts to [1,1,1] then,
we combine 1 and 1 to get 0 so the array converts to [1] then that's the value of last stone.
```

Note:

1 <= stones.length <= 30

1 <= stones[i] <= 1000


## 方法








```python
1046. Last Stone Weight
Easy

123

8

Favorite

Share
We have a collection of rocks, each rock has a positive integer weight.

Each turn, we choose the two heaviest rocks and smash them together.  Suppose the stones have weights x and y with x <= y.  The result of this smash is:

If x == y, both stones are totally destroyed;
If x != y, the stone of weight x is totally destroyed, and the stone of weight y has new weight y-x.
At the end, there is at most 1 stone left.  Return the weight of this stone (or 0 if there are no stones left.)

 

Example 1:

Input: [2,7,4,1,8,1]
Output: 1
Explanation: 
We combine 7 and 8 to get 1 so the array converts to [2,4,1,1,1] then,
we combine 2 and 4 to get 2 so the array converts to [2,1,1,1] then,
we combine 2 and 1 to get 1 so the array converts to [1,1,1] then,
we combine 1 and 1 to get 0 so the array converts to [1] then that's the value of last stone.
 

Note:

1 <= stones.length <= 30
1 <= stones[i] <= 1000


## 方法


```go

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

// PQ implements heap.Interface and holds entries.
type PQ []int

func (pq PQ) Len() int { return len(pq) }

func (pq PQ) Less(i, j int) bool { return pq[i] > pq[j] }

func (pq PQ) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

// Push 往 pq 中放 int
func (pq *PQ) Push(x interface{}) {
	temp := x.(int)
	*pq = append(*pq, temp)
}

// Pop 从 pq 中取出最优先的 int
func (pq *PQ) Pop() interface{} {
	temp := (*pq)[len(*pq)-1]
	*pq = (*pq)[0 : len(*pq)-1]
	return temp
}
```



```python
class Solution(object):
    def lastStoneWeight(self, stones):
        """
        :type stones: List[int]
        :rtype: int
        """
        max_heap = [-x for x in stones]
        heapq.heapify(max_heap)
        for i in xrange(len(stones)-1):
            x, y = -heapq.heappop(max_heap), -heapq.heappop(max_heap)
            heapq.heappush(max_heap, -abs(x-y))
        return -max_heap[0]
```