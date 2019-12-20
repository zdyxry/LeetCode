47. Permutations II


Medium


Given a collection of numbers that might contain duplicates, return all possible unique permutations.

Example:

```
Input: [1,1,2]
Output:
[
  [1,1,2],
  [1,2,1],
  [2,1,1]
]
```


## 方法


```go
func permuteUnique(nums []int) [][]int {
    sort.Ints(nums)

	n := len(nums)
	// vector 是一组可能的解答
	vector := make([]int, n)
	// taken[i] == true 表示 nums[i] 已经出现在了 vector 中
	taken := make([]bool, n)

	var ans [][]int

	makePermutation(0, n, nums, vector, taken, &ans)

	return ans
}

// 这个方式和我原来的方式相比，
// 增加了比较的次数
// 但是，减少了切片生成的次数。
// 所以，速度会快一些。
func makePermutation(cur, n int, nums, vector []int, taken []bool, ans *[][]int) {
	if cur == n {
		tmp := make([]int, n)
		copy(tmp, vector)
		*ans = append(*ans, tmp)
		return
	}

	used := make(map[int]bool, n-cur)

	for i := 0; i < n; i++ {

		if !taken[i] && !used[nums[i]] {
			used[nums[i]] = true

			// 准备使用 nums[i]，所以，taken[i] == true
			taken[i] = true
			// NOTICE: 是 vector[cur]
			vector[cur] = nums[i]

			makePermutation(cur+1, n, nums, vector, taken, ans)

			// 下一个循环中
			// vector[cur] = nums[i+1]
			// 所以，在这个循环中，恢复 nums[i] 自由
			taken[i] = false
		}
	}
}
```


```python
class Solution(object):
    def permuteUnique(self, nums):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        res = []
        nums.sort()
        self.dfs(nums, [], res)
        return res
    
    def dfs(self, remain, path, res):
        if not remain:
            res.append(path)
            return
        for i in range(len(remain)):
            if i>0 and remain[i-1]==remain[i]:
                continue
            left = remain[:i] + remain[i+1:]
            self.dfs(left, path+[remain[i]], res)
```