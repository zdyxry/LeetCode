290. Word Pattern

Easy

Given a pattern and a string str, find if str follows the same pattern.

Here follow means a full match, such that there is a bijection between a letter in pattern and a non-empty word in str.

Example 1:

```
Input: pattern = "abba", str = "dog cat cat dog"
Output: true
```

Example 2:

```
Input:pattern = "abba", str = "dog cat cat fish"
Output: false
```

Example 3:

```
Input: pattern = "aaaa", str = "dog cat cat dog"
Output: false
```

Example 4:

```
Input: pattern = "abba", str = "dog dog dog dog"
Output: false
```

Notes:

You may assume pattern contains only lowercase letters, and str contains lowercase letters that may be separated by a single space.


# 方法
拆分 pattern & str，分别作为map 的key 和 value。


```go
func wordPattern(pattern string, str string) bool {
	ps := strings.Split(pattern, "")
	ss := strings.Split(str, " ")

	if len(ps) != len(ss) {
		return false
	}

	return isMatch(ps, ss) && isMatch(ss, ps)
}

func isMatch(s1, s2 []string) bool {
	size := len(s1)

	m := make(map[string]string, size)

	var i int
	var w string
	var ok bool

	for i = 0; i < size; i++ {
		if w, ok = m[s1[i]]; ok {
			if w != s2[i] {
				return false
			}
		} else {
			m[s1[i]] = s2[i]
		}
	}
	return true
}

```

```python
class Solution(object):
    def wordPattern(self, pattern, str):
        """
        :type pattern: str
        :type str: str
        :rtype: bool
        """
        from itertools import izip 
        words = str.split()  # Space: O(n)
        if len(pattern) != len(words):
            return False

        w2p, p2w = {}, {}
        for p, w in izip(pattern, words):
            if w not in w2p and p not in p2w:
                # Build mapping. Space: O(c)
                w2p[w] = p
                p2w[p] = w
            elif w not in w2p or w2p[w] != p:
                # Contradict mapping.
                return False
        return True
```