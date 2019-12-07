class Solution(object):
    def nextGreaterElement(self, nums1, nums2):
        """
        :type nums1: List[int]
        :type nums2: List[int]
        :rtype: List[int]
        """
        st, d = [], {}
        for n in nums2:
            while st and st[-1] < n:
                d[st.pop()] = n
            st.append(n)
        return [d.get(x, -1) for x in nums1]




nums1 = [4,1,2]
nums2 = [1,3,4,2]
res = Solution().nextGreaterElement(nums1, nums2)
print(res)