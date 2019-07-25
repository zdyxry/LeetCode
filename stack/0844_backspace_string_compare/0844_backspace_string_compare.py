class Solution(object):
    def backspaceCompare(self, S1, S2):
        r1 = len(S1) - 1 
        r2 = len (S2) - 1
        
        while r1 >= 0 or r2 >= 0:
            char1 = char2 = ""
            if r1 >= 0:
                char1, r1 = self.getChar(S1, r1)
            if r2 >= 0:
                char2, r2 = self.getChar(S2, r2)
            print(char1, r1, char2, r2)
            if char1 != char2:
                return False
        return True
        
    
    def getChar(self, s , r):
        char, count = '', 0
        while r >= 0 and not char:
            if s[r] == '#':
                count += 1
            elif count == 0:
                char = s[r]
            else:
                count -= 1
            r -= 1
        return char, r


S = "ab#c"
T = "add##c"
print(Solution().backspaceCompare(S, T))