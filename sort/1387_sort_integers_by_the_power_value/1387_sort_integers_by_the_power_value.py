
class Solution(object):
    def getKth(self, lo, hi, k):
        self.memo = {}
        def helper(n):
            if n == 1:
                return 0
            if n in self.memo:
                return self.memo[n]
            if n % 2 :
                ans = helper((n*3+1)/2)+2
            else:
                ans = helper(n/2)+1
            self.memo[n] = ans
            return ans
        tmp = sorted([[helper(n),n] for n in range(lo, hi+1)])
        return tmp[k-1][1]


lo, hi, k = 12, 15, 2
res = Solution().getKth(lo, hi, k)
print(res)