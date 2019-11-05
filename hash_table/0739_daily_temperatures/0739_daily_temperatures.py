class Solution(object):
    def dailyTemperatures(self, T):
        result = [0] * len(T)
        stk = []
        for i in xrange(len(T)):
            while stk and T[stk[-1]] < T[i]:
                idx = stk.pop()
                result[idx] = i - idx
            stk.append(i)
            # print(stk, result)
        return result


T = [73,74,75,71,69,72,76,73]

res = Solution().dailyTemperatures(T)
print(res)