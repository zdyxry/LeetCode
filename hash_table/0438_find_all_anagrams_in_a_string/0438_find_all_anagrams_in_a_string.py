class Solution(object):
    def findAnagrams(self, s, p):
        """
        :type s: str
        :type p: str
        :rtype: List[int]
        """
        l1=len(s)
        l2=len(p)
        rst=[]
        d1=[0]*26
        for c in p:
            d1[ord(c)-ord('a')]+=1
        #
        d2=[0]*26
        for c in s[:l2]:
            d2[ord(c)-ord('a')]+=1
        if d1==d2:
            rst.append(0)
        #
        for i in range(1,l1-l2+1):
            d2[ord(s[i-1])-ord('a')]-=1
            d2[ord(s[i+l2-1])-ord('a')]+=1
            if d1==d2:
                rst.append(i)
        return rst


s = "cbaebabacd"
p = "abc"
print(Solution().findAnagrams(s,p))