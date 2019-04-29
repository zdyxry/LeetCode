342. Power of Four
Easy

Given an integer (signed 32 bits), write a function to check whether it is a power of 4.

Example 1:

Input: 16  
Output: true  
Example 2:  

Input: 5  
Output: false  
Follow up: Could you solve it without loops/recursion?  

# 方法
1) 对4取余，若为0，则除4循环，直至 num == 1  
2) 与 0b11 与操作，若为0，则向右移动2位循环，直至 num == 1  
3) num & (num -1） == 0 ，且 num & 0b01010101010101010101010101010101 == num


```go
func isPowerOfFour(num int) bool {
    if num < 1 {
        return false
    }
    for num %4 == 0 {
        num >>= 2
    }
    return num == 1
}

```

```python3
class Solution:
    def isPowerOfFour(self, num: int) -> bool:
        while num and not (num & 0b11):
            num >>= 2
        return num == 1
```

```python
class Solution(object):
    def isPowerOfFour(self, num):
        """
        :type num: int
        :rtype: bool
        """
        return num > 0 and (num & (num - 1)) == 0 and \
               ((num & 0b01010101010101010101010101010101) == num)
```
