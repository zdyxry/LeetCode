
class Solution(object):
    def validateStackSequences(self, pushed, poped):
        i = 0
        s = []
        for v in pushed:
            s.append(v)
            while s and i < len(poped) and s[-1] == poped[i]:
                s.pop()
                i += 1
        return i == len(poped)


pushed = [1,2,3,4,5]
poped = [4,5,3,2,1]
res = Solution().validateStackSequences(pushed, poped)
print(res)