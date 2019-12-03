
class Solution(object):
    def removeKdigits(self, num, k):
        if len(num) <= k:
            return "0"
        
        ret = []
        for i in range(len(num)):
            while k > 0 and ret and ret[-1] > num[i]:
                ret.pop()
                k-=1
            ret.append(num[i])
        while k > 0 and ret:
            ret.pop()
            k -= 1
        
        s = "".join(ret).lstrip("0")
        return s if s else "0"


num = "1432219"
k = 3
res = Solution().removeKdigits(num, k)
print(res)