class Solution:
    def minimumCost(self, cost: List[int]) -> int:
        n = len(cost)
        
        cost.sort(reverse = True)
        res = 0

        for i in range(n):
            if (i + 1) % 3 == 0:
                continue
            res += cost[i]

        return res
