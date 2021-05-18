from typing import List

class Solution:
    def maximumPopulation(self, logs: List[List[int]]) -> int:
        nums = [0] * 2051
        ret = 1950
        for i in logs:
            for j in range(i[0], i[1]):
                nums[j] += 1
        for x in range(1950, len(nums)):
            if nums[x] > nums[ret]:
                ret = x
        return ret


logs = [[1993,1999],[2000,2010]]
res = Solution().maximumPopulation(logs)
print(res)
