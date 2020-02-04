
class Solution(object):
    def findTargetSumWays(self, nums, S):
        from collections import defaultdict
        memo = {0:1}
        for x in nums:
            m = defaultdict(int)
            for s, n in memo.iteritems():
                m[s + x] += n
                m[s - x] += n
            memo = m
            print memo
        return memo[S]


nums = [1,1,1,1,1]
S = 3
res = Solution().findTargetSumWays(nums, S)
print(res)