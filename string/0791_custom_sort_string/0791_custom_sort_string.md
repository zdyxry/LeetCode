791. Custom Sort String


Medium


S and T are strings composed of lowercase letters. In S, no letter occurs more than once.

S was sorted in some custom order previously. We want to permute the characters of T so that they match the order that S was sorted. More specifically, if x occurs before y in S, then x should occur before y in the returned string.

Return any permutation of T (as a string) that satisfies this property.

Example :
```
Input: 
S = "cba"
T = "abcd"
Output: "cbad"
Explanation: 
"a", "b", "c" appear in S, so the order of "a", "b", "c" should be "c", "b", and "a". 
Since "d" does not appear in S, it can be at any position in T. "dcba", "cdba", "cbda" are also valid outputs.
```

Note:

S has length at most 26, and no character is repeated in S.  
T has length at most 200.  
S and T consist of lowercase letters only.  


## 方法


```go
func customSortString(S string, T string) string {
    res := ""
	for i := 0; i < len(S); i++ {
		count := strings.Count(T, S[i:i+1])
		res += strings.Repeat(S[i:i+1], count)
		T = strings.Replace(T, S[i:i+1], "", -1)
	}

	return res + T
}
```


```python
class Solution(object):
    def customSortString(self, S, T):
        """
        :type S: str
        :type T: str
        :rtype: str
        """
        l = []
        for i in S:
                l.append(i*T.count(i))
        for i in T:
            if i not in S:
                l.append(i)
        return ''.join(l)
```


```python
class Solution(object):
    def customSortString(self, S, T):
        """
        :type S: str
        :type T: str
        :rtype: str
        """
        counter, s = collections.Counter(T), set(S)
        result = [c*counter[c] for c in S]
        result.extend([c*counter for c, counter in counter.iteritems() if c not in s])
        return "".join(result)
```