from typing import List

class Solution:
    def maxArea(self, h: int, w: int, horizontalCuts: List[int], verticalCuts: List[int]) -> int:
        height, width = 0, 0
        mod = 10 ** 9 + 7
        horizontalCuts.sort()
        verticalCuts.sort()
        h_len, v_len = len(horizontalCuts), len(verticalCuts)

        for i in range(1, h_len):
            height = max(height, horizontalCuts[i] - horizontalCuts[i - 1])
        height = max(height, horizontalCuts[0], h - horizontalCuts[h_len - 1])

        for i in range(1, v_len):
            width = max(width, verticalCuts[i] - verticalCuts[i - 1])
        width = max(width, verticalCuts[0], w - verticalCuts[v_len - 1])

        return (height * width) % mod


h = 5
w = 4
hor = [1,2,4]
ver = [1,3]
res = Solution().maxArea(h, w, hor, ver)
print(res)