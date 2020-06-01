import collections

from typing import List


class Solution:
    
    def canBeEqual(self, target: List[int], arr: List[int]) -> bool:
        c = collections.Counter(target)
        for a in arr:
            c[a] -= 1
            if c[a] < 0:
                return False
        return True    


target = [1, 2, 3, 4]
arr = [2, 4, 1, 3]
res = Solution().canBeEqual(target, arr)
print(res)