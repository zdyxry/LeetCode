
class Solution(object):
    def decodeString(self, s):
        stack = []
        curNum = 0
        curString = ''

        for c in s:
            if c == '[':
                stack.append(curString)
                stack.append(curNum)
                curString = ''
                curNum = 0
            elif c == ']':
                num = stack.pop()
                prevString = stack.pop()
                curString = prevString + num*curString
            elif c.isdigit():
                curNum = curNum * 10 + int(c)
            else:
                curString += c
            print c, stack,curString
        return curString

s = "3[a]2[bc]"
res = Solution().decodeString(s)
print(res)