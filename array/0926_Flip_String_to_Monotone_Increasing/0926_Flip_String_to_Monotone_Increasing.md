926. Flip String to Monotone Increasing


Medium


A string of '0's and '1's is monotone increasing if it consists of some number of '0's (possibly 0), followed by some number of '1's (also possibly 0.)

We are given a string S of '0's and '1's, and we may flip any '0' to a '1' or a '1' to a '0'.

Return the minimum number of flips to make S monotone increasing.

 

Example 1:

```
Input: "00110"
Output: 1
Explanation: We flip the last digit to get 00111.
```

Example 2:

```
Input: "010110"
Output: 2
Explanation: We flip to get 011111, or alternatively 000111.
```

Example 3:

```
Input: "00011000"
Output: 2
Explanation: We flip to get 00000000.
```

Note:

1 <= S.length <= 20000  
S only consists of '0' and '1' characters.


## 方法

```go
func minFlipsMonoIncr(S string) int {
    leftOnes := make([]int, len(S)+1)
	rightZeros := make([]int, len(S)+1)

	for index, ch := range S {
		leftOnes[index+1] = leftOnes[index]
		if ch == '1' {
			leftOnes[index+1]++
		}
	}

	for index := len(S) - 1; index >= 0; index-- {
		rightZeros[index] = rightZeros[index+1]
		if S[index] == '0' {
			rightZeros[index]++
		}
	}

	minFunc := func(a int, b int) int {
		if a > b {
			return b
		} else {
			return a
		}
	}

	ans := minFunc(leftOnes[len(S)], rightZeros[0])
	for i := 0; i < len(S); i++ {
		turnCount := leftOnes[i+1] + rightZeros[i+1]
		ans = minFunc(ans, turnCount)
	}
	return ans
}
```



```python
class Solution:
    def minFlipsMonoIncr(self, S: str) -> int:
        N = len(S)
        dp = [0] * 2
        for i in range(1, N + 1):
            if S[i - 1] == '0':
                dp[0] = dp[0]
                dp[1] = min(dp[1], dp[0]) + 1
            else:
                dp[1] = min(dp[1], dp[0])
                dp[0] = dp[0] + 1
        return min(dp[0], dp[1])
```