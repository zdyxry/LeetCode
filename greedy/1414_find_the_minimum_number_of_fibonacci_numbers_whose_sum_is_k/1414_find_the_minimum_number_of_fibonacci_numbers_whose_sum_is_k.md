1414. Find the Minimum Number of Fibonacci Numbers Whose Sum Is K


Medium


Given the number k, return the minimum number of Fibonacci numbers whose sum is equal to k, whether a Fibonacci number could be used multiple times.

The Fibonacci numbers are defined as:

F1 = 1  
F2 = 1  
Fn = Fn-1 + Fn-2 , for n > 2.  
It is guaranteed that for the given constraints we can always find such fibonacci numbers that sum k.
 

Example 1:

```
Input: k = 7
Output: 2 
Explanation: The Fibonacci numbers are: 1, 1, 2, 3, 5, 8, 13, ... 
For k = 7 we can use 2 + 5 = 7.
```

Example 2:

```
Input: k = 10
Output: 2 
Explanation: For k = 10 we can use 2 + 8 = 10.
```

Example 3:

```
Input: k = 19
Output: 3 
Explanation: For k = 19 we can use 1 + 5 + 13 = 19.
```

Constraints:

1 <= k <= 10^9


## 方法


```go
func findMinFibonacciNumbers(k int) int {
    nums := []int{1, 1}
	for i := 2; ; i++ {
		v := nums[i-1] + nums[i-2]
		if v > k {
			break
		}
		nums = append(nums, v)
	}
	count := 0
	for k > 0 {
		k -= nums[len(nums)-1]
		count++
		if k == 0 {
			break
		}
		i := len(nums) - 1
		for ; i >= 0 && k < nums[i]; i-- { }
		nums = nums[:i+1]
	}
	return count
}
```


```python
class Solution:
    def findMinFibonacciNumbers(self, k: int) -> int:
        ls=self.fib(k)
        res=0
        while k:
            if k>=ls[-1]:
                k-=ls[-1]
                res+=1
            else:
                ls.pop()
        return res
            
    def fib(self, N) -> int:
        a, b = 0, 1
        res = []
        while b <= N:
            res.append(b)
            a, b = b, a + b
        return res
```