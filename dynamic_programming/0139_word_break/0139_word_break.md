139. Word Break


Medium


Given a non-empty string s and a dictionary wordDict containing a list of non-empty words, determine if s can be segmented into a space-separated sequence of one or more dictionary words.

Note:

The same word in the dictionary may be reused multiple times in the segmentation.

You may assume the dictionary does not contain duplicate words.

Example 1:

```
Input: s = "leetcode", wordDict = ["leet", "code"]
Output: true
Explanation: Return true because "leetcode" can be segmented as "leet code".
```


Example 2:

```
Input: s = "applepenapple", wordDict = ["apple", "pen"]
Output: true
Explanation: Return true because "applepenapple" can be segmented as "apple pen apple".
             Note that you are allowed to reuse a dictionary word.
```


Example 3:

```
Input: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
Output: false
```


## 方法

```go
func wordBreak(s string, wordDict []string) bool {
    if len(wordDict) == 0 {
		return false
	}

	dict := make(map[string]bool, len(wordDict))
	length := make(map[int]bool, len(wordDict))

	for _, w := range wordDict {
		length[len(w)] = true
		dict[w] = true
	}

	sizes := make([]int, 0, len(length))
	for k := range length {
		sizes = append(sizes, k)
	}

	sort.Ints(sizes)

	// dp[i] == true，等于 wordBreak(s[:i+1], wordDict) == true
	dp := make([]bool, len(s)+1)
	dp[0] = true
	n := len(s)
	for i := 0; i <= n; i++ {
		if !dp[i] {
			continue
		}

		for _, size := range sizes {
			if i+size <= n {
				dp[i+size] = dp[i+size] || dict[s[i:i+size]]
			}
		}
	}

	return dp[n]
}
```



```python
class Solution(object):
    def wordBreak(self, s, wordDict):
        """
        :type s: str
        :type wordDict: List[str]
        :rtype: bool
        """
        def canBreak(s, m, wordDict):
            if s in m: return m[s]
            if s in wordDict: 
                m[s] = True
                return True
            
            for i in range(1, len(s)):
                r = s[i:]
                if r in wordDict and canBreak(s[0:i], m, wordDict):
                    m[s] = True
                    return True
            
            m[s] = False
            return False
            
        return canBreak(s, {}, set(wordDict))
```