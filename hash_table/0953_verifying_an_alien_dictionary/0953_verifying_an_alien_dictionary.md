953. Verifying an Alien Dictionary

Easy

In an alien language, surprisingly they also use english lowercase letters, but possibly in a different order. The order of the alphabet is some permutation of lowercase letters.

Given a sequence of words written in the alien language, and the order of the alphabet, return true if and only if the given words are sorted lexicographicaly in this alien language.

 

Example 1:

```
Input: words = ["hello","leetcode"], order = "hlabcdefgijkmnopqrstuvwxyz"
Output: true
Explanation: As 'h' comes before 'l' in this language, then the sequence is sorted.
```

Example 2:

```
Input: words = ["word","world","row"], order = "worldabcefghijkmnpqstuvxyz"
Output: false
Explanation: As 'd' comes after 'l' in this language, then words[0] > words[1], hence the sequence is unsorted.
```

Example 3:

```
Input: words = ["apple","app"], order = "abcdefghijklmnopqrstuvwxyz"
Output: false
Explanation: The first three characters "app" match, and the second string is shorter (in size.) According to lexicographical rules "apple" > "app", because 'l' > '∅', where '∅' is defined as the blank character which is less than any other character (More info).
```

Note:

1 <= words.length <= 100

1 <= words[i].length <= 20

order.length == 26

All characters in words[i] and order are english lowercase letters.


# 方法

判断字符串是否按照字典顺序排序的，先保存字典正确顺序为 map，然后遍历列表：  
当字符串长度相等时，依次判断字符顺序是否正确；  
当字符串长度不等时，则长的应为后序。


```go
func isAlienSorted(words []string, order string) bool {
    indexs := make([]int, 123)
	for i, c := range order {
		indexs[c] = i
	}

	less := func(i, j int) bool {
		si, sj := len(words[i]), len(words[j])
		for k := 0; k < si && k < sj; k++ {
			ii, ij := indexs[words[i][k]], indexs[words[j][k]]
			switch {
			case ii < ij:
				return true
			case ii > ij:
				return false
			}
		}
		return si <= sj
	}

	for i := 1; i < len(words); i++ {
		if !less(i-1, i) {
			return false
		}
	}
	return true
}
```


```python

class Solution(object):
    def isAlienSorted(self, words, order):
        """
        :type words: List[str]
        :type order: str
        :rtype: bool
        """
        lookup = {c: i for i, c in enumerate(order)}
        for i in xrange(len(words)-1):
            word1 = words[i]
            word2 = words[i+1]
            for k in xrange(min(len(word1), len(word2))):
                if word1[k] != word2[k]:
                    if lookup[word1[k]] > lookup[word2[k]]:
                        return False
                    break
            else:
                if len(word1) > len(word2):
                    return False
        return True
```
