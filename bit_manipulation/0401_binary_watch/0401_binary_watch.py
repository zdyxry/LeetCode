# -*- coding: utf-8 -*-

class Solution(object):
    def readBinaryWatch(self, num):
        result = []
        for h in range(12):
            for m in range(60):
                if (bin(h) + bin(m)).count('1') == num:
                    result.append('%d:%02d' % (h, m))
        return result

print(Solution().readBinaryWatch(2))