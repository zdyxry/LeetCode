from typing import List


class Solution:
    def destCity(self, paths: List[List[str]]) -> str:
        a = set()
        b = set()
        for path in paths:
            x, y = path
            a.add(x)
            b.add(y)

        b = b-a
        assert len(b) == 1
        return b.pop()



paths = [["London","New York"],["New York","Lima"],["Lima","Sao Paulo"]]
res = Solution().destCity(paths)
print(res)