769. Max Chunks To Make Sorted


Medium


Given an array arr that is a permutation of [0, 1, ..., arr.length - 1], we split the array into some number of "chunks" (partitions), and individually sort each chunk.  After concatenating them, the result equals the sorted array.

What is the most number of chunks we could have made?

Example 1:

```
Input: arr = [4,3,2,1,0]
Output: 1
Explanation:
Splitting into two or more chunks will not return the required result.
For example, splitting into [4, 3], [2, 1, 0] will result in [3, 4, 0, 1, 2], which isn't sorted.
```

Example 2:

```
Input: arr = [1,0,2,3,4]
Output: 4
Explanation:
We can split into two chunks, such as [1, 0], [2, 3, 4].
However, splitting into [1, 0], [2], [3], [4] is the highest number of chunks possible.
```

Note:

arr will have length in range [1, 10].   
arr[i] will be a permutation of [0, 1, ..., arr.length - 1].



## 方法


```go
func maxChunksToSorted(arr []int) int {
    if len(arr) < 2 {
        return len(arr)
    }

    count := 0
    cur := arr[0]
    for i, v := range arr {
        cur = max(cur, v)
        if cur == i {
            count++
        }
    }
    return count
}
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

```


```python
class Solution:
    def maxChunksToSorted(self, arr: List[int]) -> int:
        A , B = 0 , 0
        for i,j in enumerate(arr):
            A = max(A,j)
            if(A == i):B += 1
        return B 
```