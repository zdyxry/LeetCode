
class Solution(object):
    def alphabetBoardPath(self, target: str) -> str :
        m = {c : [i // 5, i % 5] for i, c in enumerate("abcdefghijklmnopqrstuvwxyz")}
        print(m)
        x0, y0 = 0, 0
        res = []
        for c in target:
            x, y = m[c]
            if y < y0:
                res.append('L' * (y0-y))
            if x < x0:
                res.append('U' * (x0 - x))
            if x > x0:
                res.append('D' * (x - x0))
            if y > y0:
                res.append('R' * (y - y0))
            res.append('!')
            x0, y0 = x, y
        return "".join(res)


target = "leet"
res = Solution().alphabetBoardPath(target)
print(res)