78. Subsets


Medium


Given a set of distinct integers, nums, return all possible subsets (the power set).

Note: The solution set must not contain duplicate subsets.

Example:

```
Input: nums = [1,2,3]
Output:
[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]
```


## 方法

```go
func subsets(nums []int) [][]int {
    res := make([][]int, 1, 1024)
	for _, n := range nums {
		for _, r := range res {
			res = append(res, append([]int{n}, r...))
		}
	}
	return res
}
```



```python
class Solution(object):
    def subsets(self, nums):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        nums.sort()
        result = [[]]
        for i in xrange(len(nums)):
            size = len(result)
            for j in xrange(size):
                result.append(list(result[j]))
                result[-1].append(nums[i])
        return result
```