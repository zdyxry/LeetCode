from typing import List

class Solution:
    def getWinner(self, arr: List[int], k: int) -> int:
        if k >= len(arr): return max(arr) 
        temp,n = arr[0], -1
        for i in arr:
            if temp >= i:
                arr.append(i)
                n += 1
            else:
                arr.append(temp)
                temp = i
                n = 1
            if n==k:
                return temp



arr = [2,1,3,5,4,6,7]
k = 2
res = Solution().getWinner(arr, k)
print(res)