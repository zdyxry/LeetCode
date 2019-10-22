class Solution(object):
    def countBits(self, num):
        res = [0]
        for i in xrange(1, num + 1):
            res.append((i&1)+res[i >> 1])
        return res

    def countBits2(self, num):
        res = [0]
        for i in xrange(1,num+1):
            r = res[i&(i-1)] + 1
            res.append(r)

        return res

num = 5
res = Solution().countBits2(num)
print(res)