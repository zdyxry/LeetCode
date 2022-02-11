2062. Count Vowel Substrings of a String


Easy


A substring is a contiguous (non-empty) sequence of characters within a string.

A vowel substring is a substring that only consists of vowels ('a', 'e', 'i', 'o', and 'u') and has all five vowels present in it.

Given a string word, return the number of vowel substrings in word.

 

Example 1:

```
Input: word = "aeiouu"
Output: 2
Explanation: The vowel substrings of word are as follows (underlined):
- "aeiouu"
- "aeiouu"
```

Example 2:

```
Input: word = "unicornarihan"
Output: 0
Explanation: Not all 5 vowels are present, so there are no vowel substrings.
```

Example 3:

```
Input: word = "cuaieuouac"
Output: 7
Explanation: The vowel substrings of word are as follows (underlined):
- "cuaieuouac"
- "cuaieuouac"
- "cuaieuouac"
- "cuaieuouac"
- "cuaieuouac"
- "cuaieuouac"
- "cuaieuouac"
```
 

Constraints:

1 <= word.length <= 100   
word consists of lowercase English letters only.


## 方法


```
func countVowelSubstrings(word string) int {
	var res int
	var cnt [26]int
	n := len(word)
	for i := 0; i < n; i++ {
		cnt = [26]int{}
		for j := i; j < n; j++ {
			if word[j] != 'a' && word[j] != 'e' && word[j] != 'i' && word[j] != 'o' && word[j] != 'u' {
				break
			}
			cnt[word[j]-'a']++
			if cnt['a'-'a'] > 0 && cnt['e'-'a'] > 0 && cnt['i'-'a'] > 0 && cnt['o'-'a'] > 0 && cnt['u'-'a'] > 0 {
				res++
			}
		}
	}
	return res
}

```


```
class Solution:
    def countVowelSubstrings(self, word: str) -> int:
        a = {'a','e','i','o','u'}
        l, r = 0, 0
        cnt = 0
        while l < len(word):
            if word[r] in a:
                if len(set(word[l:r+1])) >= 5:
                    cnt += 1
                r += 1
                if r == len(word):
                    l += 1
                    r = l
            else:
                l += 1
                r = l
        return cnt
```