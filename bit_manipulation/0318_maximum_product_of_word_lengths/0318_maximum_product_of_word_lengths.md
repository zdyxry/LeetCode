318. Maximum Product of Word Lengths


Medium


Given a string array words, find the maximum value of length(word[i]) * length(word[j]) where the two words do not share common letters. You may assume that each word will contain only lower case letters. If no such two words exist, return 0.

Example 1:

```
Input: ["abcw","baz","foo","bar","xtfn","abcdef"]
Output: 16 
Explanation: The two words can be "abcw", "xtfn".
```


Example 2:

```
Input: ["a","ab","abc","d","cd","bcd","abcd"]
Output: 4 
Explanation: The two words can be "ab", "cd".
```

Example 3:
```
Input: ["a","aa","aaa","aaaa"]
Output: 0 
Explanation: No such pair of words.
```


## 方法


```go
func maxProduct(words []string) int {
    size := len(words)

	masks := make([]int, size)
	for i := 0; i < size; i++ {
		for _, b := range words[i] {
			// 利用 bite 位来描述 words[i]
			masks[i] |= (1 << uint32(b-'a'))
		}
	}

	var max int
	for i := 0; i < size-1; i++ {
		for j := i + 1; j < size; j++ {
			// 当 words[i] 和 words[j] 拥有相同的字母时
			// 这个字母所对应的位进行 & 运算后，为 1
			if masks[i]&masks[j] != 0 {
				continue
			}
			temp := len(words[i]) * len(words[j])
			if max < temp {
				max = temp
			}
		}
	}

	return max
}
```

```python
class Solution(object):
    def maxProduct(self, words):
        """
        :type words: List[str]
        :rtype: int
        """
        d = {}
        for w in words:
            mask = 0
            for c in set(w):
                mask |= (1 << (ord(c) - 97))
            d[mask] = max(d.get(mask, 0), len(w))
        return max([d[x] * d[y] for x in d for y in d if not x & y] or [0])
```