class Solution(object):
    def repeatedSubstringPattern(self, s):
        for l in range(1, len(s)//2+1):
            if len(s)%l == 0:
                str_sub = s[:l]
                n = len(s)//l
                if str_sub * n == s:
                    return True
        return False


print(Solution().repeatedSubstringPattern("aba"))