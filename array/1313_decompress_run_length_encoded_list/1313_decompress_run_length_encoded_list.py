from typing import List

class Solution:
    def decompressRLElist(self, nums: List[int]) -> List[int]:
        ans = list()
        for i in range(0, len(nums), 2):
            ans.extend([nums[i + 1]] * nums[i])
        return ans


nums = [1,2,3,4]
res = Solution().decompressRLElist(nums)
print(res)