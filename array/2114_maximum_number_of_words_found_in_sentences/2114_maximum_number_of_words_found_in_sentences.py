class Solution:
    def mostWordsFound(self, sentences: List[str]) -> int:
        result = 0
        for i in sentences:
            result = max(result, len(i.split()))
        return result