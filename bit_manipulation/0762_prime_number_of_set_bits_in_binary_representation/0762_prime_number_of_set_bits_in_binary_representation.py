# -*- coding: utf-8 -*-

class Solution(object):
    def countPrimeSetBits(self, L, R):
        """
        :type L: int
        :type R: int
        :rtype: int
        """
        primes = {2, 3, 5, 7, 11, 13, 17, 19}
        return sum(map(lambda x: bin(x).count('1') in primes, range(L, R+1)))

print(Solution().countPrimeSetBits(842,888))