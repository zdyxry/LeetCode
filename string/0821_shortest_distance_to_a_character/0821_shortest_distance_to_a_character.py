class Solution(object):
    def shortestToChar(self, S, C):
        prev = float('-inf')
        ans = []
        for i, x in enumerate(S):
            if x == C: prev = i
            ans.append(i - prev)

        prev = float('inf')
        for i in range(len(S) - 1, -1, -1):
            if S[i] == C: prev = i
            ans[i] = min(ans[i], prev - i)

        return ans



S = "loveleetcode"
C = "e"
res = Solution().shortestToChar(S, C)
print(res)