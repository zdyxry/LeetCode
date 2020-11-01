from typing import List

class Solution:
    def canFormArray(self, arr: List[int], pieces: List[List[int]]) -> bool:
        result = []
        for i in arr:
            for j in pieces:
                if i == j[0]:
                    result += j
        return result == arr


arr = [85]
pieces = [[85]]
res = Solution().canFormArray(arr, pieces)
print(res)