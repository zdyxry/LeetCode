1408. String Matching in an Array


Easy


Given an array of string words. Return all strings in words which is substring of another word in any order. 

String words[i] is substring of words[j], if can be obtained removing some characters to left and/or right side of words[j].

 

Example 1:

```
Input: words = ["mass","as","hero","superhero"]
Output: ["as","hero"]
Explanation: "as" is substring of "mass" and "hero" is substring of "superhero".
["hero","as"] is also a valid answer.
```

Example 2:

```
Input: words = ["leetcode","et","code"]
Output: ["et","code"]
Explanation: "et", "code" are substring of "leetcode".
```

Example 3:

```
Input: words = ["blue","green","bu"]
Output: []
```
 

Constraints:

1 <= words.length <= 100  
1 <= words[i].length <= 30  
words[i] contains only lowercase English letters.  
It's guaranteed that words[i] will be unique.  


## 方法

```go
type ByLen []string

func (a ByLen) Len() int {
	return len(a)
}

func (a ByLen) Less(i, j int) bool {
	return len(a[i]) < len(a[j])
}

func (a ByLen) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func stringMatching(words []string) []string {
	sort.Sort(ByLen(words))
	res := []string{}
	for i := range words {
		for j := i + 1; j < len(words); j++ {
			if strings.Contains(words[j], words[i]) {
				res = append(res, words[i])
			}
		}
	}

	unique := []string{}
	for _, v := range res {
		skip := false
		for _, u := range unique {
			if v == u {
				skip = true
				break
			}
		}
		if !skip {
			unique = append(unique, v)
		}
	}
	return unique
}
```



```python
class Solution:
    def stringMatching(self, words: List[str]) -> List[str]:
        words.sort(key=len) #by size in ascending order 
        ans = []
        for i, word in enumerate(words):
            for j in range(i+1, len(words)):
                if word in words[j]: 
                    ans.append(word)
                    break
        return ans
```