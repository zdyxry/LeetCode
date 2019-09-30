class Solution(object):
    def sortColors(self, nums):
        def triPartition(nums, target):
            i,j,n = 0, 0,len(nums) -1
            while j <= n:
                if nums[j] < target:
                    nums[i], nums[j] = nums[j], nums[i]
                    i += 1
                    j += 1
                elif nums[j] > target:
                    nums[j], nums[n] = nums[n], nums[j]
                    n -=1
                else:
                    j +=1

        triPartition(nums, 1)


nums = [2,0,2,1,1,0]
Solution().sortColors(nums)
print(nums)