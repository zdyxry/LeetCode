
class Solution(object):
    def summaryRanges(self, nums):
        ranges = []
        if not nums:
            return ranges

        start, end = nums[0], nums[0]
        for i in xrange(1, len(nums) + 1):
            if i < len(nums) and nums[i] == end + 1:
                end = nums[i]
            else:
                interval = str(start)
                if start != end:
                    interval += "->" + str(end)
                ranges.append(interval)
                if i < len(nums):
                    start = end = nums[i]
            
        return ranges

nums = [0,1,2,4,5,7]
res = Solution().summaryRanges(nums)
print(res)