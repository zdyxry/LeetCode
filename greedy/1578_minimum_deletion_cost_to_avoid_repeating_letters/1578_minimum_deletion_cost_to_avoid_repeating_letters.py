from typing import List

class Solution:
    def minCost(self, s: str, cost: List[int]) -> int:
        pre = 0
        res = 0
        for idx in range(1, len(s)):
            if s[idx] == s[pre]:
                res += min(cost[idx], cost[pre])
            if s[idx] != s[pre] or cost[pre] < cost[idx]:
                pre = idx
        return res


s = "abaac"
cost = [1,2,3,4,5]
res = Solution().minCost(s, cost)
print(res)