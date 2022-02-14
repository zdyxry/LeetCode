class Solution:
    def maxDistance(self, colors: List[int]) -> int:
        result = 0
        for i in range(len(colors)):
            for j in range(i, len(colors)):
                if colors[i] != colors[j]:
                    result = max(result, j-i)
        return result
