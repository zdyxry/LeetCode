class Solution(object):
    def reverseString(self, s):
        """
        :type s: str
        :rtype: str
        """
        s[:]=s[::-1]


s = ["h","e","l","l","o"]
Solution().reverseString(s)
print(s)