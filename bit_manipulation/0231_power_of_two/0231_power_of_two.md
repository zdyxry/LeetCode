231. Power of Two
Easy

Given an integer, write a function to determine if it is a power of two.

Example 1:

Input: 1
Output: true
Explanation: 20 = 1
Example 2:

Input: 16
Output: true
Explanation: 24 = 16
Example 3:

Input: 218
Output: false

# 方法
处理边界条件；
1)对2取余，若不为1，则向右移动1位
2)将数字与 1  进行 `与` 操作，若不为 0 ，则为 False


```go
func isPowerOfTwo(n int) bool {
	if n < 1 {
		return false
	}

	for n > 1 {
		if n%2 == 1 {
			return false
		}
		n >>= 1
	}

	return true
}
```


```python
class Solution:
    def isPowerOfTwo(self, n: int) -> bool:
        if n < 1:
            return False
        
        while n > 1:
            if n % 2 == 1:
                return False
            n >>= 1
        return True
    
```

