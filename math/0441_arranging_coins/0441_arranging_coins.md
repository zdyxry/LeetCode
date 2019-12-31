441. Arranging Coins


Easy


You have a total of n coins that you want to form in a staircase shape, where every k-th row must have exactly k coins.

Given n, find the total number of full staircase rows that can be formed.

n is a non-negative integer and fits within the range of a 32-bit signed integer.

Example 1:

```
n = 5

The coins can form the following rows:
¤
¤ ¤
¤ ¤

Because the 3rd row is incomplete, we return 2.
```

Example 2:

```
n = 8

The coins can form the following rows:
¤
¤ ¤
¤ ¤ ¤
¤ ¤

Because the 4th row is incomplete, we return 3.
```

## 方法

```go
func arrangeCoins(n int) int {
    if n <= 0 {
		return 0
	}
	x := math.Sqrt(2*float64(n)+0.25) - 0.5
	return int(x)
}
```

```python
class Solution(object):
    def arrangeCoins(self, n):
        """
        :type n: int
        :rtype: int
        """
        level = 0
        count = 0
        while count + level + 1 <= n:
            level += 1
            count += level
        return level
```


```python
class Solution(object):
    def arrangeCoins(self, n):
        """
        :type n: int
        :rtype: int
        """
        left, right = 0, n + 1 #[left, right)
        while left < right:
            mid = left + (right - left) / 2
            if mid * (mid + 1) / 2 <= n:
                left = mid + 1
            else:
                right = mid
        return left - 1
```