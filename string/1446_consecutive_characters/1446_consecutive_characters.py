
class Solution:
    def maxPower(self, s: str) -> int:
        if len(s) == 1:
            return 1
        tmp = s[0]
        cnt = 1
        max_cnt = 0
        for i in s[1:]:
            if i == tmp:
                cnt += 1
            else:
                tmp = i
                cnt = 1
            max_cnt = max(max_cnt, cnt)
        return max_cnt


s = "leetcode"
res = Solution().maxPower(s)
print(res)