89. Gray Code


Medium

The gray code is a binary numeral system where two successive values differ in only one bit.

Given a non-negative integer n representing the total number of bits in the code, print the sequence of gray code. A gray code sequence must begin with 0.

Example 1:

```
Input: 2
Output: [0,1,3,2]
Explanation:
00 - 0
01 - 1
11 - 3
10 - 2

For a given n, a gray code sequence may not be uniquely defined.
For example, [0,2,3,1] is also a valid gray code sequence.

00 - 0
10 - 2
11 - 3
01 - 1
```

Example 2:

```
Input: 0
Output: [0]
Explanation: We define the gray code sequence to begin with 0.
             A gray code sequence of n has size = 2n, which for n = 0 the size is 20 = 1.
             Therefore, for n = 0 the gray code sequence is [0].
```


## 方法

```go
func grayCode(n int) []int {
	if n == 0 {
		return []int{0}
	}

	ans := []int{0, 1}
	for i := 1; i < n; i++ {
		var next []int
		b := 1 << uint(i)
		for j := len(ans) - 1; j >= 0; j-- {
			next = append(next, ans[j]|b)
		}
		ans = append(ans, next...)
	}
	return ans
}


```

```python
class Solution:
    def grayCode(self, n: int) -> List[int]:
        def dfs():
            nonlocal n, code, gray
            if code in visited:
                return
            visited.add(code)
            gray.append(code)
            for i in range(0, n):
                code ^= 1 << i
                dfs()
                code ^= 1 << i

        code = 0
        visited = set()
        gray = []
        dfs()
        return gray
```