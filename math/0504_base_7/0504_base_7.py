
class Solution(object):
    def convertToBase7(self, num: int) -> str:
        pnum = abs(num)
        negate = pnum != num
        base7 = ""
        while pnum >= 7:
            mod = pnum % 7
            base7 += str(mod)
            pnum = pnum // 7
        base7 += str(pnum)
        if negate:
            base7 += "-"
        return base7[::-1]


num = 100
res = Solution().convertToBase7(num)
print(res)