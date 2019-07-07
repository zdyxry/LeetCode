557. Reverse Words in a String III
    

Easy


Given a string, you need to reverse the order of characters in each word within a sentence while still preserving whitespace and initial word order.

Example 1:
```
Input: "Let's take LeetCode contest"
Output: "s'teL ekat edoCteeL tsetnoc"
```

Note: In the string, each word is separated by single space and there will not be any extra space in the string.

## 方法

遍历字符串，将其以 `" "` 切分，并将切分后字母反转保存。



```go
func reverseWords(s string) string {
	ss := strings.Split(s, " ")
	for i, s := range ss {
		ss[i] = revers(s)
	}
	return strings.Join(ss, " ")
}

func revers(s string) string {
	bytes := []byte(s)
	i, j := 0, len(bytes)-1
	for i < j {
		bytes[i], bytes[j] = bytes[j], bytes[i]
		i++
		j--
	}
	return string(bytes)
}
```



```python
class Solution(object):
    def reverseWords(self, s):
        """
        :type s: str
        :rtype: str
        """
        return " ".join(i[::-1] for i in s.split())

```

