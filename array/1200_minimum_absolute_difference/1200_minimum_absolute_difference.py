from typing import List



class Solution:
    def minimumAbsDifference(self, arr: List[int]) -> List[List[int]]:
        arr.sort()
        m = float('inf')
        out = []
        for i in range(1, len(arr)):
            prev = arr[i - 1]
            curr = abs(prev - arr[i])
            if curr < m:
                out = [[prev, arr[i]]]
                m = curr
            elif curr == m: out.append([prev, arr[i]])
        return out


arr = [4,2,1,3]
res = Solution().minimumAbsDifference(arr)
print(res)