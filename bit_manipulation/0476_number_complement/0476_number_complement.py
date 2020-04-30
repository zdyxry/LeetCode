class Solution:
    def findComplement(self, num: int) -> int:
        n = ''.join(str(1-int(i)) for i in bin(num)[2:]) 
        return int(n, 2)


num = 5
res = Solution().findComplement(num)
print(res)