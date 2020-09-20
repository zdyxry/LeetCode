from typing import List

class Solution:
    def sumOddLengthSubarrays(self, arr: List[int]) -> int:
        res = 0
        for i in range(len(arr)+1):
            if i % 2:
                if i == 1:
                    res += sum(arr)
                else:
                    j = 0
                    while j + i <= len(arr):
                        res += sum(arr[j:j+i])
                        j += 1
        return res


arr = [1,4,2,5,3]
res = Solution().sumOddLengthSubarrays(arr)
print(res)