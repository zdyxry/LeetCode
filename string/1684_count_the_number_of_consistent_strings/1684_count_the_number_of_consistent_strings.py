from typing import List

class Solution:
    def countConsistentStrings(self, allowed: str, words: List[str]) -> int:
        allowed_set = set(allowed)
        return sum(set(word) <= allowed_set for word in words)




allowed = "ab"
words = ["ad","bd","aaab","baa","badab"]
res = Solution().countConsistentStrings(allowed, words)
print(res)