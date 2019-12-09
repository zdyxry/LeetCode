907. Sum of Subarray Minimums


Medium

Given an array of integers A, find the sum of min(B), where B ranges over every (contiguous) subarray of A.

Since the answer may be large, return the answer modulo 10^9 + 7.

 

Example 1:

```
Input: [3,1,2,4]
Output: 17
Explanation: Subarrays are [3], [1], [2], [4], [3,1], [1,2], [2,4], [3,1,2], [1,2,4], [3,1,2,4]. 
Minimums are 3, 1, 2, 4, 1, 1, 2, 1, 1, 1.  Sum is 17.
```

Note:

1 <= A.length <= 30000   
1 <= A[i] <= 30000   


## 方法


```go
func sumSubarrayMins(A []int) int {
    res, n, mod := 0, len(A), 1000000007
	lefts, rights, leftStack, rightStack := make([]int, n), make([]int, n), []*pair{}, []*pair{}
	for i := 0; i < n; i++ {
		count := 1
		for len(leftStack) != 0 && leftStack[len(leftStack)-1].val > A[i] {
			count += leftStack[len(leftStack)-1].count
			leftStack = leftStack[:len(leftStack)-1]
		}
		leftStack = append(leftStack, &pair{val: A[i], count: count})
		lefts[i] = count
	}

	for i := n - 1; i >= 0; i-- {
		count := 1
		for len(rightStack) != 0 && rightStack[len(rightStack)-1].val >= A[i] {
			count += rightStack[len(rightStack)-1].count
			rightStack = rightStack[:len(rightStack)-1]
		}
		rightStack = append(rightStack, &pair{val: A[i], count: count})
		rights[i] = count
	}
    
    fmt.Println(lefts, rights)

	for i := 0; i < n; i++ {
		res = (res + A[i]*lefts[i]*rights[i]) % mod
	}
	return res
}

type pair struct {
	val   int
	count int
}
```



```go
func sumSubarrayMins(A []int) int {
    stack, dp, res, mod := []int{}, make([]int, len(A)+1), 0, 1000000007
	stack = append(stack, -1)

	for i := 0; i < len(A); i++ {
		for stack[len(stack)-1] != -1 && A[i] <= A[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		dp[i+1] = (dp[stack[len(stack)-1]+1] + (i-stack[len(stack)-1])*A[i]) % mod
		stack = append(stack, i)
		res += dp[i+1]
		res %= mod
	}
	return res
}

type pair struct {
	val   int
	count int
}
```



```python
class Solution(object):
    def sumSubarrayMins(self, A):
        """
        :type A: List[int]
        :rtype: int
        """
        dot = 0
        sums = 0
        stack = []
        
        for elem in A:
            new_count = 1
            while stack and stack[-1][0] >= elem:
                val, count = stack.pop()
                dot -= val * count
                new_count += count
            
            stack.append((elem, new_count))
            dot += elem * new_count
            sums += dot
        
        return sums % (10**9+7)
```