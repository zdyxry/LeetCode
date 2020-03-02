
class Solution(object):
    def smallerNumbersTHanCurrent(self, nums):
        nums2 = sorted(nums)
        mapping = {}
        for i, num in enumerate(nums2):
            if i > 0 and nums2[i] == nums2[i-1]:
                mapping[nums2[i]] = mapping[nums2[i-1]]
            else:
                mapping[nums2[i]] = i
        res = []
        for num in nums:
            res.append(mapping[num])
        return res