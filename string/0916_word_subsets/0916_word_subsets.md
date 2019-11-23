916. Word Subsets


Medium


We are given two arrays A and B of words.  Each word is a string of lowercase letters.

Now, say that word b is a subset of word a if every letter in b occurs in a, including multiplicity.  For example, "wrr" is a subset of "warrior", but is not a subset of "world".

Now say a word a from A is universal if for every b in B, b is a subset of a. 

Return a list of all universal words in A.  You can return the words in any order.

 

Example 1:

```
Input: A = ["amazon","apple","facebook","google","leetcode"], B = ["e","o"]
Output: ["facebook","google","leetcode"]
```

Example 2:

```
Input: A = ["amazon","apple","facebook","google","leetcode"], B = ["l","e"]
Output: ["apple","google","leetcode"]
```

Example 3:

```
Input: A = ["amazon","apple","facebook","google","leetcode"], B = ["e","oo"]
Output: ["facebook","google"]
```

Example 4:

```
Input: A = ["amazon","apple","facebook","google","leetcode"], B = ["lo","eo"]
Output: ["google","leetcode"]
```

Example 5:

```
Input: A = ["amazon","apple","facebook","google","leetcode"], B = ["ec","oc","ceo"]
Output: ["facebook","leetcode"]
```
 

Note:

1 <= A.length, B.length <= 10000  
1 <= A[i].length, B[i].length <= 10  
A[i] and B[i] consist only of lowercase letters.  
All words in A[i] are unique: there isn't i != j with A[i] == A[j].  

## 方法


```go
func wordSubsets(A []string, B []string) []string {
    clt := new([26]int)
	for _, b := range B {
		collect(clt, count(b))
	}

	res := make([]string, 0, len(A))
	for _, a := range A {
		if isSubset(count(a), clt) {
			res = append(res, a)
		}
	}

	return res
}

func count(s string) *[26]int {
	res := [26]int{}
	for _, b := range s {
		res[b-'a']++
	}
	return &res
}

func isSubset(s, clt *[26]int) bool {
	isSubset := true
	i := 0
	for isSubset && i < 26 {
		isSubset = isSubset && (s[i] >= clt[i])
		i++
	}
	return isSubset
}

// collect to clt
func collect(clt, b *[26]int) {
	for i := range clt {
		clt[i] = max(clt[i], b[i])
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```



```python
class Solution(object):
    def wordSubsets(self, A, B):
        """
        :type A: List[str]
        :type B: List[str]
        :rtype: List[str]
        """
        s = set(A)
        letters_required = {}
        for i in B:
            for j in i:
                count = i.count(j)
                if j not in letters_required or count > letters_required[j]:
                    letters_required[j] = count

        for i in A:
            for j in letters_required:
                if i.count(j) < letters_required[j]:
                    s.remove(i)
                    break
        return list(s)
```