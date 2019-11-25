583. Delete Operation for Two Strings


Medium


Given two words word1 and word2, find the minimum number of steps required to make word1 and word2 the same, where in each step you can delete one character in either string.

Example 1:
```
Input: "sea", "eat"
Output: 2
Explanation: You need one step to make "sea" to "ea" and another step to make "eat" to "ea".
```

Note:   
The length of given words won't exceed 500.   
Characters in given words can only be lower-case letters.   

## 方法

```go
func minDistance(word1 string, word2 string) int {
    n1, n2 := len(word1), len(word2)

	// dp[i][j] == k 表示 word1[:i] 和 word2[:j] 的最大公共子序列的长度为 k
	dp := make([][]int, n1+1)
	for i := range dp {
		dp[i] = make([]int, n2+1)
	}

	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
		}
	}

	return n1 + n2 - dp[n1][n2]*2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```


```go
/*动态规划

方法1:最长公共子串
dp[i][j]表示word1[1, i]和word2[1, j]的最长公共子串;
if word1[i-1] == word2[j-1], dp[i][j] = dp[i-1][j-1] + 1;
if word1[i-1] != word2[j-1], dp[i][j] = max(dp[i-1][j], dp[i][j-1]);
and dp[0][j] = dp[i][0] = 0;
result = word1.length + word2.length - 2 * dp[word1.length][word2.length]

方法2:最小修改操作
dp[i][j]表示word1[1, i]和word2[1, j]的最小修改操作;
if word1[i-1] == word2[j-1], dp[i][j] = dp[i-1][j-1];
if word1[i-1] != word2[j-1], dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + 1;
and dp[0][j] = j, dp[i][0] = i => i == 0 || j == 0, dp[i][j] = i + j;
result = dp[word1.length][word2.length]
*/

func minDistance(word1 string, word2 string) int {
    dp := make([][]int, len(word1) + 1)
    for i := 0; i <= len(word1); i++ {
        dp[i] = make([]int, len(word2) + 1)
    }
    for i := 0; i <= len(word1); i++ {
        for j := 0; j <= len(word2); j++ {
            if i == 0 || j == 0 {
                dp[i][j] = i + j
            } else if word1[i-1] == word2[j-1] {
                dp[i][j] = dp[i-1][j-1]
            } else {
                dp[i][j] = Min(dp[i-1][j], dp[i][j-1]) + 1
            }
        }
    }
    return dp[len(word1)][len(word2)]
}

func Min(x, y int) int {
    if x < y {
        return x
    } else {
        return y
    }
}
```



```python
class Solution(object):
    def minDistance(self, word1, word2):
        """
        :type word1: str
        :type word2: str
        :rtype: int
        """
        len1, len2 = len(word1), len(word2)
        
        matrix = [[0]*(len2+1) for i in range (len1+1)]
        
        for i in range(1, len1+1):
            for j in range(1, len2+1):
                if word1[i-1] == word2[j-1]:
                    matrix[i][j] = matrix[i-1][j-1] + 1
                else:
                    matrix[i][j] = max(matrix[i][j-1], matrix[i-1][j])
        return len1 + len2 - 2*(matrix[-1][-1])
```