1497. Check If Array Pairs Are Divisible by k


Medium


Given an array of integers arr of even length n and an integer k.

We want to divide the array into exactly n / 2 pairs such that the sum of each pair is divisible by k.

Return True If you can find a way to do that or False otherwise.

 

Example 1:

```
Input: arr = [1,2,3,4,5,10,6,7,8,9], k = 5
Output: true
Explanation: Pairs are (1,9),(2,8),(3,7),(4,6) and (5,10).
```

Example 2:

```
Input: arr = [1,2,3,4,5,6], k = 7
Output: true
Explanation: Pairs are (1,6),(2,5) and(3,4).
```

Example 3:

```
Input: arr = [1,2,3,4,5,6], k = 10
Output: false
Explanation: You can try all possible pairs to see that there is no way to divide arr into 3 pairs each with sum divisible by 10.
```

Example 4:

```
Input: arr = [-10,10], k = 2
Output: true
```

Example 5:

```
Input: arr = [-1,1,-2,2,-3,3,-4,4], k = 3
Output: true
```

Constraints:

arr.length == n  
1 <= n <= 10^5  
n is even.  
-10^9 <= arr[i] <= 10^9  
1 <= k <= 10^5

## 方法


```go
func canArrange(arr []int, k int) bool {
    mod := make([]int, k)
	n := len(arr)
	for i := 0; i < n; i++ {
		mod[(arr[i]%k+k)%k]++
	}
	if mod[0] % 2 > 0 {
		return false
	}
	for i := 1; i <= k/2; i++ {
		if 2 * i == k {
			if mod[i] % 2 > 0 {
				return false
			}
		} else {
			if mod[i] != mod[k-i] {
				return false
			}
		}
	}
	return true
}
```


```python
class Solution:
    def canArrange(self, arr: List[int], k: int) -> bool:
        mod = collections.Counter(num % k for num in arr)
        for t, occ in mod.items():
            if t > 0 and (k - t not in mod or mod[k - t] != occ):
                return False
        return mod[0] % 2 == 0

```