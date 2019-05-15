# -*- coding: utf-8 -*-


class Solution(object):
    def maxProfit(self, prices):
        """
        :type prices: List[int]
        :rtype: int
        """
        n = len(prices)
        profit = 0
        for i in range(1, n):
            if prices[i] > prices[i-1]:
                profit += (prices[i] - prices[i - 1])
        return profit
            

prices = [7,1,5,3,6,4]
print(Solution().maxProfit(prices))