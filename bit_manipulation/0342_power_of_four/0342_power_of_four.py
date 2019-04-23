# -*- coding: utf-8 -*-

class Solution:
    def isPowerOfFour(self, num):
        while num and not (num & 0b11):
            num >>= 2
        return num == 1

print(Solution().isPowerOfFour(16))