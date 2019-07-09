696. Count Binary Substrings


Easy


Give a string s, count the number of non-empty (contiguous) substrings that have the same number of 0's and 1's, and all the 0's and all the 1's in these substrings are grouped consecutively.

Substrings that occur multiple times are counted the number of times they occur.

Example 1:
```
Input: "00110011"
Output: 6
Explanation: There are 6 substrings that have equal number of consecutive 1's and 0's: "0011", "01", "1100", "10", "0011", and "01".

Notice that some of these substrings repeat and are counted the number of times they occur.

Also, "00110011" is not a valid substring because all the 0's (and 1's) are not grouped together.
```



Example 2:
```
Input: "10101"
Output: 4
Explanation: There are 4 substrings: "10", "01", "10", "01" that have equal number of consecutive 1's and 0's.
```

Note:


s.length will be between 1 and 50,000.

s will only consist of "0" or "1" characters.


## 方法

统计当前字符连续出现的次数，当字符与前一个字符不同，则将当前字符数置为 1 ，最终结果为当前连续字符出现次数与前一个连续字符出现次数的最小值相加。


```
  001110010    prev = 0    cur = 0
 0 01110010    prev = 0    cur = 1
0 0 1110010    prev = 0    cur = 2
00 1 110010    prev = 2    cur = 1    01
001 1 10010    prev = 2    cur = 2    0011
0011 1 0010    prev = 2    cur = 3
00111 0 010    prev = 3    cur = 1    10
001110 0 10    prev = 3    cur = 2    1100
0011100 1 0    prev = 2    cur = 1    01
00111001 0     prev = 1    cur = 1    10
```



```go
func countBinarySubstrings(s string) int {
    count, countZero, countOne := 0, 0, 0
	prev := rune(s[0])

	for _, r := range s {
		if prev == r {
			if r == '0' {
				countZero++
			} else {
				countOne++
			}
		} else {
			count += min(countZero, countOne)
			if r == '0' {
				countZero = 1
			} else {
				countOne = 1
			}
		}
		prev = r
	}

	return count + min(countZero, countOne)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

```


```python
class Solution(object):
    def countBinarySubstrings(self, s):
        """
        :type s: str
        :rtype: int
        """
        result, prev, curr = 0, 0, 1
        for i in xrange(1, len(s)):
            if s[i-1] != s[i]:
                result += min(prev, curr)
                prev, curr = curr, 1
            else:
                curr += 1
        result += min(prev, curr)
        return result

```


```python
class Solution:
    def countBinarySubstrings(self, s):
        """
        :type s: str
        :rtype: int
        """
        result = 0
        prev_length = 0
        cur_length = 1
        for i in range(1, len(s)):
            if (s[i] == s[i - 1]):
                cur_length += 1
            else:
                prev_length = cur_length
                cur_length = 1
            if prev_length >= cur_length:
                result += 1
        return result
```