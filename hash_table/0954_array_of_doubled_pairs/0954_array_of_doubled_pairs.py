from collections import Counter


class Solution(object):
    def canReorderDoubled(self, A):
        count = Counter(A)
        for x in sorted(count, key=abs):
            if count[x] > count[2 * x]:
                return False
        return True


A = [3,1,3,6]
res = Solution().canReorderDoubled(A)
print(res)
