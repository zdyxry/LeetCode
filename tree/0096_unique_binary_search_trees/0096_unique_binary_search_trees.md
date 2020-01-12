96. Unique Binary Search Trees


Medium


Given n, how many structurally unique BST's (binary search trees) that store values 1 ... n?

Example:

```
Input: 3
Output: 5
Explanation:
Given n = 3, there are a total of 5 unique BST's:

   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3
```

## 方法

```go
func numTrees(n int) int {
    if n == 0 {
		return 1
	}
	if n == 1 || n == 2 {
		return n
	}

	if n == 3 {
		return 5
	}

	res := 0
	// 左右对称，所以只做一半
	for i := 1; i <= n/2; i++ {
		res += numTrees(i-1) * numTrees(n-i)
	}
	res *= 2
	// 处理 n 为奇数的情况
	if n%2 == 1 {
		temp := numTrees(n / 2)
		res += temp * temp
	}

	return res
}
```


```python
class Solution(object):
    def numTrees(self, n):
        """
        :type n: int
        :rtype: int
        """
        if n == 0:
            return 1

        def combination(n, k):
            count = 1
            # C(n, k) = (n) / 1 * (n - 1) / 2 ... * (n - k + 1) / k
            for i in xrange(1, k + 1):
                count = count * (n - i + 1) / i
            return count

        return combination(2 * n, n) - combination(2 * n, n - 1)
```