from typing import List

class Solution:
    def countSubstrings(self, s: str, t: str) -> int:
        # 相同长度下s1s2是否满足只差一个字符
        def isvalid(s1,s2):
            diff = 0
            for i in range(len(s1)):
                if s1[i] != s2[i]:
                    diff += 1
            return True if diff == 1 else False

        # 预处理s所有子串
        substr = []
        for i in range(len(s)):
            for j in range(len(s) - i):
                substr.append(s[i:i + j + 1])
        ans = 0
        n = len(t)
        for w in substr:
            l = len(w)
            for left in range(len(t) - l + 1):
                if isvalid(w,t[left:left + l]):
                    ans += 1
        return ans


s = "aba"
t = "baba"
res = Solution().countSubstrings(s, t)
print(res)