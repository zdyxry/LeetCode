class Solution(object):
    def reverseOnlyLetters(self, S):
        """
        :type S: str
        :rtype: str
        """
        letter = 'abcdefghijklmnopqrstuvwxyz'
        s = list(S)
        i,j = 0, len(s)-1

        while i < j: 
            if s[i].lower() not in letter: 
                i = i + 1
            elif s[j].lower() not in letter: 
                j = j - 1
            else:
                s[i],s[j] = s[j], s[i]
                i = i+1
                j = j-1
                
        return ''.join(s)


    def reverseOnlyLetters2(self, S):
        """
        :type S: str
        :rtype: str
        """
        def getNext(S):
            for i in reversed(xrange(len(S))):
                if S[i].isalpha():
                    yield S[i]

        result = []
        letter = getNext(S)
        for i in xrange(len(S)):
            if S[i].isalpha():
                result.append(letter.next())
            else:
                result.append(S[i])
        return "".join(result)




print(Solution().reverseOnlyLetters("ab-cd"))
print(Solution().reverseOnlyLetters2("ab-cd"))