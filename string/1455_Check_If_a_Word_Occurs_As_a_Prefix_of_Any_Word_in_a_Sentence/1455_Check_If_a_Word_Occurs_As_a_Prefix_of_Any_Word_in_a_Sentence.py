
class Solution:
    def isPrefixOfWord(self, sentence: str, searchWord: str) -> int:
        for idx, w in enumerate(sentence.split(' ')):
            if w.startswith(searchWord):
                return idx + 1
        else:
            return -1
            

sentence = "i love eating burger"
searchWord = "burg"
res = Solution().isPrefixOfWord(sentence, searchWord)
print(res)