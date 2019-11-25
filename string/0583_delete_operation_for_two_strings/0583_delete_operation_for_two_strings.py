
class Solution(object):
    def minDistance(self, word1, word2):
        len1, len2 = len(word1), len(word2)

        matrix = [[0]*(len2+1) for i in range(len1+1)]

        for i in range(1, len1+1):
            for j in range(1, len2+1):
                if word1[i-1] == word2[j-1]:
                    matrix[i][j] = matrix[i-1][j-1]+1
                else:
                    matrix[i][j] = max(matrix[i][j-1], matrix[i-1][j])
        
        return len1+ len2 - 2 * (matrix[-1][-1])


word1 = "sea"
word2 = "eat"

res = Solution().minDistance(word1, word2)
print(res)