class Solution:
    def removeDuplicates(self, s: str, k: int) -> str:
        stack = []
        for i in s:
            print(i)
            if len(stack) == 0 or stack[-1][0] != i:
                stack.append([i, 1])
            else:
                stack[-1][1] += 1
                if stack[-1][1] == k:
                    stack.pop()
                    
        ret = ""
        for t in stack:
            ret += t[0] * t[1]
        return ret


s = "abcd"
k = 2
res = Solution().removeDuplicates(s, k)
print(res)