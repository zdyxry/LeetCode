
class Solution(object):
    def findTheLongestSubstring(self, s):
        vowels = {'a':1, 'e':2, 'i':4, 'o':8, 'u':16}
        d, n , r = {0:-1}, 0, 0
        for i, c in enumerate(s):
            if c in vowels:
                n ^= vowels[c]
            if n not in d:
                d[n] = i
            else:
                r = max(r, i - d[n])
            # print c, bin(n), d, r
        return r


s = "eleetminicoworoep"
res = Solution().findTheLongestSubstring(s)
print(res)