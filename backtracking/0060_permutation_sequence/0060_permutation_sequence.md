60. Permutation Sequence


Medium


The set [1,2,3,...,n] contains a total of n! unique permutations.

By listing and labeling all of the permutations in order, we get the following sequence for n = 3:

```
"123"
"132"
"213"
"231"
"312"
"321"
Given n and k, return the kth permutation sequence.
```

Note:

Given n will be between 1 and 9 inclusive.  
Given k will be between 1 and n! inclusive.  

Example 1:

```
Input: n = 3, k = 3
Output: "213"
```

Example 2:
```
Input: n = 4, k = 9
Output: "2314"
```

## 方法

```go
func getPermutation(n int, k int) string {
    if n == 0 {
		return ""
	}

	// 用于存放结果的字符
	res := make([]byte, n)

	// 存放待抓取的数字
	rec := make([]byte, n)
	for i := 0; i < n; i++ {
		rec[i] = byte(i) + '1'
	}

	// 由于排列的序号是从 1 开始的。
	// k 需要减去 1 ，好从 0 开始表示
	k--

	base := 1
	for i := 2; i < n; i++ {
		base *= i
	}

	for i := 0; i < n-1; i++ {
		idx := k / base
		res[i] = rec[idx]
		// 从 rec 中去除已经使用的数 rec[idx]
		rec = append(rec[:idx], rec[idx+1:]...)
		k %= base
		base /= (n - i - 1)
	}
	// 不要忘记最后一个数
	res[n-1] = rec[0]

	return string(res)
}
```


```python
class Solution(object):
    def getPermutation(self, n, k):
        """
        :type n: int
        :type k: int
        :rtype: str
        """
        res = ''
        nums = [i for i in range(1,n+1)]
        for i in range(n-1,0,-1):
            base = math.factorial(i)
            current = 0
            while k > base:
                k -= base
                current += 1
            res += str(nums[current])
            nums.pop(current)
        res += str(nums[0])
        return res
```


```python
class Solution(object):
    def getPermutation(self, n, k):
        """
        :type n: int
        :type k: int
        :rtype: str
        """
        ans = ''
        fact = [1] * n
        num = [str(i) for i in range(1, 10)]
        for i in range(1, n):
            fact[i] = fact[i - 1] * i
        k -= 1
        for i in range(n, 0, -1):
            first = k // fact[i - 1]
            k %= fact[i - 1]
            ans += num[first]
            num.pop(first)
        return ans
```