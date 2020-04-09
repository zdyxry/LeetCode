
class Solution(object):
    def longestDiverseString(self, a: int, b: int, c: int) -> str:
        x = [[a, 'a'], [b, 'b'], [c, 'c']]
        res = ""
        while True:
            for num in sorted(x, reverse=True):
                if num[0] <= 0:
                    return res
                if len(res) >= 2 and res[-2:] == num[1]*2:
                    continue
                res += num[1]
                num[0] -= 1
                break
        return res


a = 1
b = 1
c = 7
res = Solution().longestDiverseString(a, b, c)
print(res)