from typing import List

class Solution:
    def closestCost(self, baseCosts: List[int], toppingCosts: List[int], target: int) -> int:
        ans = baseCosts[0]
        li = baseCosts
        for cost in toppingCosts:
            for c in li.copy():
                li.append(c + cost)
                li.append(c + 2 * cost)
        for a in li:
            if abs(a - target) < abs(ans - target) or abs(a - target) == abs(ans - target) and a < ans:
                ans = a
        return ans



baseCosts = [1,7]
toppingCosts = [3,4]
target = 10
res = Solution().closestCost(baseCosts, toppingCosts, target)
print(res)