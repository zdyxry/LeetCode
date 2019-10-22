class Solution(object):
    def findMaximumXOR(self, nums):
        result = 0

        for i in reversed(xrange(32)):
            result <<= 1
            prefixes = set()
            for n in nums:
                prefixes.add(n>>i)
            for p in prefixes:
                if (result | 1) ^ p in prefixes:
                    result += 1
                    break

        return result

nums = [3,10,5,25,2,8]
res = Solution().findMaximumXOR(nums)
print(res)