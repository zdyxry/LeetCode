844. Backspace String Compare


Easy


Given two strings S and T, return if they are equal when both are typed into empty text editors. # means a backspace character.

Example 1:
```
Input: S = "ab#c", T = "ad#c"
Output: true
Explanation: Both S and T become "ac".
```

Example 2:
```
Input: S = "ab##", T = "c#d#"
Output: true
Explanation: Both S and T become "".
```

Example 3:
```
Input: S = "a##c", T = "#a#c"
Output: true
Explanation: Both S and T become "c".
```

Example 4:
```
Input: S = "a#c", T = "b"
Output: false
Explanation: S becomes "c" while T becomes "b".
```


Note:

1 <= S.length <= 200

1 <= T.length <= 200

S and T only contain lowercase letters and '#' characters.

Follow up:

Can you solve it in O(N) time and O(1) space?


## 方法

依次遍历两个字符串，若为 # 号，则删除最后一个字符。




```go
func backspaceCompare(S string, T string) bool {
	s := make([]rune, 0)
	for _, c := range S {
		if c == '#' {
			if len(s) > 0 {
				s = s[:len(s)-1]
			}
		} else {
			s = append(s, c)
		}
	}
	s2 := make([]rune, 0)
	for _, c := range T {
		if c == '#' {
			if len(s2) > 0 {
				s2 = s2[:len(s2)-1]
			}
		} else {
			s2 = append(s2, c)
		}
	}
	return string(s) == string(s2)
}
```

```go
func backspaceCompare(S string, T string) bool {
    i := len(S)
	j := len(T)

	for i >= 0 || j >= 0 {
		i = nextIndex(&S, i)
		j = nextIndex(&T, j)

		if i >= 0 && j >= 0 && S[i] != T[j] {
			return false
		}

	}

	return i == j
}

// 返回 s[:i] 中，不是 '#' 的字符的最大的索引号
func nextIndex(s *string, i int) int {
	i--
	count := 0
	for i >= 0 && ((*s)[i] == '#' || count > 0) {
        if (*s)[i] == '#' {
			count++
		} else {
			count--
		}
		i--
	}
	return i
}
```


```python
class Solution(object):
    def backspaceCompare(self, S, T):
        """
        :type S: str
        :type T: str
        :rtype: bool
        """
        Sstk= []
        Tstk= []
        for i in S:
            if i!='#':
                Sstk.append(i)
            elif Sstk:
                Sstk.pop()
                
                
        for j in T:
            if j!='#':
                Tstk.append(j)
            elif Tstk:
                Tstk.pop()
                
        return Sstk== Tstk

```