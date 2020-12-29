from typing import List

class Solution:
    def maximumBinaryString(self, binary: str) -> str:
        n = len(binary)
        cnt0 = 0
        start = -1
        for i in range(n):
            if binary[i] == '0':
                cnt0 += 1
                if start == -1:
                    start = i
        if start == -1: return binary
        return '1' * (start + cnt0 - 1) + '0' + '1' * (n - start - cnt0)


binary = "000110"
res = Solution().maximumBinaryString(binary)
print(res)