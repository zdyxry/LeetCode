
class Solution(object):
    def minSteps(self, s, t):
        cnt = [0] * 26
        for i in range(len(s)):
            cnt[ord(s[i]) - ord('a')] += 1
            cnt[ord(t[i]) - ord('a')] -= 1
        res = 0
        for i in cnt:
            if i > 0:
                res += i
        return res


s = "bab"
t = "aba"
res = Solution().minSteps(s, t)
print(res)