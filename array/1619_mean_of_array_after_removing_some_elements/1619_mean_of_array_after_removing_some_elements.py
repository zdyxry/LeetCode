from typing import List

class Solution:
    def trimMean(self, arr: List[int]) -> float:
        arr.sort()
        n = int(len(arr) * 0.05)
        result = arr[n:len(arr)-n]
        return sum(result) / (len(arr) - 2 * n)


arr = [1,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,3]
res = Solution().trimMean(arr)
print(res)