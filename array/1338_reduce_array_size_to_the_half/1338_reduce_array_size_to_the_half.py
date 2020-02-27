import collections


class Solution(object):
    def minSetSize(self, A):
        res, temp , n = 0, 0, len(A)/ 2
        for a, freq in collections.Counter(A).most_common():
            temp += freq
            res += 1
            if temp >= n:
                return res
        return res


A = [3,3,3,3,5,5,5,2,2,7]
res = Solution().minSetSize(A)
print(res)