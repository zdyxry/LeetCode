class Solution:
    def areOccurrencesEqual(self, s: str) -> bool:
        m = {}
        for c in s:
            if c in m:
                m[c] += 1
            else:
                m[c] = 0
        return len(set([x for x in m.values()])) == 1