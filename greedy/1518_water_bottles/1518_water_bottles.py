class Solution:
    def numWaterBottles(self, numBottles: int, numExchange: int) -> int:
        res = numBottles
        while numBottles >= numExchange:
            cnt = numBottles // numExchange
            res += cnt
            numBottles = numBottles - (cnt * numExchange)
            numBottles += cnt
        return res


numBottles = 9
numExchange = 3
res = Solution().numWaterBottles(numBottles, numExchange)
print(res)