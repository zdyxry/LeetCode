class Solution(object):
    def reverseWords(self, s):
        return ' '.join(reversed(s.split()))

s = "the sky is blue"
res = Solution().reverseWords(s)
print(res)