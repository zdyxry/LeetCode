678. Valid Parenthesis String


Medium


Given a string containing only three types of characters: '(', ')' and '*', write a function to check whether this string is valid. We define the validity of a string by these rules:

Any left parenthesis '(' must have a corresponding right parenthesis ')'.  
Any right parenthesis ')' must have a corresponding left parenthesis '('.  
Left parenthesis '(' must go before the corresponding right parenthesis ')'.  
'*' could be treated as a single right parenthesis ')' or a single left parenthesis '(' or an empty string.  
An empty string is also valid.  

Example 1:
```
Input: "()"
Output: True
```

Example 2:
```
Input: "(*)"
Output: True
```

Example 3:
```
Input: "(*))"
Output: True
```

Note:  
The string size will be in the range [1, 100].  

## 方法


```go
func checkValidString(s string) bool {
    l, r := 0, 0
	n := len(s)
	for i := 0; i < n; i++ {
		// 从左侧开始的扫描
		// 并认为所有的 '*' 是 '('
		if s[i] == ')' {
			l--
		} else {
			l++
		}

		// 从右侧开始的扫描
		// 并认为所有的 '*' 是 ')'
		j := n - i - 1
		if s[j] == '(' {
			r--
		} else {
			r++
		}

		if l < 0 || r < 0 {
			// l < 0 意味着
			// 就算所有的 '*' 变成的 '('
			// s[:i] 中也没有能与 s[i] == ')' 匹配的 '('
			// r < 0 意味着
			// 就算所有的 '*' 变成的 ')'
			// s[j+1:] 中也没有能与 s[j] == '(' 匹配的 ')'
			return false
		}
	}

	return true
}
```



```python
class Solution(object):
    def checkValidString(self, s):
        """
        :type s: str
        :rtype: bool
        """
        lower, upper = 0, 0  # keep lower bound and upper bound of '(' counts
        for c in s:
            # 如果是 （，则 lower += 1， 否则 -=1
            lower += 1 if c == '(' else -1
            # 如果不是 ），则upper += 1，否则-=1
            upper -= 1 if c == ')' else -1
            # 如果 upper <0 ，说明哪怕所有的 （ 和 * 都算上，也无法抵消 ） 数量
            if upper < 0: break
            # 此时如果 lower < 0，那么说明 ）和 * 的数量比 （ 多，因为 * 可以作为空，所以不能说明此时无效
            lower = max(lower, 0)
        return lower == 0  # range of '(' count is
```


```python
class Solution(object):
    def checkValidString(self, s):
        """
        :type s: str
        :rtype: bool
        """
        vset = set([0])
        for c in s:
            nset = set()
            if c == '*':
                for v in vset:
                    nset.add(v + 1)
                    nset.add(v)
                    if v >= 1: nset.add(v - 1)
            elif c == '(':
                for v in vset:
                    nset.add(v + 1)
            elif c == ')':
                for v in vset:
                    if v >= 1: nset.add(v - 1)
            vset = nset
            print vset
        return 0 in vset
```