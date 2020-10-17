from typing import List



class Solution:
    def grayCode(self, n: int) -> List[int]:
        def dfs():
            nonlocal n, code, gray
            if code in visited:
                return
            visited.add(code)
            gray.append(code)
            for i in range(0, n):
                code ^= 1 << i
                dfs()
                code ^= 1 << i

        code = 0
        visited = set()
        gray = []
        dfs()
        return gray



n = 2
res = Solution().grayCode(n)
print(res)