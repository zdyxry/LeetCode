class Solution:
    def countPrefixes(self, words: List[str], s: str) -> int:
        result = 0
        for w in words:
            if s.startswith(w):
                result += 1
        return result