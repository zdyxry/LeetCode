
class Solution:
    def rotateString(self, A: str, B: str) -> bool:
        if len(A) != len(B):
            return False
        if A == B:
            return True
        for i in range(1, len(A)):
            if A[i:] + A[:i] == B:
                return True
        return False


A = "abcde"
B = "cdeab"
res = Solution().rotateString(A, B)
print(res)