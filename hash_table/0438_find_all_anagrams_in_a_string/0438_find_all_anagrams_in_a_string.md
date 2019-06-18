438. Find All Anagrams in a String

Easy

Given a string s and a non-empty string p, find all the start indices of p's anagrams in s.

Strings consists of lowercase English letters only and the length of both strings s and p will not be larger than 20,100.

The order of output does not matter.

Example 1:

```
Input:
s: "cbaebabacd" p: "abc"

Output:
[0, 6]

Explanation:
The substring with start index = 0 is "cba", which is an anagram of "abc".
The substring with start index = 6 is "bac", which is an anagram of "abc".
```

Example 2:

```
Input:
s: "abab" p: "ab"

Output:
[0, 1, 2]

Explanation:
The substring with start index = 0 is "ab", which is an anagram of "ab".
The substring with start index = 1 is "ba", which is an anagram of "ab".
The substring with start index = 2 is "ab", which is an anagram of "ab".
```


# 方法
统计 p 中字符出现次数，遍历 s ，当长度等于 p 时，判断当前字符出现次数是否相当，若相当，则添加当前字符中起始索引。


```go
func findAnagrams(s string, p string) []int {
    var res = []int{}

	var target, window [26]int
	for i := 0; i < len(p); i++ {
		target[p[i]-'a']++
	}

	check := func(i int) {
		if window == target {
			res = append(res, i)
		}
	}

	for i := 0; i < len(s); i++ {
		window[s[i]-'a']++
		if i == len(p)-1 {
			check(0)
		} else if len(p) <= i {
			window[s[i-len(p)]-'a']--
			check(i - len(p) + 1)
		}
	}

	return res
}
```


```python
class Solution(object):
    def findAnagrams(self, s, p):
        """
        :type s: str
        :type p: str
        :rtype: List[int]
        """
        l1=len(s)
        l2=len(p)
        rst=[]
        d1=[0]*26
        for c in p:
            d1[ord(c)-ord('a')]+=1
        #
        d2=[0]*26
        for c in s[:l2]:
            d2[ord(c)-ord('a')]+=1
        if d1==d2:
            rst.append(0)
        #
        for i in range(1,l1-l2+1):
            d2[ord(s[i-1])-ord('a')]-=1
            d2[ord(s[i+l2-1])-ord('a')]+=1
            if d1==d2:
                rst.append(i)
        return rst

```