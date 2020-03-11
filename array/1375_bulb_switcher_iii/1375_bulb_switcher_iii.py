
class Solution(object):
    def numTimesAllBlue(self, light):
        res = 0
        max_idx = 0
        for i,b in enumerate(light):
            max_idx = max(max_idx, b)
            if i + 1 == max_idx:
                res += 1
        return res


light =[2,1,3,5,4]
res = Solution().numTimesAllBlue(light)
print(res)