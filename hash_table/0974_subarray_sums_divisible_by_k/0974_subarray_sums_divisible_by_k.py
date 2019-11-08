from collections import defaultdict

class Solution(object):
    def subarrayDivByK(self, A, K):
        count = defaultdict(int)
        count[0] = 1
        result, prefix = 0, 0
        for a in A:
            prefix = (prefix+a) % K
            result += count[prefix]
            count[prefix] += 1
        return result


A = [4,5,0,-2,-3,1]
K = 5

res = Solution().subarrayDivByK(A, K)
print(res)