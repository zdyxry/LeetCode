from typing import List

class Solution:
    def simplifiedFractions(self, n: int) -> List[str]:
        res = []
        v = set()
        for fenmu in range(2, n + 1):
            for fenzi in range(1, fenmu):
                value = fenzi / fenmu
                if value not in v:
                    v.add(value)
                    res.append(str(fenzi) + '/' + str(fenmu))
        return res


n = 3
res = Solution().simplifiedFractions(n)
print(res)