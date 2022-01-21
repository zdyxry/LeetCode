class Solution:
    def isPrefixString(self, s: str, words: List[str]) -> bool:
        comp = ""
        for w in words:
            comp += w
            if comp == s:
                return True
        return False
