class Solution(object):
    def multiply(self, num1, num2):
        inum1, inum2 = 0, 0
        for d in num1:
            inum1 = inum1*10 + (ord(d)-48)
        for d in num2:
            inum2 = inum2*10 + (ord(d)-48)
        
        return str(inum1 * inum2)


num1, num2 = "2", "3"
res = Solution().multiply(num1, num2)
print(res)