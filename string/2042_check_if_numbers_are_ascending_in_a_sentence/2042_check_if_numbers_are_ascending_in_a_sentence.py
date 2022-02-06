class Solution:
    def areNumbersAscending(self, s: str) -> bool:
        result = []
        for c in s.split(' '):
            if str.isdigit(c):
                result.append(int(c))

        for i in range(1, len(result)):
            if result[i] <= result[i-1]:
                return False
        return True