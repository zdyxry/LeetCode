from typing import List

class Solution:
    def kidsWithCandies(self, candies: List[int], extraCandies: int) -> List[bool]:
        return [x+extraCandies>=max(candies) for x in candies]


candies = [2,3,5,1,3]
extraCandies = 3
res = Solution().kidsWithCandies(candies, extraCandies)
print(res)