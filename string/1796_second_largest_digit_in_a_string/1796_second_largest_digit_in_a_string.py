from typing import List
import heapq


class Solution:
    def secondHighest(self, s: str) -> int:
        q = []
        for ch in s:
            if ch.isdigit():
                n = int(ch)
                if not q or -n != q[0]:
                    heapq.heappush(q, -n)
        if not q: return -1
        heapq.heappop(q)
        return -heapq.heappop(q) if q else -1


s = "dfa12321afd"
res = Solution().secondHighest(s)
print(res)