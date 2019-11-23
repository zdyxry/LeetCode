890. Find and Replace Pattern


Medium


You have a list of words and a pattern, and you want to know which words in words matches the pattern.

A word matches the pattern if there exists a permutation of letters p so that after replacing every letter x in the pattern with p(x), we get the desired word.

(Recall that a permutation of letters is a bijection from letters to letters: every letter maps to another letter, and no two letters map to the same letter.)

Return a list of the words in words that match the given pattern. 

You may return the answer in any order.

 

Example 1:

```
Input: words = ["abc","deq","mee","aqq","dkd","ccc"], pattern = "abb"
Output: ["mee","aqq"]
Explanation: "mee" matches the pattern because there is a permutation {a -> m, b -> e, ...}. 
"ccc" does not match the pattern because {a -> c, b -> c, ...} is not a permutation,
since a and b map to the same letter.
```

Note:

1 <= words.length <= 50  
1 <= pattern.length = words[i].length <= 20  

## 方法

```go
func findAndReplacePattern(words []string, pattern string) []string {
    np := normalize(pattern)
	res := make([]string, 0, len(words))
	for _, w := range words {
		if normalize(w) == np {
			res = append(res, w)
		}
	}
	return res
}

// 把 str 标准化
func normalize(str string) string {
	m := make(map[rune]byte, len(str))
	/**按照字符在 s 中第一次出现的顺序，依次映射到 a,b,c,... */
	for _, c := range str {
		if _, ok := m[c]; !ok {
			m[c] = byte(len(m)) + 'a'
		}
	}
	/**按照字符的映射关系，映射 s 到 res */
	res := make([]byte, len(str))
	for i, c := range str {
		res[i] = m[c]
	}
	return string(res)
}
```


```python
class Solution(object):
    def findAndReplacePattern(self, words, pattern):
        """
        :type words: List[str]
        :type pattern: str
        :rtype: List[str]
        """
        res = []

        def match(word, pattern):
            if len(word) != len(pattern):
                return False
            dir = {}
            for w, p in zip(word, pattern):
                if w not in dir:
                    if p in dir.values():
                        return False
                    dir[w] = p
                else:
                    if dir[w] != p:
                        return False
            return True

        for w in words:
            if match(w, pattern):
                res.append(w)
        return res

```