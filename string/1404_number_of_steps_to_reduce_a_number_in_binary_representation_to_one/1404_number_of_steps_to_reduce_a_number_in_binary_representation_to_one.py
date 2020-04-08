
class Solution(object):
    def numSteps(self, s):
        cnt = 0
        s_num = int(s, 2)
        while s_num != 1:
            if s_num % 2 == 1:
                s_num += 1
            else:
                s_num = s_num // 2
            cnt += 1
        return cnt


s = "1101"
res = Solution().numSteps(s)
print(res)