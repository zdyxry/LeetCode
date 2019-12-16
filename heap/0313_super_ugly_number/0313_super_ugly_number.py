import heapq


class Solution(object):
    def nthSuperUglyNumber(self, n, primes):
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
            heapq.heappush(heap, (primes[k] * uglies[idx[k]],k))

        return uglies[-1]


n = 12
primes = [2,7,13,19]
res = Solution().nthSuperUglyNumber(n, primes)
print(res)