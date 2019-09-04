class Solution(object):
    def intersect(self, nums1, nums2):
        hash_table = {}
        res = []

        for i in nums1:
            if i in hash_table:
                hash_table[i] += 1
            else:
                hash_table[i] = 1
        
        for i in nums2:
            if i in hash_table:
                if hash_table[i] > 0:
                    res.append(i)
                    hash_table[i] -=1

        return res

nums1 = [1,2,2,1]
nums2 = [2,2]

res = Solution().intersect(nums1, nums2)
print(res)