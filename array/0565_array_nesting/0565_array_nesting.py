from typing import List


class Solution:
    def arrayNesting(self, nums: List[int]) -> int:
        n=len(nums)
        visited=[0]*n
        res=0
        for i in range(n):
            cnt=0
            while not visited[i]:
                visited[i]=1
                cnt+=1
                i=nums[i]
                
            res=max(res,cnt)
        return res


nums = [5,4,0,3,1,6,2]
res = Solution().arrayNesting(nums)
print(res)