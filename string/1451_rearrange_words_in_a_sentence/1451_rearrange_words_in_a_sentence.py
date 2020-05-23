import collections

class Solution:
    def arrangeWords(self, text: str) -> str:
        a = text.split()
        a[0] = a[0].lower()
        d = collections.defaultdict(list)
        for s in a:
            d[len(s)].append(s)
        res = []
        for k in sorted(d.keys()):
            res.extend(d[k])
        res[0] = res[0].title()
        return ' '.join(res)


text = "Leetcode is cool"
res = Solution().arrangeWords(text)
print(res)