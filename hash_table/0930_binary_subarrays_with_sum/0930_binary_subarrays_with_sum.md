930. Binary Subarrays With Sum


Medium


In an array A of 0s and 1s, how many non-empty subarrays have sum S?

 

Example 1:

```
Input: A = [1,0,1,0,1], S = 2
Output: 4
Explanation: 
The 4 subarrays are bolded below:
[1,0,1,0,1]
[1,0,1,0,1]
[1,0,1,0,1]
[1,0,1,0,1]
```

Note:

A.length <= 30000

0 <= S <= A.length

A[i] is either 0 or 1.


## 方法

```go
func numSubarraysWithSum(A []int, S int) int {
    preSum, res := 0, 0
	count := make([]int, len(A)+1)
	count[0] = 1
	for _, n := range A {
		preSum += n
		if preSum >= S {
			res += count[preSum-S]
		}
		count[preSum]++
	}
	return res
}
```


```python
class Solution(object):
    def numSubarraysWithSum(self, A, S):
        """
        :type A: List[int]
        :type S: int
        :rtype: int
        """
        result = 0
        left, right, sum_left, sum_right = 0, 0, 0, 0
        for i, a in enumerate(A):
            sum_left += a
            while left < i and sum_left > S:
                sum_left -= A[left]
                left += 1
            sum_right += a
            while right < i and \
                  (sum_right > S or (sum_right == S and not A[right])):
                sum_right -= A[right]
                right += 1
            if sum_left == S:
                result += right-left+1
        return result
```