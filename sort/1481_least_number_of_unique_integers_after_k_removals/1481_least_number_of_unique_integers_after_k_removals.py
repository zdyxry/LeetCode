from typing import List
from collections import Counter


class Solution:
    def findLeastNumOfUniqueInts(self, arr: List[int], k: int) -> int:
        res = dict(Counter(arr))
        res = sorted(res.items(),key=lambda item:item[1])
        print(res)
        sum = 0
        len_elem = len(res)
        for (elem_k,elem_v) in res:
            if sum + elem_v <= k:
                print(elem_k, elem_v)
                len_elem -=1
                sum += elem_v
            else:
                break
        return len_elem


arr = 4,3,1,1,3,3,2
k = 3
res = Solution().findLeastNumOfUniqueInts(arr, k)
print(res)