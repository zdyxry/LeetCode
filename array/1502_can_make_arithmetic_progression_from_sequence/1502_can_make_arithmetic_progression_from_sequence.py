from typing import List


class Solution:
    def canMakeArithmeticProgression(self, arr: List[int]) -> bool:
        if len(arr) == 2:
            return True
        arr.sort(reverse=True)
        diff = arr[0] - arr[1]
        pre = arr[1]
        for i in arr[2:]:
            if pre - i != diff:
                return False
            pre = i
        return True


arr = [1,2,4]
res = Solution().canMakeArithmeticProgression(arr)
print(res)