
class Solution(object):
    def countSegments(self, s):
        result = int(len(s) and s[-1] != ' ')
        for i in xrange(1, len(s)):
            if s[i] == ' ' and s[i-1] != ' ':
                result += 1
        return result


s = "Hello, my name is John"
res = Solution().countSegments(s)
print(res)