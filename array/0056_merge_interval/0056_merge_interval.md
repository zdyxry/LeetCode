56. Merge Intervals


Medium


Given a collection of intervals, merge all overlapping intervals.

Example 1:

```
Input: [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].
```

Example 2:

```
Input: [[1,4],[4,5]]
Output: [[1,5]]
Explanation: Intervals [1,4] and [4,5] are considered overlapping.
```


## 方法

```go
func merge(intervals [][]int) [][]int {
	for i := 0; i < len(intervals); i++ {
		for j := i + 1; j < len(intervals); j++ {
			a := intervals[i]
			b := intervals[j]
			if a[1] >= b[0] && a[0] <= b[1] {
				// 有重叠
				arr := []int{
					a[0],
					a[1],
					b[0],
					b[1],
				}
				sort.Ints(arr)

				intervals[i] = []int{arr[0], arr[3]}
				// delete j
				intervals = append(intervals[:j], intervals[j+1:]...)
				j = i
			}
		}
	}
	return intervals
}
```



```python
class Solution(object):
    def merge(self, intervals):
        """
        :type intervals: List[List[int]]
        :rtype: List[List[int]]
        """
        if not intervals:
            return []
        
        length = len(intervals)
        if length == 1:
            return intervals
        intervals.sort(key=lambda x:x[0])  ## or intervals.sort()
        res = [intervals[0]]
    
        for i in range(length):
			# compare with the last
            if intervals[i][0] <= res[-1][1]:
				# take the maximum
                res[-1][1] = max(intervals[i][1], res[-1][1])
            else:
                res.append(intervals[i])
                    
        return res
```