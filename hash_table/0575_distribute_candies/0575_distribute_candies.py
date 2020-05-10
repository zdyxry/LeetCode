from typing import List

class Solution:
    def distributeCandies(self, candies: List[int]) -> int:
        return min(len(set(candies)), len(candies)//2)


candies = [1,1,2,2,3,3]
res = Solution().distributeCandies(candies)
print(res)