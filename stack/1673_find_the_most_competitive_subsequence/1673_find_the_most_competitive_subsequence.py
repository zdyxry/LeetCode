from typing import List


class Solution:
    def mostCompetitive(self, nums: List[int], k: int) -> List[int]:
        ans = []
        r = len(nums) - k
        for i in range(len(nums)):
            while len(ans) > 0 and ans[-1] > nums[i] and r > 0:
                ans.pop()
                r -= 1
            ans.append(nums[i])
        return ans[:k]
        

nums = [3,5,2,6]
k = 2
res = Solution().mostCompetitive(nums, k)
print(res)