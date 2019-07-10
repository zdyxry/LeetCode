class Solution(object):
    def buddyStrings(self, A, B):
        """
        :type A: str
        :type B: str
        :rtype: bool
        """
        if len(A) != len(B):
            return False
        A = list(A)
        diff = []
        for i in range(len(A)):
            if A[i] != B[i]:
                diff.append(i)
        if len(set(A)) < len(A) and len(diff) == 0:
            return True
        return len(diff) == 2 and A[diff[0]] == B[diff[1]] and A[diff[1]] == B[diff[0]]


A = "ab"
B = "ba"
print(Solution().buddyStrings(A, B))