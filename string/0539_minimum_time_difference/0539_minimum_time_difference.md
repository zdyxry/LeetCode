539. Minimum Time Difference


Medium

Given a list of 24-hour clock time points in "Hour:Minutes" format, find the minimum minutes difference between any two time points in the list.

Example 1:
```
Input: ["23:59","00:00"]
Output: 1
```

Note:  
The number of time points in the given list is at least 2 and won't exceed 20000.  
The input time is legal and ranges from 00:00 to 23:59.  

## 方法


```go
func findMinDifference(timePoints []string) int {
    n := len(timePoints)
	// 把时间按升序排列
	sort.Strings(timePoints)

	// 把 minDiff 设置为 timePoints[n-1] 到 timePoints[0] 的分钟数
	minDiff := 24*60 - diff(timePoints[0], timePoints[n-1])

	for i := 1; i < n; i++ {
		minDiff = min(minDiff, diff(timePoints[i-1], timePoints[i]))
	}

	return minDiff
}

// 返回从 a 到 b 的分钟数
func diff(a, b string) int {
	ha, ma := getHourAndMintueOf(a)
	hb, mb := getHourAndMintueOf(b)
	return (hb-ha)*60 + (mb - ma)
}

func getHourAndMintueOf(s string) (hour, minute int) {
	hour, _ = strconv.Atoi(s[:2])
	minute, _ = strconv.Atoi(s[3:])
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```


```python
class Solution(object):
    def findMinDifference(self, timePoints):
        """
        :type timePoints: List[str]
        :rtype: int
        """
        if len(timePoints) > 1440:
            return 0
        s = sorted(map(lambda t: int(t[:2]) * 60 + int(t[3:]), timePoints))
        return min(s2 - s1 for s1, s2 in zip(s, s[1:] + [1440+s[0]]))
```


```python
class Solution(object):
    def findMinDifference(self, timePoints):
        """
        :type timePoints: List[str]
        :rtype: int
        """
        if len(timePoints) > 1440:
            return 0
        buckets = [0] * 1440
        for tp in timePoints:
            seconds = int(tp[:2]) * 60 + int(tp[3:])
            buckets[seconds] += 1
            if buckets[seconds] > 1:
                return 0
        s = [i for i, cnt in enumerate(buckets) if cnt]
        return min(s2 - s1 for s1, s2 in zip(s, s[1:] + [1440 + s[0]]))
```