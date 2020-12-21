from typing import List


class Solution:
    def maximumUniqueSubarray(self, nums: List[int]) -> int:
        dic = {}  
        j = 0
        res = 0
        s = 0
        length = len(nums)
        for i in range(length):
            if nums[i] not in dic:
                dic[nums[i]] = 1
            else:
                dic[nums[i]] += 1
            s += nums[i]
            while dic[nums[i]] > 1:
                dic[nums[j]] -= 1
                s -= nums[j]
                j += 1
            res = max(res,s)
        return res if res > s else s



nums = [4,2,4,5,6]
res = Solution().maximumUniqueSubarray(nums)
print(res)