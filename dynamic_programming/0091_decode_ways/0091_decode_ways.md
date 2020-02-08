91. Decode Ways


Medium


A message containing letters from A-Z is being encoded to numbers using the following mapping:

```
'A' -> 1
'B' -> 2
...
'Z' -> 26
Given a non-empty string containing only digits, determine the total number of ways to decode it.
```

Example 1:

```
Input: "12"
Output: 2
Explanation: It could be decoded as "AB" (1 2) or "L" (12).
```

Example 2:

```
Input: "226"
Output: 3
Explanation: It could be decoded as "BZ" (2 26), "VF" (22 6), or "BBF" (2 2 6).
```


## 方法

```go
func numDecodings(s string) int {
    if s[0] == '0' {
		return 0
	}

	n := len(s)
	dp := make([]int, n+1)
	// dp[i] 表示 s[:i] 的组成方式数
	dp[0], dp[1] = 1, 1
	lastOne, lastTwo := int(s[0]-'0'), 0
	for i := 2; i <= n; i++ {
		last := int(s[i-1] - '0')
		lastOne, lastTwo = last, lastOne*10+last
		if 1 <= lastOne {
			dp[i] = dp[i-1]
		}
		if 10 <= lastTwo && lastTwo <= 26 {
			dp[i] += dp[i-2]
		}
	}

	return dp[n]
}
```

```go
func numDecodings(s string) int {
    if len(s) == 0 {
		return 0
	}
	dp := make([]int, len(s)+1)
	dp[0] = 1
	if s[:1] == "0" {
		dp[1] = 0
	} else {
		dp[1] = 1
	}
	for i := 2; i <= len(s); i++ {
		lastNum, _ := strconv.Atoi(s[i-1 : i])
		if lastNum >= 1 && lastNum <= 9 {
			dp[i] += dp[i-1]
		}
		lastNum, _ = strconv.Atoi(s[i-2 : i])
		if lastNum >= 10 && lastNum <= 26 {
			dp[i] += dp[i-2]
		}
	}
	return dp[len(s)]
}
```


```python
class Solution(object):
    def numDecodings(self, s):
        """
        :type s: str
        :rtype: int
        """
        if len(s) == 0 or s[0] == '0':
            return 0
        prev, prev_prev = 1, 0
        for i in xrange(len(s)):
            cur = 0
            if s[i] != '0':
                cur = prev
            if i > 0 and (s[i - 1] == '1' or (s[i - 1] == '2' and s[i] <= '6')):
                cur += prev_prev
            prev, prev_prev = cur, prev
        return prev
```