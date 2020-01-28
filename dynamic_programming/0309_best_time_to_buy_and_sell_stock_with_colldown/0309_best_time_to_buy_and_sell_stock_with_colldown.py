
class Solution(object):
    def maxProfit(self, prices):
        n = len(prices)
        if n < 2:
            return 0
        if n == 2:
            return max(0, prices[n - 1] - prices[n - 2])
        dp_prev_prev = 0
        dp_prev = max(0, prices[n - 1] - prices[n - 2])
        dp_curr = max(dp_prev, prices[n - 2] - prices[n - 3], prices[n - 1] - prices[n - 3])
        max_so_far = max(prices[n - 1], prices[n - 2])
        for i in xrange(n - 4, -1, -1):
            max_so_far = max(max_so_far, prices[i + 1] + dp_prev_prev)
            dp_prev_prev = dp_prev
            dp_prev = dp_curr
            dp_curr = max(dp_curr, max_so_far - prices[i])
        return dp_curr


prices = [1, 2, 3, 0, 2]
res = Solution().maxProfit(prices)
print(res)