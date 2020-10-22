821. Shortest Distance to a Character


Easy


Given a string S and a character C, return an array of integers representing the shortest distance from the character C in the string.

Example 1:

```
Input: S = "loveleetcode", C = 'e'
Output: [3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0]
```

Note:

S string length is in [1, 10000].   
C is a single character, and guaranteed to be in string S.   
All letters in S and C are lowercase.


## 方法


```go
func shortestToChar(S string, C byte) []int {
	var l,r int
	ans := []int{}
	for i := 0;i < len(S);i++ {
		l = i
		r = i
		mindist := 0
		for {
			if S[l] == C || S[r] == C {
				ans = append(ans,mindist)
				break
			}
			if l - 1 >= 0 {
				l--
			}
			if r + 1 <len(S) {
				r++
			}
			mindist++
		}
	}
	return ans
}

```

```python
class Solution(object):
    def shortestToChar(self, S, C):
        prev = float('-inf')
        ans = []
        for i, x in enumerate(S):
            if x == C: prev = i
            ans.append(i - prev)

        prev = float('inf')
        for i in range(len(S) - 1, -1, -1):
            if S[i] == C: prev = i
            ans[i] = min(ans[i], prev - i)

        return ans
```