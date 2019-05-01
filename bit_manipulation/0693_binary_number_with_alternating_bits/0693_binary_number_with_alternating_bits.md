693. Binary Number with Alternating Bits

Easy

Given a positive integer, check whether it has alternating bits: namely, if two adjacent bits will always have different values.  


Example 1:  
Input: 5  
Output: True  
Explanation:  
The binary representation of 5 is: 101  
Example 2:  
Input: 7  
Output: False  
Explanation:  
The binary representation of 7 is: 111.  
Example 3:  
Input: 11  
Output: False  
Explanation:  
The binary representation of 11 is: 1011.  
Example 4:  
Input: 10  
Output: True  
Explanation:  
The binary representation of 10 is: 1010.  

# 方法
1）遍历二进制数字，记录最后一位，若相同，则 false   
2）移位异或，加 1 相与为 0 

```go
func hasAlternatingBits(n int) bool {
	prevBit := -1
	for n > 0 {
		lastBit := n & 1
		if lastBit == prevBit {
			return false
		}
		prevBit = lastBit
		n >>= 1
	}
	return true
}
```




```python
class Solution:
    def hasAlternatingBits(self, n):
        if not n:
            return False
        num = n ^ (n >> 1)
        return not (num & (num + 1))
```