796. Rotate String


Easy


We are given two strings, A and B.

A shift on A consists of taking string A and moving the leftmost character to the rightmost position. For example, if A = 'abcde', then it will be 'bcdea' after one shift on A. Return True if and only if A can become B after some number of shifts on A.

Example 1:
```
Input: A = 'abcde', B = 'cdeab'
Output: true
```


Example 2:

```
Input: A = 'abcde', B = 'abced'
Output: false
```

Note:

A and B will have length at most 100.

## 方法

```go
func rotateString(A string, B string) bool {
    if A == B {
		return true
	}
	for i := 0; i < len(A); i++ {
		if B == string(A[i+1:])+string(A[0:i+1]) {
			return true
		}
	}
	return false
}
```



```python
class Solution:
    def rotateString(self, A: str, B: str) -> bool:
        if len(A) != len(B):
            return False
        if A == B:
            return True
        for i in range(1, len(A)):
            if A[i:] + A[:i] == B:
                return True
        return False
```