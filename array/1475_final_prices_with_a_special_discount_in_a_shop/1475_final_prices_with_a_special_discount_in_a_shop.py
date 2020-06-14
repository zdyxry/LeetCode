from typing import List


class Solution:
    def finalPrices(self, prices: List[int]) -> List[int]:
        ans = []
        n = len(prices)
        for i in range(n):
            dis = 0
            for j in range(i+1, n):
                if prices[j]<=prices[i]:
                    dis = prices[j]
                    break
            ans.append(prices[i]-dis)
        return ans


prices = [8,4,6,2,3]
res = Solution().finalPrices(prices)
print(res)