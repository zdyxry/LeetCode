
class Solution:
    def balancedStringSplit(self, s: str) -> int:
        split = 0
        unbalance = 0
        for i in s:
            unbalance += 1 if i == 'R' else -1    # 'R' = +1, 'L' = -1
            if not unbalance:   # if unbalance == 0
                split += 1
        return split


s = "RLRRLLRLRL"
res = Solution().balancedStringSplit(s)
print(res)