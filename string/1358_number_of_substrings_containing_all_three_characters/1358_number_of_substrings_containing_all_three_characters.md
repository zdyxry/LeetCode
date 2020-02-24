1358. Number of Substrings Containing All Three Characters


Medium


Given a string s consisting only of characters a, b and c.

Return the number of substrings containing at least one occurrence of all these characters a, b and c.

 

Example 1:

```
Input: s = "abcabc"
Output: 10
Explanation: The substrings containing at least one occurrence of the characters a, b and c are "abc", "abca", "abcab", "abcabc", "bca", "bcab", "bcabc", "cab", "cabc" and "abc" (again). 
```

Example 2:

```
Input: s = "aaacb"
Output: 3
Explanation: The substrings containing at least one occurrence of the characters a, b and c are "aaacb", "aacb" and "acb". 
```

Example 3:

```
Input: s = "abc"
Output: 1
```
 

Constraints:

3 <= s.length <= 5 x 10^4  
s only consists of a, b or c characters.

## 方法


```go
func numberOfSubstrings(s string) int {
	ans := 0
	lo := -1
	cnt := []int{0, 0, 0}

	for hi, c := range s {
		cnt[c-'a']++
		for {
			if zeroArray(cnt) {
				break
			} else {
				ans += len(s) - hi
				lo++
				cnt[s[lo]-'a']--
			}
		}
	}
	return ans
}

func zeroArray(nums []int) bool {
	for _, c := range nums {
		if c == 0 {
			return true
		}
	}
	return false
}
```





```python
class Solution(object):
    def numberOfSubstrings(self, s):
        """
        :type s: str
        :rtype: int
        """
        res = i = 0
        count = {c: 0 for c in 'abc'}
        for j in xrange(len(s)):
            count[s[j]] += 1
            while all(count.values()):
                count[s[i]] -= 1
                i += 1
            res += i
        return res
```


```python
class Solution(object):
    def numberOfSubstrings(self, s):
        """
        :type s: str
        :rtype: int
        """
        ans, lo, cnt = 0, -1, {c : 0 for c in 'abc'}
        for hi, c in enumerate(s):
            cnt[c] += 1
            while all(cnt.values()):
                ans += len(s) - hi
                lo += 1
                cnt[s[lo]] -= 1
        return ans
```