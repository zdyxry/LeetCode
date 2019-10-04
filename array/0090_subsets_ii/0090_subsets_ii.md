90. Subsets II


Medium


Given a collection of integers that might contain duplicates, nums, return all possible subsets (the power set).

Note: The solution set must not contain duplicate subsets.

Example:

```
Input: [1,2,2]
Output:
[
  [2],
  [1],
  [1,2,2],
  [2,2],
  [1,2],
  []
]
```

## 方法

```go
func subsetsWithDup(nums []int) [][]int {
	c, res := []int{}, [][]int{}
	sort.Ints(nums) // 这里是去重的关键逻辑
	for k := 0; k <= len(nums); k++ {
		generateSubsetsWithDup(nums, k, 0, c, &res)
	}
	return res
}

func generateSubsetsWithDup(nums []int, k, start int, c []int, res *[][]int) {
	if len(c) == k {
		b := make([]int, len(c))
		copy(b, c)
		*res = append(*res, b)
		return
	}
	// i will at most be n - (k - c.size()) + 1
	for i := start; i < len(nums)-(k-len(c))+1; i++ {
		fmt.Printf("i = %v start = %v c = %v\n", i, start, c)
		if i > start && nums[i] == nums[i-1] { // 这里是去重的关键逻辑,本次不取重复数字，下次循环可能会取重复数字
			continue
		}
		c = append(c, nums[i])
		generateSubsetsWithDup(nums, k, i+1, c, res)
		c = c[:len(c)-1]
	}
	return
}
```


```python
class Solution(object):
    def subsetsWithDup(self, nums):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        nums.sort()
        result = [[]]
        previous_size = 0
        for i in xrange(len(nums)):
            size = len(result)
            for j in xrange(size):
                # Only union non-duplicate element or new union set.
                if i == 0 or nums[i] != nums[i - 1] or j >= previous_size:
                    result.append(list(result[j]))
                    result[-1].append(nums[i])
            previous_size = size
        return result
```