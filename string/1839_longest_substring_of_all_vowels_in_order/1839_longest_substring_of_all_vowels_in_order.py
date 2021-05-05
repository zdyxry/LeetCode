class Solution:
    def longestBeautifulSubstring(self, word: str) -> int:
        n = len(word)
        if n < 5:
            return 0 
        l, r = 0, 1
        res = 0
        curr = word[0]
        ss = set()
        ss.add(word[0])
        while r < n:
            if word[r] < curr:
                if len(ss) == 5:
                    res = max(res, r-l)
                l = r
                ss.clear()
            curr = word[r]
            ss.add(word[r])
            r += 1
        if len(ss) == 5:
            res = max(res, r - l)
        return res

word = "aeiaaioaaaaeiiiiouuuooaauuaeiu"
res = Solution().longestBeautifulSubstring(word)
print(res)