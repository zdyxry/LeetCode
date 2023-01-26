class Solution:
    def getSmallestString(self, n: int, k: int) -> str:
        result = ['a'] * n
        i, d = n-1, k-n
        while d > 25 :
            result[i] = 'z'
            i -= 1
            d -= 25
        result[i] = chr(ord(result[i]) + d )
        return ''.join(result)