class Solution:
    def minInsertions(self, s: str) -> int:
        # use stack
        stack = []
        count = 0
        i=0
        while i < len(s):
            if s[i] == '(':
                stack.append(s[i])
                i += 1
            else:
                # when facing empty stack
                if not stack:
                    if i+1 < len(s) and s[i+1] == ")":
                        count += 1  # add one "("
                        i += 2
                    else:
                        count += 2   # add one "("  and one ")"
                        i += 1
                else:
                    # check two positions
                    if i + 1 < len(s) and s[i+1] == ")":
                        stack.pop()
                        i += 2
                    else:
                        count += 1   # add one ")"
                        stack.pop()
                        i += 1
        
        rest = len(stack)*2  # still have "(" on the stack. one "(" pairs with two ")"
        return count + rest


s = "(()))"
res = Solution().minInsertions(s)
print(res)