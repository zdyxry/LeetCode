461. Hamming Distance

Easy

The Hamming distance between two integers is the number of positions at which the corresponding bits are different.

Given two integers x and y, calculate the Hamming distance.

Note:
0 ≤ x, y < 231.  

Example:

Input: x = 1, y = 4

Output: 2

Explanation:  
1   (0 0 0 1)  
4   (0 1 0 0)  


The above arrows point to positions where the corresponding bits are different.

# 方法
汉明距离，两个数字二进制中不同位数。   
异或，不同为 1，计算 1 数量。


```go
func hammingDistance(x int, y int) int {
    x ^= y

	res := 0
	for x > 0 {
		res += x & 1
		x >>= 1
	}

	return res
}
```

```python
class Solution(object):
    def hammingDistance(self, x, y):
        """
        :type x: int
        :type y: int
        :rtype: int
        """
        x ^= y 
        
        res = 0 
        while x:
            res += x & 1
            x >>= 1
        return res
```