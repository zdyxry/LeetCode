class Solution:
    def reversePrefix(self, word: str, ch: str) -> str:
        result = ""
        index = len(word)
        for idx, c in enumerate(word):
            if c == ch:
                index = idx
                break
        if index == len(word):
            return word
        
        for i in range(index, -1, -1):
            result += word[i]
        return result + word[index+1:]