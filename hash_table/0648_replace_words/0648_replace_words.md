648. Replace Words


Medium


In English, we have a concept called root, which can be followed by some other words to form another longer word - let's call this word successor. For example, the root an, followed by other, which can form another word another.

Now, given a dictionary consisting of many roots and a sentence. You need to replace all the successor in the sentence with the root forming it. If a successor has many roots can form it, replace it with the root with the shortest length.

You need to output the sentence after the replacement.

Example 1:

```
Input: dict = ["cat", "bat", "rat"]
sentence = "the cattle was rattled by the battery"
Output: "the cat was rat by the bat"
```

Note:

The input will only have lower-case letters.

1 <= dict words number <= 1000

1 <= sentence words number <= 1000

1 <= root length <= 100

1 <= sentence words length <= 1000


## 方法

```go
func replaceWords(dict []string, sentence string) string {
    hasRoot := make(map[string]bool, len(dict))
	hasLen := make([]bool, 101)
	for i := range dict {
		hasRoot[dict[i]] = true
		hasLen[len(dict[i])] = true
	}

	// ls 收集了所有 root 的长度
	ls := make([]int, 0, 101)
	for i, ok := range hasLen {
		if ok {
			ls = append(ls, i)
		}
	}
	sort.Ints(ls)

	words := strings.Split(sentence, " ")
	for i, w := range words {
		for _, l := range ls {
			// w 的前 l 个字符是字根就修改 w 并结束 for 循环
			if l < len(w) && hasRoot[w[:l]] {
				words[i] = w[:l]
				break
			}
		}
	}

	return strings.Join(words, " ")
}
```


```python
class Solution(object):
    def replaceWords(self, dict, sentence):
        """
        :type dict: List[str]
        :type sentence: str
        :rtype: str
        """
        rootset = set(dict)
        
        def replace(word):
            for i in xrange(1, len(word)):
                if word[:i] in rootset:
                    return word[:i]
            return word
        return " ".join(map(replace, sentence.split()))
```


```python
class Solution(object):
    def replaceWords(self, dict, sentence):
        """
        :type dict: List[str]
        :type sentence: str
        :rtype: str
        """
        _trie = lambda: collections.defaultdict(_trie)
        trie = _trie()
        END = True
        for root in dict:
            cur = trie
            for letter in root:
                cur = cur[letter]
            cur[END] = root

        def replace(word):
            cur = trie
            for letter in word:
                if letter not in cur: break
                cur = cur[letter]
                if END in cur:
                    return cur[END]
            return word

        return " ".join(map(replace, sentence.split()))
```