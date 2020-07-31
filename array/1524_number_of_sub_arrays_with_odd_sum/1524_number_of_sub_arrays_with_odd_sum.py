from typing import List

class Solution:
    def numOfSubarrays(self, arr: List[int]) -> int:
        pre = 0
        odd, even = 0, 1
        for num in arr:
            pre += num
            if pre%2 == 0:
                even +=1
            else:
                odd +=1
        return (even*odd)%(10**9 + 7)


arr = [1,3,5]

res = Solution().numOfSubarrays(arr)
print(res)