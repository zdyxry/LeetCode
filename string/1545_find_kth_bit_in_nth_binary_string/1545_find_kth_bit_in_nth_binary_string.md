1545. Find Kth Bit in Nth Binary String


Medium


Given two positive integers n and k, the binary string  Sn is formed as follows:

```
S1 = "0"
Si = Si-1 + "1" + reverse(invert(Si-1)) for i > 1
Where + denotes the concatenation operation, reverse(x) returns the reversed string x, and invert(x) inverts all the bits in x (0 changes to 1 and 1 changes to 0).
```

For example, the first 4 strings in the above sequence are:

```
S1 = "0"
S2 = "011"
S3 = "0111001"
S4 = "011100110110001"
Return the kth bit in Sn. It is guaranteed that k is valid for the given n.
```
 

Example 1:

```
Input: n = 3, k = 1
Output: "0"
Explanation: S3 is "0111001". The first bit is "0".
```

Example 2:

```
Input: n = 4, k = 11
Output: "1"
Explanation: S4 is "011100110110001". The 11th bit is "1".
```

Example 3:

```
Input: n = 1, k = 1
Output: "0"
```

Example 4:

```
Input: n = 2, k = 3
Output: "1"
```
 

Constraints:

1 <= n <= 20  
1 <= k <= 2n - 1

## 方法


```python
class Solution:
    def findKthBit(self, n: int, k: int) -> str:
        flip = 0
        l = 2 ** n - 1
        while k > 1:
            if k == l // 2 + 1:
                return str(1 ^ flip)
            if k > l // 2:
                k = l + 1 - k
                flip = 1 - flip
            l //= 2
        return str(flip)
```

```python
class Solution:
    def findKthBit(self, n: int, k: int) -> str:
        s = [0] * (n+1)
        s[0] = '0'
        for i in range(1, n+1):
            s[i] = s[i-1] + '1' + self.reverse(s[i-1])
        return s[n][k-1]

    def reverse(self, s):
        res = []
        for i in s:
            if i == '1':
                res.append('0')
            else:
                res.append('1')
        return ''.join(res[::-1])
```