class Solution(object):
    def detectCapitalUse(self, word):
        """
        :type word: str
        :rtype: bool
        """
        count = 0
        for i in range(len(word)):
            if word[i].isupper():
                count+=1
            
        if count == 1:
            if word[0].isupper():
                return True
            else:
                return False
        if count == len(word):
            return True
        elif count == 0:
            return True
        else:
            return False


print(Solution().detectCapitalUse("USA"))