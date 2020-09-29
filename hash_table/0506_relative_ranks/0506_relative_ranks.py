from typing import List

class Solution:
    def findRelativeRanks(self, nums: List[int]) -> List[str]:
        hashmap = {}
        ans = [0] * len(nums)
        for i in range(len(nums)):
            hashmap[nums[i]] = i
        nums.sort(reverse = True)
        for i in range(len(nums)):
            if i == 0:
                ans[hashmap[nums[i]]] = 'Gold Medal'
            if i == 1:
                ans[hashmap[nums[i]]] = 'Silver Medal'
            if i == 2:
                ans[hashmap[nums[i]]] = 'Bronze Medal'
            if i > 2:
                ans[hashmap[nums[i]]] = str(i + 1)
        return ans



nums = [5,4,3,2,1]
res = Solution().findRelativeRanks(nums)
print(res)