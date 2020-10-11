import math


class Solution:
    def minimumOneBitOperations(self, n: int) -> int:
        if not n:return 0
        head = 1<<int(math.log2(n))
        return head + self.minimumOneBitOperations((n^head)^(head>>1))

n = 5
res = Solution().minimumOneBitOperations(n)
print(res)