import collections
from typing import List


class Solution:
    def canArrange(self, arr: List[int], k: int) -> bool:
        mod = collections.Counter(num % k for num in arr)
        for t, occ in mod.items():
            if t > 0 and (k - t not in mod or mod[k - t] != occ):
                return False
        return mod[0] % 2 == 0


arr =[1,2,3,4,5,10,6,7,8,9]
k = 5
res = Solution().canArrange(arr, k)
print(res)