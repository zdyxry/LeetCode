from typing import List

class Solution:
    def containsPattern(self, arr: List[int], m: int, k: int) -> bool:
        i = 0
        while i < len(arr):
            p = arr[i:i+m]
            if p * k == arr[i:i+m*k]:
                return True
            i += 1
        return False


arr = [1,2,4,4,4,4]
m = 1
k = 3
res = Solution().containsPattern(arr,m , k)
print(res)