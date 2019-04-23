191. Number of 1 Bits
Easy

Write a function that takes an unsigned integer and return the number of '1' bits it has (also known as the Hamming weight).


Example 1:

Input: 00000000000000000000000000001011
Output: 3
Explanation: The input binary string 00000000000000000000000000001011 has a total of three '1' bits.
Example 2:

Input: 00000000000000000000000010000000
Output: 1
Explanation: The input binary string 00000000000000000000000010000000 has a total of one '1' bit.
Example 3:

Input: 11111111111111111111111111111101
Output: 31
Explanation: The input binary string 11111111111111111111111111111101 has a total of thirty one '1' bits.

# 方法
按位与，若为真则计数器 +1

```go
func hammingBits(num uint32) int {
    var res uint32
    for num != 0 {
        res += (num & 1)
        num >>= 1
    }
    return int(res)
}

func hammingBits2(num uint32) int {
    var res int
    for num != 0 {
        res += 1
        num &= (num - 1)
    }
    return res
}
```

```python
class Solution(object):
    def hammingWeight(self, n):
        """
        :type n: int
        :rtype: int
        """
        res = 0
        while n:
            res += (n & 1)
            n >>= 1
        return res
```
