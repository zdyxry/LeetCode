from typing import List


class Solution:
    def duplicateZeros(self, arr: List[int]) -> None:
        """
        Do not return anything, modify arr in-place instead.
        """
        nums = []
        for num in arr:
            nums.append(num)
            if num == 0:
                nums.append(0)
        for i in range(len(arr)):
            arr[i] = nums[i]


arr = [1,0,2,3,0,4,5,0]
Solution().duplicateZeros(arr)
print(arr)