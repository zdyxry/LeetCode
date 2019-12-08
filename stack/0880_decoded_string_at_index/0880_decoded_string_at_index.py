
class Solution(object):
    def decodeAtIndex(self, S, K):
        i = 0
        for c in S:
            if c.isdigit():
                i *= int(c)
            else:
                i += 1
        
        for c in reversed(S):
            K %= i
            if K == 0 and c.isalpha():
                return c
            
            if c.isdigit():
                i /= int(c)
            else:
                i -= 1


S = "leet2code3"
K = 10
res = Solution().decodeAtIndex(S, K)
print(res)
