class Solution(object):
    def singleNumber(self, A):
        one, two, carry = 0, 0, 0
        for x in A:
            two |= one & x
            one ^= x
            carry = one & two
            one &= ~carry
            two &= ~carry
        return one

A = [2,2,3,2]
res = Solution().singleNumber(A)
print(res)