228. Summary Ranges


Medium


Given a sorted integer array without duplicates, return the summary of its ranges.

Example 1:

```
Input:  [0,1,2,4,5,7]
Output: ["0->2","4->5","7"]
Explanation: 0,1,2 form a continuous range; 4,5 form a continuous range.
```

Example 2:

```
Input:  [0,2,3,4,6,8,9]
Output: ["0","2->4","6","8->9"]
Explanation: 2,3,4 form a continuous range; 8,9 form a continuous range.
```

## 方法

```go
func summaryRanges(a []int) []string {
    res := []string{}

	l := len(a)
	if l == 0 {
		return res
	}

	begin := a[0]
	str := ""

	for i := 0; i < l; i++ {
		if i == l-1 || a[i]+1 != a[i+1] {
			if a[i] == begin {
				str = fmt.Sprintf("%d", begin)
			} else {
				str = fmt.Sprintf("%d->%d", begin, a[i])
			}

			if i+1 < l {
				begin = a[i+1]
			}

			res = append(res, str)
		}
	}

	return res
}
```


```python
class Solution(object):
    def summaryRanges(self, nums):
        """
        :type nums: List[int]
        :rtype: List[str]
        """
        ranges = []
        if not nums:
            return ranges

        start, end = nums[0], nums[0]
        for i in xrange(1, len(nums) + 1):
            if i < len(nums) and nums[i] == end + 1:
                end = nums[i]
            else:
                interval = str(start)
                if start != end:
                    interval += "->" + str(end)
                ranges.append(interval)
                if i < len(nums):
                    start = end = nums[i]

        return ranges
```