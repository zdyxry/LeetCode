class Solution:
    def truncateSentence(self, s: str, k: int) -> str:
        if len(s.split(' ')) <= k:
            return s
        else:
            return ' '.join(s.split(' ')[:k])


s = "Hello how are you Contestant"
k = 4
res = Solution().truncateSentence(s, k)
print(res)