class Solution(object):
    def scoreOfParentheses(self, S):
        result, depth = 0, 0
        for i in xrange(len(S)):
            if S[i] == '(':
                depth += 1
            else:
                depth -= 1
                if S[i-1] == '(':
                    result += 2 ** depth
            print S[i], depth,result
        return result

S = "(()(()))"
res = Solution().scoreOfParentheses(S)
print(res)