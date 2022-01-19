class Solution:
    def isThree(self, n: int) -> bool:
        result = 0
        for i in range(1,n+1):
            if n % i == 0 :
                result += 1
        return result == 3
        