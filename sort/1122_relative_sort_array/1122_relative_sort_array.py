from collections import Counter

class Solution(object):
    def relativeSortArray(self, arr1, arr2):
        c = Counter(arr1)
        res = []
        for i in arr2:
            res += [i] * c.pop(i)
        return res + sorted(c.elements())


arr1 = [2,3,1,3,2,4,6,7,9,2,19]
arr2 = [2,1,4,3,9,6]
res = Solution().relativeSortArray(arr1, arr2)
print(res)