class Solution:
    def oddCells(self, n: int, m: int, indices: List[List[int]]) -> int:
        rows = [0] * n
        cols = [0] * m
        for x, y in indices:
            rows[x] += 1
            cols[y] += 1
        return sum((rows[x] + cols[y]) % 2 == 1 for x in range(n) for y in range(m))
