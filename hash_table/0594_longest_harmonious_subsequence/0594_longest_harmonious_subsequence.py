class Solution(object):
    def findLHS(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        if not nums:
            return 0
        
        d = {}
        for n in nums:
            if n in d:
                d[n] += 1
            else:
                d[n] = 1

        max_length = 0
        for n in d:
            if n+1 in d:
                max_length = max(max_length, d[n + 1] + d[n])
            
        return max_length

nums = [1,3,2,2,5,2,3,7]

print(Solution().findLHS(nums))