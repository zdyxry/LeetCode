# -*- coding: utf-8 -*-

class Solution:
    # @param n, an integer
    # @return an integer
    def reverseBits(self, n):
        res = 0
        for i in range(32):
            bit = n & 1
            n >>= 1
            res <<= 1
            res += bit
        return res

n = 43261596

print(Solution().reverseBits(n))