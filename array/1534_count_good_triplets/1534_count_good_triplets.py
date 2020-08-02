from typing import List

class Solution:
    def countGoodTriplets(self, arr: List[int], a: int, b: int, c: int) -> int:
        if len(arr) < 3: return 0
        
        ans = 0
        for i in range(0, len(arr)):
            for j in range(i+1, len(arr)):
                for k in range(j+1, len(arr)):
                    if abs(arr[i] - arr[j]) <= a and abs(arr[j] - arr[k]) <= b and abs(arr[i] - arr[k]) <= c:
                        ans += 1
        return ans


arr = [3,0,1,1,9,7]
a = 7
b = 2
c = 3
res = Solution().countGoodTriplets(arr, a, b, c)
print(res)