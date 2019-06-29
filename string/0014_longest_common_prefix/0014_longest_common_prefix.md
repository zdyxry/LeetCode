14. Longest Common Prefix

Easy

Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

Example 1:

```
Input: ["flower","flow","flight"]
Output: "fl"
```
Example 2:

```
Input: ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.
```

Note:

All given inputs are in lowercase letters a-z.



## 方法
因为要找到所有单词中公共部分最长的，所以以第一个单词为基准就好（也可以找到最短的单词做基准），遍历后续所有单词，依次比较字母是否相同。

```go
func longestCommonPrefix(strs []string) string {
	short := shortest(strs)

	for i, r := range short {
		for j := 0; j < len(strs); j++ {
			if strs[j][i] != byte(r) {
				return strs[j][:i]
			}
		}
	}

	return short
}

func shortest(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	res := strs[0]
	for _, s := range strs {
		if len(res) > len(s) {
			res = s
		}
	}

	return res
}

```


```python

class Solution(object):
    def longestCommonPrefix(self, strs):
        """
        :type strs: List[str]
        :rtype: str
        """
        if not strs:
            return ""

        for i in xrange(len(strs[0])):
            for string in strs[1:]:
                if i >= len(string) or string[i] != strs[0][i]:
                    return strs[0][:i]
        return strs[0]
```