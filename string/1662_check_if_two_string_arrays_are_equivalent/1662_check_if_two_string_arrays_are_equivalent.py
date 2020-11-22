from typing import List


class Solution:
    def arrayStringsAreEqual(self, word1: List[str], word2: List[str]) -> bool:
        return ''.join(word1) == ''.join(word2)


word1 = ["ab", "c"]
word2 = ["a", "bc"]
res = Solution().arrayStringsAreEqual(word1, word2)
print(res)