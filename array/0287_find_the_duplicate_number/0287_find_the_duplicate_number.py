class Solution(object):
    def findDuplicates(self, nums):
        slow = nums[0]
        fast = nums[nums[0]]
        while slow != fast:
            slow = nums[slow]
            fast = nums[nums[fast]]
        
        fast = 0
        while slow != fast:
            slow = nums[slow]
            fast = nums[fast]
        return slow

nums = [1,3,4,2,2]

res = Solution().findDuplicates(nums)
print(res)