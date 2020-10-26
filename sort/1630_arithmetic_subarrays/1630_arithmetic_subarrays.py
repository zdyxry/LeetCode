from typing import List

class Solution:
    def checkArithmeticSubarrays(self, nums: List[int], l: List[int], r: List[int]) -> List[bool]:
        li = []
        for i,j in zip(l,r):
            n = nums[i:j+1] 
            n.sort()
            for i in range(2, len(n)):
                if n[i] - n[i-1] != n[1] - n[0]: 
                    li.append(False) 
                    break 
            else: 
                li.append(True)
        return li


nums = [4,6,5,9,3,7]
l = [0,0,2]
r = [2,3,5]
res = Solution().checkArithmeticSubarrays(nums, l, r)
print(res)