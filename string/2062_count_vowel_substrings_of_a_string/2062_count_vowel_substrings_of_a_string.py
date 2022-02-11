class Solution:
    def countVowelSubstrings(self, word: str) -> int:
        a = {'a','e','i','o','u'}
        l, r = 0, 0
        cnt = 0
        while l < len(word):
            if word[r] in a:
                if len(set(word[l:r+1])) >= 5:
                    cnt += 1
                r += 1
                if r == len(word):
                    l += 1
                    r = l
            else:
                l += 1
                r = l
        return cnt
