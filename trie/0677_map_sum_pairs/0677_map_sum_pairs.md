677. Map Sum Pairs


Medium


Implement a MapSum class with insert, and sum methods.

For the method insert, you'll be given a pair of (string, integer). The string represents the key and the integer represents the value. If the key already existed, then the original key-value pair will be overridden to the new one.

For the method sum, you'll be given a string representing the prefix, and you need to return the sum of all the pairs' value whose key starts with the prefix.

Example 1:

```
Input: insert("apple", 3), Output: Null
Input: sum("ap"), Output: 3
Input: insert("app", 2), Output: Null
Input: sum("ap"), Output: 5
```

## 方法

```go
type MapSum struct {
	children       map[rune]*MapSum
	value          int
}

/** Initialize your data structure here. */
func Constructor() MapSum {
	return MapSum{
		children:       make(map[rune]*MapSum),
		value:          0,
	}
}

/** Inserts a word into the trie. */
func (t *MapSum) Insert(word string, value int) {
	curr := t
	for _, ch := range word {
		if _, ok := curr.children[rune(ch)]; !ok {
			trie := Constructor()
			curr.children[rune(ch)] = &trie
		}
		curr = curr.children[rune(ch)]
	}
	curr.value = value
}

func (t *MapSum) Sum(prefix string) int {
	curr := t
	for _, ch := range prefix {
		if _, ok := curr.children[rune(ch)]; !ok {
			return 0
		}
		curr = curr.children[rune(ch)]
	}

	return SumChildren(*curr)
}

func SumChildren(t MapSum) int {
	curr := t
	sum := t.value
	for k := range curr.children {
		sum = sum + SumChildren(*curr.children[k])
	}

	return sum
}
```

```python
class MapSum(object):
    def __init__(self):
        """
        Initialize your data structure here.
        """
        self.d = {}

    def insert(self, key, val):
        """
        :type key: str
        :type val: int
        :rtype: None
        """
        t = self.d
        for c in key:
            if c not in t:
                t[c] = {}
            t = t[c]
        t['val'] = val
    
    def dfs(self, t, sum):
        for c in t:
            if c == 'val':
                sum += t[c] #前缀求和
            else:
                sum = self.dfs(t[c], sum)   #深搜
        return sum
    
    def sum(self, prefix):
        t = self.d
        for c in prefix:
            if c not in t:
                return 0        #判断前缀是否存在
            t = t[c]
        ans = self.dfs(t, 0)
        return ans
```



```python
677. Map Sum Pairs
Medium

401

71

Add to List

Share
Implement a MapSum class with insert, and sum methods.

For the method insert, you'll be given a pair of (string, integer). The string represents the key and the integer represents the value. If the key already existed, then the original key-value pair will be overridden to the new one.

For the method sum, you'll be given a string representing the prefix, and you need to return the sum of all the pairs' value whose key starts with the prefix.

Example 1:
Input: insert("apple", 3), Output: Null
Input: sum("ap"), Output: 3
Input: insert("app", 2), Output: Null
Input: sum("ap"), Output: 5
```