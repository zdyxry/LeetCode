class Solution:
    def checkIfCanBreak(self, s1: str, s2: str) -> bool:
        s1= sorted(s1)
        s2= sorted(s2)
        return all(s1[i] >= s2[i] for i in range(len(s1))) or all(s1[i] <= s2[i] for i in range(len(s1)))


s1 = "abc"
s2 = "efg"
res = Solution().checkIfCanBreak(s1, s2)
print(res)