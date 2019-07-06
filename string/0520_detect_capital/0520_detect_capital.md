520. Detect Capital


Easy


Given a word, you need to judge whether the usage of capitals in it is right or not.

We define the usage of capitals in a word to be right when one of the following cases holds:

All letters in this word are capitals, like "USA".
All letters in this word are not capitals, like "leetcode".
Only the first letter in this word is capital, like "Google".
Otherwise, we define that this word doesn't use capitals in a right way.
 

Example 1:
```
Input: "USA"
Output: True
```

Example 2:
```
Input: "FlaG"
Output: False
```

Note: The input will be a non-empty word consisting of uppercase and lowercase latin letters.


## 方法

遍历字符串，统计其大写字母个数，根据需求判断。



```go
func detectCapitalUse(word string) bool {
    head := word[:1]
	tail := word[1:]
	if isIn(head, 'A', 'Z') {
		return isIn(tail, 'A', 'Z') || isIn(tail, 'a', 'z')
	}
	return isIn(tail, 'a', 'z')
}

func isIn(s string, first, last byte) bool {
	for i := range s {
		if !(first <= s[i] && s[i] <= last) {
			return false
		}
	}
	return true
}
```


```python
class Solution(object):
    def detectCapitalUse(self, word):
        """
        :type word: str
        :rtype: bool
        """
        count = 0
        for i in range(len(word)):
            if word[i].isupper():
                count+=1
            
        if count == 1:
            if word[0].isupper():
                return True
            else:
                return False
        if count == len(word):
            return True
        elif count == 0:
            return True
        else:
            return False
```