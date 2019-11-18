import collections


class Solution(object):
    def customSortString1(self, S, T):
        l = []
        for i in S:
            l.append(i*T.count(i))
        for i in T:
            if i not in S:
                l.append(i)
        return ''.join(l)

    def customSortString2(self, S, T):
        counter, s = collections.Counter(T), set(S)
        result = [c*counter[c] for c in S]
        result.extend([c*counter for c, counter in counter.iteritems() if c not in s])
        return "".join(result)

S = "cba"
T = "abcd"
res = Solution().customSortString2(S, T)
print(res)