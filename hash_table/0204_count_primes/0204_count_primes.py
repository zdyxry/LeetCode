# -*- coding: utf-8 -*-

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


print(Solution().countPrimes(50))
