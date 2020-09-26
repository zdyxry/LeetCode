1288. Remove Covered Intervals


Medium


Given a list of intervals, remove all intervals that are covered by another interval in the list. Interval [a,b) is covered by interval [c,d) if and only if c <= a and b <= d.

After doing so, return the number of remaining intervals.

 

Example 1:

```
Input: intervals = [[1,4],[3,6],[2,8]]
Output: 2
Explanation: Interval [3,6] is covered by [2,8], therefore it is removed.
```

Constraints:

1 <= intervals.length <= 1000  
0 <= intervals[i][0] < intervals[i][1] <= 10^5  
intervals[i] != intervals[j] for all i != j


## 方法

[x, y]按照x升序 y降序排列
显然后面所有左边界和它一样的区间都将被删除
若小于 右置位最大值， 一样被删除

```go
func removeCoveredIntervals(intervals [][]int) int {
    if len(intervals) == 0 || len(intervals[0]) == 0 {
        return 0
    }
    sort.Slice(intervals, func(i, j int) bool { 
        if intervals[i][0] != intervals[j][0] {
            return intervals[i][0] < intervals[j][0]
        }else {
            return intervals[i][1] > intervals[j][1]
        }
    })
    result := 0 
    max := intervals[0][1]
    for i:=1; i < len(intervals); i++ {
        if intervals[i][0] == intervals[i-1][0] {
            result++ 
            continue
        }
        if max < intervals[i][1] {
            max = intervals[i][1]
        } else  {
            result++
        }
    }
    
    return len(intervals) - result
}

```


```python
class Solution:
    def removeCoveredIntervals(self, intervals: List[List[int]]) -> int:
        # Sort by start point.
        # If two intervals share the same start point
        # put the longer one to be the first.
        intervals.sort(key = lambda x: (x[0], -x[1]))
        count = 0
        
        prev_end = 0
        for _, end in intervals:
            # if current interval is not covered
            # by the previous one
            if end > prev_end:
                count += 1    
                prev_end = end
        
        return count
```