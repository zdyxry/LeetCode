
class Solution:
    def maxRepeating(self, sequence: str, word: str) -> int:
        ans = 1
        while sequence.find(word * ans) >= 0:
            ans += 1
        return ans - 1


sequence = "ababc"
word = "ab"
res = Solution().maxRepeating(sequence, word)
print(res)