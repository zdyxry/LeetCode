
class Solution(object):
    def arrangeCoins1(self, n):
        level = 0
        count = 0
        while count + level + 1 <= n:
            level += 1
            count += level
        return level

    def arrangeCoins2(self, n):
        left, right = 0, n +1
        while left < right:
            mid = left + (right - left) / 2
            if mid * (mid + 1 ) / 2 <= n:
                left = mid + 1
            else:
                right = mid
        return left - 1


n = 8
res = Solution().arrangeCoins2(n)
print(res)