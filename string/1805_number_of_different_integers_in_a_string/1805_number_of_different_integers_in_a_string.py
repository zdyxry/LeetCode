class Solution:
    def numDifferentIntegers(self, word: str) -> int:
        ret = set()
        stack = []
        for i in word:
            if i.isdigit():
                if len(stack) == 1 and stack[-1] == '0':
                    stack.pop()
                stack.append(i)
            else:
                if stack:
                    ret.add(''.join(stack))
                stack = []
        if stack:
            ret.add(''.join(stack))
        return len(ret)



word = "a123bc34d8ef34"
res = Solution().numDifferentIntegers(word)
print(res)