from typing import List


class Solution:
    def shuffle(self, nums: List[int], n: int) -> List[int]:
        res = []
        for i, j in zip(nums[0:n],nums[n:]):
            res+=[i,j]
        return res


res = Solution().shuffle([2,5,1,3,4,7], 3)
print(res)