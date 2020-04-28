from typing import List


class Solution:
    def findDiagonalOrder(self, nums: List[List[int]]) -> List[int]:
        m = []
        for i, row in enumerate(nums):
            for j, v in enumerate(row):
                if i + j >= len(m):
                    m.append([])
                m[i+j].append(v)
        return [v for d in m for v in reversed(d)]


nums = [[1,2,3],[4,5,6],[7,8,9]]
res = Solution().findDiagonalOrder(nums)
print(res)