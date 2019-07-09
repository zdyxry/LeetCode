class Solution:
    def countBinarySubstrings(self, s):
        """
        :type s: str
        :rtype: int
        """
        result = 0
        prev_length = 0
        cur_length = 1
        for i in range(1, len(s)):
            if (s[i] == s[i - 1]):
                cur_length += 1
            else:
                prev_length = cur_length
                cur_length = 1
            if prev_length >= cur_length:
                result += 1
        return result

    def countBinarySubstrings2(self, s):
        """
        :type s: str
        :rtype: int
        """
        result, prev, curr = 0, 0, 1
        for i in xrange(1, len(s)):
            if s[i-1] != s[i]:
                result += min(prev, curr)
                prev, curr = curr, 1
            else:
                curr += 1
        result += min(prev, curr)
        return result

print(Solution().countBinarySubstrings('00110'))
print(Solution().countBinarySubstrings2('00110'))