859. Buddy Strings


Easy


Given two strings A and B of lowercase letters, return true if and only if we can swap two letters in A so that the result equals B.

 

Example 1:
```
Input: A = "ab", B = "ba"
Output: true
```

Example 2:
```
Input: A = "ab", B = "ab"
Output: false
```
Example 3:
```
Input: A = "aa", B = "aa"
Output: true
```
Example 4:
```
Input: A = "aaaaaaabc", B = "aaaaaaacb"
Output: true
```
Example 5:
```
Input: A = "", B = "aa"
Output: false
```


Note:

0 <= A.length <= 20000

0 <= B.length <= 20000

A and B consist only of lowercase letters.


## 方法

将相同字符和不同字符分开考虑，最多允许有2个不同的字符，判断两个字符串中不同的字符是否相同。

考虑两个字符串相同时场景，判断是否有重复字母可以用于交换。



```go
func buddyStrings(A string, B string) bool {
    if len(A) != len(B) {
		return false
	}

	if A == B {
		return hasDouble(A)
	}

	size := len(A)
	i := 0
	countDown := 2
	ca, cb := byte(0), byte(0)
	for countDown > 0 && i < size {
		if A[i] != B[i] {
			ca += A[i]
			cb += B[i]
			countDown--
		}
		i++
	}

	return ca == cb && A[i:] == B[i:]
}

func hasDouble(s string) bool {
	seen := [26]bool{}
	for i := range s {
		b := s[i] - 'a'
		if seen[b] {
			return true
		}
		seen[b] = true
	}
	return false
}

```



```python

class Solution(object):
    def buddyStrings(self, A, B):
        """
        :type A: str
        :type B: str
        :rtype: bool
        """
        if len(A) != len(B):
            return False
        A = list(A)
        diff = []
        for i in range(len(A)):
            if A[i] != B[i]:
                diff.append(i)
        if len(set(A)) < len(A) and len(diff) == 0:
            return True
        return len(diff) == 2 and A[diff[0]] == B[diff[1]] and A[diff[1]] == B[diff[0]]
```