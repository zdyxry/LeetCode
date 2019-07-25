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


```python
# https://leetcode.com/problems/backspace-string-compare/discuss/145786/Python-tm
# 对两个字符串从右往左遍历，用子方程getChar() 从两个字符串分别取值，如取值不等则返回False，相等就继续迭代，知道迭代技术返回True

# 子方程getChar()的取值规则：

# While循环退出条件：下标出界或者返回值不为空。
# Case 1: 当值等于`#`, `count`增值
# Case 2: 如果`count == 0` 说明这个值没有被`#`给抵消，返回
# Case 3: 如果`count != 0` 且这个值不为`#`，这个值要被`#`抵消掉
# 三个Case运行完后，记得锁紧下标`r`
class Solution(object):
    def backspaceCompare(self, S1, S2):
        r1 = len(S1) - 1 
        r2 = len (S2) - 1
        
        while r1 >= 0 or r2 >= 0:
            char1 = char2 = ""
            if r1 >= 0:
                char1, r1 = self.getChar(S1, r1)
            if r2 >= 0:
                char2, r2 = self.getChar(S2, r2)
            if char1 != char2:
                return False
        return True
        
    
    def getChar(self, s , r):
        char, count = '', 0
        while r >= 0 and not char:
            if s[r] == '#':
                count += 1
            elif count == 0:
                char = s[r]
            else:
                count -= 1
            r -= 1
        return char, r
```