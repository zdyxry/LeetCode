class Solution:
    def minimumMoves(self, s: str) -> int:
        l = len(s)
        i = 0
        result = 0
        while i < l:
            if s[i] == 'X':
                result += 1
                i += 3
            else:
                i += 1
        return result
