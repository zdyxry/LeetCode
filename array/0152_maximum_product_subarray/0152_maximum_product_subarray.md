152. Maximum Product Subarray


Medium


Given an integer array nums, find the contiguous subarray within an array (containing at least one number) which has the largest product.

Example 1:
```
Input: [2,3,-2,4]
Output: 6
Explanation: [2,3] has the largest product 6.
```


Example 2:

```
Input: [-2,0,-1]
Output: 0
Explanation: The result cannot be 2, because [-2,-1] is not a subarray.
```

## 方法

```go
func maxProduct(a []int) int {
    cur, neg, max := 1, 1, a[0]

	for i := 0; i < len(a); i++ {

		switch {
		case a[i] > 0:
			cur, neg = a[i]*cur, a[i]*neg
		case a[i] < 0:
			cur, neg = a[i]*neg, a[i]*cur
		default:
			cur, neg = 0, 1
		}

		if max < cur {
			max = cur
		}

		if cur <= 0 {
			cur = 1
		}
	}

	return max
}
```



```python
class Solution(object):
    def maxProduct(self, A):
        """
        :type nums: List[int]
        :rtype: int
        """
        global_max, local_max, local_min = float("-inf"), 1, 1
        for x in A:
            local_max, local_min = max(x, local_max * x, local_min * x), min(x, local_max * x, local_min * x)
            global_max = max(global_max, local_max)
        return global_max
```