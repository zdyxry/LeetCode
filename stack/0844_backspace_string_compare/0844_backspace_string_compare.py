class Solution(object):
    def backspaceCompare(self, S, T):
        Sstk = []
        Tstk = []
        for i in S:
            if i != "#":
                Sstk.append(i)
            elif Sstk:
                Sstk.pop()

        for j in T:
            if j != "#":
                Tstk.append(j)
            elif Tstk:
                Tstk.pop()
        
        return Sstk == Tstk


S = "ab#c"
T = "add##c"
print(Solution().backspaceCompare(S, T))