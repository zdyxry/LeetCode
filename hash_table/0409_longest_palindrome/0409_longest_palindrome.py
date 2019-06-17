class Solution(object):
    def longestPalindrome(self, s):
        """
        :type s: str
        :rtype: int
        """
        if s is None or len(s) == 0:
            return 0
        dic = dict()
        for c in s:
            if dic.get(c) is None:
                dic[c] = 1
            else:
                dic[c] += 1
        f = 0
        ans = 0
        for key, val in dic.items():
            if val % 2 == 0:
                ans += val
            elif val == 1:
                f = 1
            else:
                ans += val-1
                f = 1
        return ans+f


s = "abccccdd"
print(Solution().longestPalindrome(s))