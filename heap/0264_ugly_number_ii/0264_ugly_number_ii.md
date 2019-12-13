264. Ugly Number II


Medium


Write a program to find the n-th ugly number.

Ugly numbers are positive numbers whose prime factors only include 2, 3, 5. 

Example:

```
Input: n = 10
Output: 12
Explanation: 1, 2, 3, 4, 5, 6, 8, 9, 10, 12 is the sequence of the first 10 ugly numbers.
```

Note:  

1 is typically treated as an ugly number.

n does not exceed 1690.

## 方法


```go
func nthUglyNumber(n int) int {
    if n <= 6 {
		return n
	}

	pos := []int{0, 0, 0}
	factors := []int{2, 3, 5}
	candidates := []int{2, 3, 5}

	res := make([]int, n)

	res[0] = 1

	for i := 1; i < n; i++ {
		res[i] = min(candidates)
		for j := 0; j < 3; j++ {
			if res[i] == candidates[j] {
				pos[j]++
				candidates[j] = res[pos[j]] * factors[j]
			}
		}
	}

	return res[n-1]
}

func min(candidates []int) int {
	return lesser(
		candidates[0],
		lesser(
			candidates[1],
			candidates[2],
		),
	)
}

func lesser(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```


```python
def nums():
    ret = [0, 1]
    i2 = i3 = i5 = 1
    for i in range(2, 1691):
        n2 = ret[i2] * 2
        n3 = ret[i3] * 3
        n5 = ret[i5] * 5
        n = min(n2, n3, n5)
        ret.append(n)
        if n == n2:
            i2 += 1
        if n == n3:
            i3 += 1
        if n == n5:
            i5 += 1
    return ret

class Solution(object):
    u = nums()
    def nthUglyNumber(self, n):
        """
        :type n: int
        :rtype: int
        """
        return self.u[n]
```


```python
class Solution(object):
    # @param {integer} n
    # @return {integer}
    def nthUglyNumber(self, n):
        ugly_number = 0

        heap = []
        heapq.heappush(heap, 1)
        for _ in xrange(n):
            ugly_number = heapq.heappop(heap)
            if ugly_number % 2 == 0:
                heapq.heappush(heap, ugly_number * 2)
            elif ugly_number % 3 == 0:
                heapq.heappush(heap, ugly_number * 2)
                heapq.heappush(heap, ugly_number * 3)
            else:
                heapq.heappush(heap, ugly_number * 2)
                heapq.heappush(heap, ugly_number * 3)
                heapq.heappush(heap, ugly_number * 5)

        return ugly_number
```