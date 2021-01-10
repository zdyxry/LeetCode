class Solution:
    def totalMoney(self, n: int) -> int:
        cur = 7
        ret = 0
        for i in range(n):
            if i % 7 == 0:
                cur -= 6
            ret += cur
            cur += 1
        return ret


n = 10
res = Solution().totalMoney(n)
print(res)