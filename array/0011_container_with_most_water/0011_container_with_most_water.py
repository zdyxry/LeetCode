class Solution(object):
    def maxArea(self, height):
        max_area, i, j = 0, 0, len(height) - 1
        while i < j:
            max_area = max(max_area, min(height[i], height[j]) * (j - i))
            if height[i] < height[j]:
                i += 1
            else:
                j -= 1
        return max_area

height = [1,8,6,2,5,4,8,3,7]

res = Solution().maxArea(height)
print(res)