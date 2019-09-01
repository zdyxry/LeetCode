class Solution(object):
    def isSubsequence(self, s, t):
        if not s :
            return True
        
        i = 0
        for c in t:
            if c == s[i]:
                i += 1
            if i == len(s):
                break
        return i == len(s)

s = "abc"
t = "ahbgdc"
res = Solution().isSubsequence(s,t)
print(res)