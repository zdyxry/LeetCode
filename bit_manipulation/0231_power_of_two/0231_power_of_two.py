# -*- coding: utf-8 -*-

class Solution:
    def isPowerOfTwo(self, n):
        if n < 1:
            return False
        
        while n > 1:
            if n % 2 == 1:
                return False
            n >>= 1
        return True

print(Solution().isPowerOfTwo(5))