
class Solution(object):
    def coinChange(self, coins, amount):
        INF = 0x7fffffff
        amounts = [INF] * (amount +1)
        amounts[0] = 0
        for i in xrange(amount+1):
            if amounts[i] != INF:
                for coin in coins:
                    if i + coin <= amount:
                        amounts[i + coin] = min(amounts[i + coin], amounts[i] +1)
        return amounts[amount] if amounts[amount] != INF else -1


coins = [1,2,5]
amount = 11
res = Solution().coinChange(coins, amount)
print(res)