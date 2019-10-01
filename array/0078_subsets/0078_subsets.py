class Solution(object):
    def subsets(self,nums):
        nums.sort()
        result = [[]]
        for i in xrange(len(nums)):
            size = len(result)
            for j in xrange(size):
                result.append(list(result[j]))
                result[-1].append(nums[i])
        return result

nums = [1,2,3]
res = Solution().subsets(nums)
print(res)