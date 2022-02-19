class Solution:
    def countPoints(self, rings: str) -> int:
        m = {}
        for i in range(10):
            m[i] = set()

        for j in range(0, len(rings), 2):
            s = rings[j:j+2]
            color, idx = s[0], s[1]
            m[int(idx)].add(color)

        result = 0
        for i in m:
            if len(m[i]) == 3:
                result += 1
        return result