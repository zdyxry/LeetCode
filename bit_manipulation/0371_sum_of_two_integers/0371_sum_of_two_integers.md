371. Sum of Two Integers
Easy

Calculate the sum of two integers a and b, but you are not allowed to use the operator + and -.

Example 1:  

Input: a = 1, b = 2  
Output: 3  
Example 2:  

Input: a = -2, b = 3  
Output: 1  


# 方法
`与` 运算求进位
`异或` 运算求低位

```go
func getSum(a int, b int) int {
    return ((a & b) << 1) + (a ^ b)
}
```


```python
class Solution(object):
    def getSum(self, a, b):
        """
        :type a: int
        :type b: int
        :rtype: int
        """
        return ((a & b) << 1) + (a ^ b)
```