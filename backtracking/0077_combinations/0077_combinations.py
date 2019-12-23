
class Solution(object):
    def combine(self, n, k):
        result, combination = [], []
        i = 1
        while True:
            if len(combination) == k:
                result.append(combination[:])
            if len(combination) == k or \
                len(combination)+(n-i+1) < k:
                if not combination:
                    break
                i = combination.pop() +1
            else:
                combination.append(i)
                i += 1
        return result


n = 4
k = 2
res = Solution().combine(n,k)
print(res)