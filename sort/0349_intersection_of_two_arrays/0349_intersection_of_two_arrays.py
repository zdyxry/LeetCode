class Solution(object):
    def intersection(self, nums1, nums2):
        if len(nums1) > len(nums2):
            return self.intersection(nums2, nums1)

        lookup = set()
        for i in nums1:
            lookup.add(i)

        res = []
        for i in nums2:
            if i in lookup:
                res += i,
                lookup.discard(i)
        
        return res

nums1 = [1,2,2,1]
nums2 = [2,2]
res = Solution().intersection(nums1, nums2)
print(res)