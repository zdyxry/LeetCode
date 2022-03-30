class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        res = []
        nums.sort()
        length = len(nums)
        for i in range(length - 2):
            if nums[i] > 0 :
                continue
            if i > 0 and nums[i] == nums[i-1]:
                continue

            l, r = i + 1, length -1
            while l < r :
                total = nums[i] + nums[l] + nums[r]

                if total < 0 :
                    l += 1
                elif total > 0 :
                    r -= 1
                else:
                    res.append(nums[i], nums[l], nums[r])
                    while l < r and nums[l] == nums[l+1]:
                        l += 1
                    while l < r and nums[r] == nums[r-1]:
                        r -= 1
                    l += 1
                    r -= 1
        return res

nums = [-1,0,1,2,-1,-4]
res = Solution().threeSum(nums)

print(res)