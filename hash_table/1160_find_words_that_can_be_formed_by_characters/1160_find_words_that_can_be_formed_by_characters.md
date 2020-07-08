1160. Find Words That Can Be Formed by Characters


Easy


You are given an array of strings words and a string chars.

A string is good if it can be formed by characters from chars (each character can only be used once).

Return the sum of lengths of all good strings in words.

 

Example 1:

```
Input: words = ["cat","bt","hat","tree"], chars = "atach"
Output: 6
Explanation: 
The strings that can be formed are "cat" and "hat" so the answer is 3 + 3 = 6.
```

Example 2:

```
Input: words = ["hello","world","leetcode"], chars = "welldonehoneyr"
Output: 10
Explanation: 
The strings that can be formed are "hello" and "world" so the answer is 5 + 5 = 10.
```

Note:

1 <= words.length <= 1000  
1 <= words[i].length, chars.length <= 100  
All strings contain lowercase English letters only.


## 方法



```go
func countCharacters(words []string, chars string) int {
    magic := [26]int{}
	for _, c := range chars {
		magic[c-'a']++
	}
	good := func(word string, magic [26]int) bool {
		for _, c := range word {
			magic[c-'a']--
			if magic[c-'a'] < 0 {
				return false
			}
		}
		return true
	}
	
	goodWord := 0
	for _, word := range words {
		if good(word, magic) {
            goodWord += len(word)
		}
	}
	return goodWord
}
```


```python
class Solution:
    def countCharacters(self, words: List[str], chars: str) -> int:
        chars_cnt = collections.Counter(chars)
        ans = 0
        for word in words:
            word_cnt = collections.Counter(word)
            for c in word_cnt:
                if chars_cnt[c] < word_cnt[c]:
                    break
            else:
                ans += len(word)
        return ans
```