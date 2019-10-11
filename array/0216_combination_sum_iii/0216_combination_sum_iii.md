216. Combination Sum III


Medium


Find all possible combinations of k numbers that add up to a number n, given that only numbers from 1 to 9 can be used and each combination should be a unique set of numbers.

Note:

All numbers will be positive integers.

The solution set must not contain duplicate combinations.


Example 1:

```
Input: k = 3, n = 7
Output: [[1,2,4]]
```


Example 2:

```
Input: k = 3, n = 9
Output: [[1,2,6], [1,3,5], [2,3,4]]
```

## 方法

```go
func combinationSum3(k int, n int) [][]int {
    if k == 0 {
		return [][]int{}
	}
	c, res := []int{}, [][]int{}
	findcombinationSum3(k, n, 1, c, &res)
	return res
}

func findcombinationSum3(k, target, index int, c []int, res *[][]int) {
	if target == 0 {
		if len(c) == k {
			b := make([]int, len(c))
			copy(b, c)
			*res = append(*res, b)
		}
		return
	}
	for i := index; i < 10; i++ {
		if target >= i {
			c = append(c, i)
			findcombinationSum3(k, target-i, i+1, c, res)
			c = c[:len(c)-1]
		}
	}
}
```


```python
class Solution(object):
    def combinationSum3(self, k, n):
        """
        :type k: int
        :type n: int
        :rtype: List[List[int]]
        """
        paths = []
        self.helper(k, n, 0, 1, 0, [], paths)
        return paths
    
    def helper(self, k, n, cur_sum, cur_num, cur_cnt, path, paths):
        if cur_sum == n:
            if cur_cnt == k:
                paths.append(path[:])
            return
        
        for next_num in xrange(cur_num, 10):
            if cur_sum + next_num <= n and cur_cnt + 1 <= k:
                path.append(next_num)
                self.helper(k, n, cur_sum+next_num, next_num+1, cur_cnt+1, path, paths)
                path.pop()
```