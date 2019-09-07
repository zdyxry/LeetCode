class Solution(object):
    def convertToTitle(self, n):
        if n == 0:
            return ""
        else:
            r = (n-1)%26
            q = (n-1)//26
            return self.convertToTitle(q) + chr(r + ord('A'))


n = 701
res = Solution().convertToTitle(n)
print(res)