from typing import List


class Solution:
    def countTriplets(self, arr: List[int]) -> int:
        if len(arr) < 2:
            return 0

        count = 0

        for i in range(len(arr)-1):
            temp = arr[i]
            for j in range(i+1, len(arr)):
                temp = temp^arr[j]
                if temp == 0:
                    count += j-i

        return count



arr = [2,3,1,6,7]
res = Solution().countTriplets(arr)
print(res)