868. Binary Gap  

Easy

Given a positive integer N, find and return the longest distance between two consecutive 1's in the binary representation of N.

If there aren't two consecutive 1's, return 0.

Example 1:

Input: 22  
Output: 2  
Explanation:  
22 in binary is 0b10110.  
In the binary representation of 22, there are three ones, and two consecutive pairs of 1's.  
The first consecutive pair of 1's have distance 2.  
The second consecutive pair of 1's have distance 1.  
The answer is the largest of these two distances, which is 2.  
Example 2:  

Input: 5  
Output: 2  
Explanation:  
5 in binary is 0b101.  
Example 3:  

Input: 6  
Output: 1  
Explanation:  
6 in binary is 0b110.  
Example 4:  

Input: 8  
Output: 0  
Explanation:  
8 in binary is 0b1000.  
There aren't any consecutive pairs of 1's in the binary representation of 8, so we return 0.  

Note:

1 <= N <= 10^9


# 方法
转换为二进制，记录第一个 1 出现的位置，若下一位不为 1，则距离 +1，取距离最大值。



```go
func binaryGap(N int) int {
	res := 0
	gap := 0

	for N > 0 {
		if N&1 == 1 {
			res = max(res, gap)
			gap = 1
		} else if gap > 0 {
			gap++
		}
		N >>= 1
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```


```python
class Solution:
    def binaryGap(self, N):
        pre = dist = 0
        for i, c in enumerate(bin(N)[2:]):
            if c == "1":
                dist = max(dist, i - pre)
                pre = i
        return dist
```
