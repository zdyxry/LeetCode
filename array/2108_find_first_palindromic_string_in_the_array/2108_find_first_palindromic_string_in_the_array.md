2108. Find First Palindromic String in the Array


Easy

Given an array of strings words, return the first palindromic string in the array. If there is no such string, return an empty string "".

A string is palindromic if it reads the same forward and backward.

 

Example 1:

```
Input: words = ["abc","car","ada","racecar","cool"]
Output: "ada"
Explanation: The first string that is palindromic is "ada".
Note that "racecar" is also palindromic, but it is not the first.
```

Example 2:

```
Input: words = ["notapalindrome","racecar"]
Output: "racecar"
Explanation: The first and only string that is palindromic is "racecar".
```

Example 3:

```
Input: words = ["def","ghi"]
Output: ""
Explanation: There are no palindromic strings, so the empty string is returned.
```

Constraints:

```
1 <= words.length <= 100
1 <= words[i].length <= 100
words[i] consists only of lowercase English letters.
```


## 方法


```
func firstPalindrome(words []string) string {
	for _, word := range words {
		if isPalindrome(word) {
			return word
		}
	}
	return ""
}
func isPalindrome(word string) bool {
	n := len(word)
	l, r := 0, n-1
	for l < r {
		if word[l] != word[r] {
			return false
		} else {
			l++
			r--
		}
	}
	return true
}

```


```
class Solution:
    def firstPalindrome(self, words: List[str]) -> str:
        def isPalindrome(word: str) -> bool:
            n = len(word)
            l, r = 0, n - 1
            while l < r:
                if word[l] != word[r]:
                    return False
                l += 1
                r -= 1
            return True
        
        # 顺序遍历字符串数组，如果遇到回文字符串则返回，未遇到则返回空字符串
        for word in words:
            if isPalindrome(word):
                return word
        return ""

```