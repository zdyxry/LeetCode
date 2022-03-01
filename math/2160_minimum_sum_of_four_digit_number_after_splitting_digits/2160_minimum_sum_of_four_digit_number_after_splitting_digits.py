class Solution:
    def minimumSum(self, num: int) -> int:
        digits = []
        while num:
            digits.append(num % 10)
            num //= 10
        digits.sort()
        return 10 * (digits[0] + digits[1]) + digits[2] + digits[3]
