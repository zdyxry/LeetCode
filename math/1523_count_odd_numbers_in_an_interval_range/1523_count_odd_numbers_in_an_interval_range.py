class Solution:
    def countOdds(self, low: int, high: int) -> int:
        res = high // 2
        if high % 2 == 1:
            res += 1
        if low - 1 >= 0:
            res -= (low - 1) // 2
            if (low - 1) % 2 == 1:
                res -= 1
        return res


low = 3
high = 7
res = Solution().countOdds(low, high)
print(res)
