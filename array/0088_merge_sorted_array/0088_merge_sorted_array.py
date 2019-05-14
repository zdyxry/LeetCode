# -*- coding: utf-8 -*-

class Solution(object):
    # @param nums1  a list of integers
    # @param m  an integer, length of nums1
    # @param nums2  a list of integers
    # @param n  an integer, length of nums2
    # @return nothing
    def merge(self, nums1, m, nums2, n):
        last, i, j = m + n - 1, m - 1, n - 1

        while i >= 0 and j >= 0:
            if nums1[i] > nums2[j]:
                nums1[last] = nums1[i]
                last, i = last - 1, i - 1
            else:
                nums1[last] = nums2[j]
                last, j = last - 1, j - 1
        print(i, j, last)
        while j >= 0:
            nums1[last] = nums2[j]
            last, j = last - 1, j - 1

nums1 = [1,2,3,0,0,0,0]
m = 3
nums2 = [2,5,6,7]
n = 4

nums3 = [0]
o = 0 
nums4 = [1]
p = 1

nums5 = [4,5,6,7,0,0,0,0]
a = 4
nums6 = [1,2,3,4,5]
b = 4

nums7 = [1,2,3,4,0,0,0]
c = 4
nums8 = [5,6,7]
d = 3
# Solution().merge(nums1,m,nums2,n)
# Solution().merge(nums3, o, nums4, p)
# Solution().merge(nums5, a, nums6, b)
Solution().merge(nums7, c, nums8, d)
print(nums7)
