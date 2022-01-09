class Solution:
    def countTriples(self, n: int) -> int:
        res = 0
        for a in range(1, n + 1):
            for b in range(1, n + 1):
                c = sqrt(a ** 2 + b ** 2)
                if not (c % 1) and int(c) <= n and c ** 2 == a ** 2 + b ** 2:
                    res += 1
        return res
