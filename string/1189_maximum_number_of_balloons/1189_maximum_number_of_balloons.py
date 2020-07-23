import collections


class Solution:
    def maxNumberOfBalloons(self, text: str) -> int:
        d = collections.Counter(text)
        return min(d['b'], d['a'], d['l'] // 2, d['o'] // 2, d['n'])


text = "nlaebolko"
res = Solution().maxNumberOfBalloons(text)
print(res)