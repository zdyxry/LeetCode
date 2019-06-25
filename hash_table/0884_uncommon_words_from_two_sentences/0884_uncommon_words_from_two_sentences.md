884. Uncommon Words from Two Sentences

Easy

We are given two sentences A and B.  (A sentence is a string of space separated words.  Each word consists only of lowercase letters.)

A word is uncommon if it appears exactly once in one of the sentences, and does not appear in the other sentence.

Return a list of all uncommon words. 

You may return the list in any order.

 

Example 1:

```
Input: A = "this apple is sweet", B = "this apple is sour"
Output: ["sweet","sour"]
```

Example 2:

```
Input: A = "apple apple", B = "banana"
Output: ["banana"]
```
 

Note:

0 <= A.length <= 200

0 <= B.length <= 200

A and B both contain only spaces and lowercase letters.

# 方法
只需要求解出现1次的单词并记录。

python 可以用 collections 简化实现。

```go
func uncommonFromSentences(A string, B string) []string {
    words := map[string]bool{}
	ret := []string{}

	var help = func (s string) {
		for _, word := range strings.Split(s, " ") {
			_, ok := words[word]
			words[word] = !ok
		}
	}

	help(A)
	help(B)

	for key, val := range(words) {
		if val {
			ret = append(ret, key)
		}
	}

	return ret
}

```

```python
class Solution(object):
    def uncommonFromSentences(self, A, B):
        """
        :type A: str
        :type B: str
        :rtype: List[str]
        """
        res = []
        d = {}
        for s in A.split(" ") + B.split(" "):
            if s in d:
                d[s] += 1
            else:
                d[s] = 1
        
        for k,v in d.items():
            if v == 1:
                res.append(k)

        return res
```



```python
class Solution(object):
    def uncommonFromSentences(self, A, B):
        """
        :type A: str
        :type B: str
        :rtype: List[str]
        """
        c = collections.Counter((A + " " + B).split())
        return [w for w in c if c[w] == 1]
```