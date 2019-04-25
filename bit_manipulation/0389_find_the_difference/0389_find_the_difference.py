# -*- coding: utf-8 -*-
class Solution(object):
    def findTheDifference(self, s, t):
        """
        :type s: str
        :type t: str
        :rtype: str
        """
        nums=numt=0
        
        #take ascii value of both strings
        s=s.encode('ascii')
        t=t.encode('ascii')
        
        #calculate the sum of ascii value of both strings
        for s_ch in s:
            nums+=ord(s_ch)
        for t_ch in t:
            numt+=ord(t_ch)

        #return character for ascii value. obtained by differnce of both strings
        return (chr(abs(numt-nums)))

print(Solution().findTheDifference("abcd","abcdf"))