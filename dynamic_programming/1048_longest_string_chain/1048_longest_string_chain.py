from typing import List

class Solution:
    def longestStrChain(self, words: List[str]) -> int:
        words.sort(key=len)
        note = {}
        maxChain = 1
        for word in words:
            if word not in note:
                note[word] = 1
            for i in range(0, len(word)):
                newWord = word[:i] + word[i+1:]
                if newWord in note:
                    note[word] = max(note[word], note[newWord]+1)
            maxChain = max(maxChain, note[word])
        return maxChain


words = ["a","b","ba","bca","bda","bdca"]
res = Solution().longestStrChain(words)
print(res)