class Solution:
    def reachNumber(self, target: int) -> int:
        target = abs(target)
        p, i = 0, 0
        while p < target or (p + target) % 2 != 0:
            i += 1
            p = p + i
        return i


target = 1
res = Solution().reachNumber(target)
print(res)