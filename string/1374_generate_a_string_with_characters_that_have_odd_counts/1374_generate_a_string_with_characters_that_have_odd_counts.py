
class Solution(object):
    def generateTheString(self, n):
        if n % 2 == 0 :
            return 'a' * (n-1) + 'b'
        else :
            return 'a' * n 


n = 4 
res = Solution().generateTheString(n)
print(res)