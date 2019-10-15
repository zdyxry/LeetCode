class Solution(object):
    def findDuplicates(self, nums):
        result = []
        for i in nums:
            if nums[abs(i)-1] < 0:
                result.append(abs(i))
            else:
                nums[abs(i)-1] *= -1
        return result


nums = [4,3,2,7,8,2,3,1]
res = Solution().findDuplicates(nums)
print(res)