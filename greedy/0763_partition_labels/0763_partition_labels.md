763. Partition Labels


Medium


A string S of lowercase English letters is given. We want to partition this string into as many parts as possible so that each letter appears in at most one part, and return a list of integers representing the size of these parts.

 

Example 1:

```
Input: S = "ababcbacadefegdehijhklij"
Output: [9,7,8]
Explanation:
The partition is "ababcbaca", "defegde", "hijhklij".
This is a partition so that each letter appears in at most one part.
A partition like "ababcbacadefegde", "hijhklij" is incorrect, because it splits S into less parts.
```
 

Note:

S will have length in range [1, 500].  
S will consist of lowercase English letters ('a' to 'z') only.


## 方法

```go
func partitionLabels(S string) []int {
	// Use a 26 length slice to
	// store all leters' frequency.
	lastIndex := make([]int, 26)
	for i, c := range S {
		lastIndex[int(c - 'a')] = i
	}
	ans := make([]int, 0)
	start := 0
	end := 0
	for i := 0; i < len(S); i++ {
		end = max(lastIndex[int(S[i] - 'a')], end)
		if i == end {
			ans = append(ans, end-start+1)
			start = end + 1
		}
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

```


```python
class Solution:
    def partitionLabels(self, S: str) -> List[int]:
        last = {c: i for i, c in enumerate(S)}
        j = anchor = 0
        ans = []
        for i, c in enumerate(S):
            j = max(j, last[c])
            if i == j:
                ans.append(i - anchor + 1)
                anchor = i + 1
            
        return ans

```