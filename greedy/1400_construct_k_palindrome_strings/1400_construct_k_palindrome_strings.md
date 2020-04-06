1400. Construct K Palindrome Strings


Medium


Given a string s and an integer k. You should construct k non-empty palindrome strings using all the characters in s.

Return True if you can use all the characters in s to construct k palindrome strings or False otherwise.

 

Example 1:

```
Input: s = "annabelle", k = 2
Output: true
Explanation: You can construct two palindromes using all characters in s.
Some possible constructions "anna" + "elble", "anbna" + "elle", "anellena" + "b"
```

Example 2:

```
Input: s = "leetcode", k = 3
Output: false
Explanation: It is impossible to construct 3 palindromes using all the characters of s.
```

Example 3:

```
Input: s = "true", k = 4
Output: true
Explanation: The only possible solution is to put each character in a separate string.
```

Example 4:

```
Input: s = "yzyzyzyzyzyzyzy", k = 2
Output: true
Explanation: Simply you can put all z's in one string and all y's in the other string. Both strings will be palindrome.
```

Example 5:

```
Input: s = "cr", k = 7
Output: false
Explanation: We don't have enough characters in s to construct 7 palindromes.
```

Constraints:

1 <= s.length <= 10^5  
All characters in s are lower-case English letters.  
1 <= k <= 10^5

## 方法

```go
func canConstruct(s string, k int) bool {
    n := len(s)
	if n < k {
		return false
	}
	if n == k {
		return true
	}

	cnt := make([]int, 26)

	for i := 0; i < n; i++ {
		cnt[int(s[i]-'a')]++
	}

	var odd int

	for i := 0; i < 26; i++ {
		odd += cnt[i] & 1
	}

	return odd <= k
}
```



```python
class Solution:
    def canConstruct(self, s: str, k: int) -> bool:
        digit_count = {}
        if len(s) < k:
            return False
        elif len(s) == k:
            return True
        else:
            odd = 0
            for i in set(s):
                digit_count[i] = s.count(i)
            
            for i in digit_count.values():
                if i % 2 != 0:
                    odd += 1
            if odd > k:
                return False
            else:
                return True
```