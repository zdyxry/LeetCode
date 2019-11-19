import itertools

class Solution(object):
    def ambiguousCoordinates(self, S):
        S = S[1:-1]
        res = []
        def f(S):
            if not S or (len(S) > 1 and S[0] == "0" and S[-1] == "0"):
                return []
            if S == "0":
                return [S[0]]
            if S[0] == "0" and len(S) > 1:
                return [S[0] + "." + S[1:]]
            if S[-1] == "0" and len(S) > 1:
                return [S]
            return [S] + [S[:i] + "." + S[i:] for i in range(1, len(S))]

        for i in range(1, len(S)):
            for x, y in itertools.product(f(S[:i]), f(S[i:])):
                res.append("(%s, %s)" % (x, y))
        return res


S = "(123)"
res = Solution().ambiguousCoordinates(S)
print(res)