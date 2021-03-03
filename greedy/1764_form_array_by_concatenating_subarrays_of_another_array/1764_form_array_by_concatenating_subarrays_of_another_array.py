from typing import List

class Solution:
    def canChoose(self, groups: List[List[int]], nums: List[int]) -> bool:
        i = 0
        n = len(nums)
        for group in groups:
            find = False
            while i + len(group) <= n:
                if group == nums[i:i+len(group)]:
                    find = True
                    i += len(group)
                    break
                else:
                    i += 1
            if not find:
                return False
        return True


groups = [[1,-1,-1],[3,-2,0]]
nums = [1,-1,0,1,-1,-1,3,-2,0]
res = Solution().canChoose(groups, nums)
print(res)