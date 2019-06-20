594. Longest Harmonious Subsequence

Easy

We define a harmounious array as an array where the difference between its maximum value and its minimum value is exactly 1.

Now, given an integer array, you need to find the length of its longest harmonious subsequence among all its possible subsequences.

Example 1:

```
Input: [1,3,2,2,5,2,3,7]
Output: 5
Explanation: The longest harmonious subsequence is [3,2,2,2,3].
```
 

Note: 

The length of the input array will not exceed 20,000.


# 方法

遍历数组，记录当前数字出现次数；遍历 map，查找当前数值 +1 是否在 map 中，若在，则max=max(max, c1 + c2)。


```go
func findLHS(nums []int) int {
    r := make(map[int]int, len(nums))
	for _, n := range nums {
		r[n]++
	}

	max := 0
	for n, c1 := range r {
		c2, ok := r[n+1]
		if ok {
			t := c1 + c2
			if max < t {
				max = t
			}
		}
	}

	return max
}
```


```python
class Solution(object):
    def findLHS(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        if not nums:
            return 0
        
        d = {}
        for n in nums:
            if n in d:
                d[n] += 1
            else:
                d[n] = 1

        max_length = 0
        for n in d:
            if n+1 in d:
                max_length = max(max_length, d[n + 1] + d[n])
            
        return max_length

```