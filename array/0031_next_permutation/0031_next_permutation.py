class Solution(object):
    def nextPermutation(self, nums):
        i = len(nums) - 1
        j = len(nums) - 1
        while i > 0 and nums[i-1] >= nums[i]:
            i -= 1
        if i == 0 :
            nums.reverse()
            return
        k = i - 1
        print(i,j,k)
        while nums[j] <= nums[k]:
            j -= 1
        nums[k], nums[j] = nums[j], nums[k]
        l, r = k+1, len(nums) -1
        while l < r:
            nums[l], nums[r] = nums[r], nums[l]
            l += 1
            r -= 1

nums = [1,2,7,4,3,1]
Solution().nextPermutation(nums)
print(nums)