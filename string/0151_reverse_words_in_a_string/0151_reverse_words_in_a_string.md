151. Reverse Words in a String


Medium


Given an input string, reverse the string word by word.

 

Example 1:

```
Input: "the sky is blue"
Output: "blue is sky the"
```

Example 2:

```
Input: "  hello world!  "
Output: "world! hello"

Explanation: Your reversed string should not contain leading or trailing spaces.
```

Example 3:

```
Input: "a good   example"
Output: "example good a"
Explanation: You need to reduce multiple spaces between two words to a single space in the reversed string.
```
 

Note:

A word is defined as a sequence of non-space characters.   
Input string may contain leading or trailing spaces. However, your reversed string should not contain leading or trailing spaces.   
You need to reduce multiple spaces between two words to a single space in the reversed string.   
 

Follow up:

For C programmers, try to solve it in-place in O(1) extra space.   


## 方法

```go
func reverseWords(s string) string {
    ss := strings.Fields(s)
	reverse(&ss, 0, len(ss)-1)
	return strings.Join(ss, " ")
}

func reverse(m *[]string, i int, j int) {
	for i <= j {
		(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
		i++
		j--
	}
}
```



```python
class Solution(object):
    def reverseWords(self, s):
        """
        :type s: str
        :rtype: str
        """
        return ' '.join(reversed(s.split()))
```