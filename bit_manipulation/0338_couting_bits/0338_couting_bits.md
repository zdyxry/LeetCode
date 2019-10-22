338. Counting Bits


Medium


Given a non negative integer number num. For every numbers i in the range 0 ≤ i ≤ num calculate the number of 1's in their binary representation and return them as an array.

Example 1:

```
Input: 2
Output: [0,1,1]
```

Example 2:

```
Input: 5
Output: [0,1,1,2,1,2]
```

Follow up:

It is very easy to come up with a solution with run time O(n*sizeof(integer)). But can you do it in linear time O(n) /possibly in a single pass?

Space complexity should be O(n).

Can you do it like a boss? Do it without using any builtin function like __builtin_popcount in c++ or in any other language.


## 方法

```go
func countBits(num int) []int {
    res := make([]int, num+1)
	for i := 1; i <= num; i++ {
		// i>>1 == i/2
		// i&1  == i%2
		// 只是 位运算 更快
		//
		// 观察以下三个数的二进制表示
		// 5 : 101
		// 10: 1010, 10>>1 == 5
		// 11: 1011, 11>>1 == 5
		// 10 的二进制表示，含有 1 的个数，可以由 5 的答案 + 10%2 计算
		// 11 同理
		res[i] = res[i>>1] + i&1
	}
	return res
}
```


```python
class Solution(object):
    def countBits(self, num):
        """
        :type num: int
        :rtype: List[int]
        """
        res = [0]
        for i in xrange(1, num + 1):
            # Number of 1's in i = (i & 1) + number of 1's in (i / 2).
            res.append((i & 1) + res[i >> 1])
        return res
```