import collections

class Solution:
    def minDeletions(self, s: str) -> int:
        c = collections.Counter(s)
        seen = set()
        res = 0
        for k, v in c.items():
            cur = v
            while cur>0 and cur in seen:
                cur -= 1
            res += v - cur
            if cur>0:
                seen.add(cur)
        return res


s = "aab"
res = Solution().minDeletions(s)
print(res)