class Solution(object):
    def removeDuplicates(self, nums):
        i = 0
        for n in nums:
            print("n is", n)
            if i< 2 or n > nums[i-2]:
                nums[i] = n
                i+= 1
                print("nums is", nums)
        return i

nums = [1,1,1,2,2,3]
res = Solution().removeDuplicates(nums)
print(res)