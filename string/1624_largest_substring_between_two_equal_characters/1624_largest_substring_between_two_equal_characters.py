

class Solution:
    def maxLengthBetweenEqualCharacters(self, s: str) -> int:
        left = {}
        right = {}
        res = -1
        for i, c in enumerate(s):
            if c not in left:
                left[c] = i
            else:
                right[c] = i
                res = max(res, right[c] - left[c]-1)
        return res



s = "aa"
res = Solution().maxLengthBetweenEqualCharacters(s)
print(res)