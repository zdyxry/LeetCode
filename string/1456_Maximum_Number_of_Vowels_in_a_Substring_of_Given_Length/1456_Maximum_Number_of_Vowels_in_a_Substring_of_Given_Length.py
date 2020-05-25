
class Solution:
    def maxVowels(self, s: str, k: int) -> int:
        d = {'a': 0, 'e': 0, 'i': 0, 'o': 0, 'u': 0}
        for c in s[:k]:
            if c in d:
                d[c] += 1
        res = sum(d.values())
        for i in range(k, len(s)):
            if s[i-k] in d:
                d[s[i-k]] -= 1
            if s[i] in d:
                d[s[i]] += 1
            res = max(res, sum(d.values()))
        return res


s = "abciiidef"
k = 3
res = Solution().maxVowels(s, k)
print(res)