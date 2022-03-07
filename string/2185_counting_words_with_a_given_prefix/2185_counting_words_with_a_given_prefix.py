class Solution:
    def prefixCount(self, words: List[str], pref: str) -> int:
        result = 0
        for w in words:
            if w.startswith(pref):
                result += 1
        return result