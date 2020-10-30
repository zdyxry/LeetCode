from typing import List

class Solution:
    def letterCasePermutation(self, S: str) -> List[str]:
        ans = [[]]

        for char in S:
            n = len(ans)
            if char.isalpha():
                for i in range(n):
                    ans.append(ans[i][:])
                    ans[i].append(char.lower())
                    ans[n+i].append(char.upper())
            else:
                for i in range(n):
                    ans[i].append(char)

        return map("".join, ans)


S = "a1b2"
res = Solution().letterCasePermutation(S)
print(res)