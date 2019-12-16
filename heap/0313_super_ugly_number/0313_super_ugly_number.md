313. Super Ugly Number


Medium


Write a program to find the nth super ugly number.

Super ugly numbers are positive numbers whose all prime factors are in the given prime list primes of size k.

Example:

```
Input: n = 12, primes = [2,7,13,19]
Output: 32 
Explanation: [1,2,4,7,8,13,14,16,19,26,28,32] is the sequence of the first 12 
             super ugly numbers given primes = [2,7,13,19] of size 4.
```

Note:

1 is a super ugly number for any given primes.   
The given numbers in primes are in ascending order.   
0 < k ≤ 100, 0 < n ≤ 106, 0 < primes[i] < 1000.   
The nth super ugly number is guaranteed to fit in a 32-bit signed integer.


## 方法


```go
func nthSuperUglyNumber(n int, primes []int) int {
    ugly := make([]int, n)
    idx := make([]int, len(primes))
    val := make([]int, len(primes))
    for i := 0; i < len(primes); i++{
        val[i] = 1
    }
    
    next := 1
    for i := 0; i < n; i++{
        ugly[i] = next
        
        next = math.MaxInt32
        for j := 0; j < len(primes); j++{
            if val[j] == ugly[i]{
                
                val[j] = ugly[idx[j]] * primes[j]
                idx[j]++
                
            }
            next = min(next, val[j])
        }
    }
    return ugly[n-1]
}

func min(x, y int) int{
    if x < y{
        return x
    }
    return y
}
```


```python
class Solution(object):
    def nthSuperUglyNumber(self, n, primes):
        """
        :type n: int
        :type primes: List[int]
        :rtype: int
        """
        heap, uglies, idx, ugly_by_last_prime = [], [0] * n, [0] * len(primes), [0] * n
        uglies[0] = 1

        for k, p in enumerate(primes):
            heapq.heappush(heap, (p, k))

        for i in xrange(1, n):
            uglies[i], k = heapq.heappop(heap)
            ugly_by_last_prime[i] = k
            idx[k] += 1
            while ugly_by_last_prime[idx[k]] > k:
                idx[k] += 1
            heapq.heappush(heap, (primes[k] * uglies[idx[k]], k))

        return uglies[-1]
```