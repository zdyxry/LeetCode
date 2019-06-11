204. Count Primes

Easy

Count the number of prime numbers less than a non-negative number, n.

Example:
```
Input: 10
Output: 4
Explanation: There are 4 prime numbers less than 10, they are 2, 3, 5, 7.
```


# 方法

筛选法，假设所有数字为质数，若当前数为某数的倍数，则肯定不为质数，标记其为 false，并总数-1。


```go
func countPrimes(n int) int {
    flagArray := make([]bool, n)
    result := 0
    for i := 2; i < n; i++ {
        if flagArray[i] == false{
            result++
        
        for j:=2;i * j < n;j++ {
                flagArray[i * j] = true
            }
        }
    }
    return result
}

```

```go
func countPrimes(n int) int {
    if n < 3 {
		return 0
	}

	// composite：合数，prime 的反义词
	isComposite := make([]bool, n)
	// 小于 n 的数中，有一半是偶数，除2以外，都不是质数
	// 所以，count 的上限是 n / 2
	count := n / 2

	for i := 3; i*i < n; i += 2 {
		if isComposite[i] {
			continue
		}

		// j 是 i 的倍数，所以 j 肯定不是 质数
		// 对所有的 j 进行标记
		for j := i * i; j < n; j += 2 * i {
			if !isComposite[j] {
				// j 还没有被别的 i 标记过了
				// 修改 count
				count--
				// 对 j 进行标记
				isComposite[j] = true
			}
		}
	}

	return count
}
```


```python
class Solution(object):
    def countPrimes(self, n):
        """
        :type n: int
        :rtype: int
        """
        if n <= 2:
            return 0

        is_prime = [True] * n
        num = n / 2
        for i in xrange(3, n, 2):
            if i * i >= n:
                break

            if not is_prime[i]:
                continue

            for j in xrange(i*i, n, 2*i):
                if not is_prime[j]:
                    continue

                num -= 1
                is_prime[j] = False

        return num

```