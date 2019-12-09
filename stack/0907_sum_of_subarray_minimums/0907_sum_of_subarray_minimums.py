
class Solution(object):
    def sumSubarrayMins(self, A):
        dot = 0
        sums = 0
        stack = []

        for elem in A:
            new_count = 1
            while stack and stack[-1][0] >= elem:
                val, count = stack.pop()
                dot -= val * count
                new_count += count
            
            stack.append((elem, new_count))
            dot += elem * new_count
            sums += dot

        return sums % (10**9+7)


A = [3,1,2,4]
res = Solution().sumSubarrayMins(A)
print(res)