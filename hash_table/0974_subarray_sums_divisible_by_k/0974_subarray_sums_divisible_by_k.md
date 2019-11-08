974. Subarray Sums Divisible by K


Medium


Given an array A of integers, return the number of (contiguous, non-empty) subarrays that have a sum divisible by K.

 

Example 1:
```
Input: A = [4,5,0,-2,-3,1], K = 5
Output: 7
Explanation: There are 7 subarrays with a sum divisible by K = 5:
[4, 5, 0, -2, -3, 1], [5], [5, 0], [5, 0, -2, -3], [0], [0, -2, -3], [-2, -3]
```

Note:

1 <= A.length <= 30000  
-10000 <= A[i] <= 10000  
2 <= K <= 10000  


## 方法


```go
func subarraysDivByK(A []int, K int) int {
    count := [10000]int{}
	count[0] = 1
	prefix, res := 0, 0
	for _, a := range A {
		prefix = (prefix + a%K + K) % K
		res += count[prefix] // 减去相同的 prefix，剩下的就都是 K 的倍数
		count[prefix]++
	}
	return res
}
```



```python
class Solution(object):
    def subarraysDivByK(self, A, K):
        """
        :type A: List[int]
        :type K: int
        :rtype: int
        """
        count = collections.defaultdict(int)
        count[0] = 1
        result, prefix = 0, 0
        for a in A:
            prefix = (prefix+a) % K
            result += count[prefix]
            count[prefix] += 1
        return result
```