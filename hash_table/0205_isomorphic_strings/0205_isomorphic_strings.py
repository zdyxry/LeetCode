# -*- coding: utf-8 -*-


class Solution(object):
    def isIsomorphic(self, s, t):
        """
        :type s: str
        :type t: str
        :rtype: bool
        """
        a = [0] * 256
        b = [0] * 256
        for x in xrange(len(s)):
            if a[ord(s[x])] != b[ord(t[x])]:
                return False
            a[ord(s[x])] = x + 1
            b[ord(t[x])] = x + 1
        return True


s = 'egg'
t = 'add'

print(Solution().isIsomorphic(s,t))