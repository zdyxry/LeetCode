class Solution:
    def canBeTypedWords(self, text: str, brokenLetters: str) -> int:
        result = 0
        for w in text.split(' '):
            if not any([ c in w for c in brokenLetters]):
                result += 1
        return result
                