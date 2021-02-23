from typing import List


class Solution:
    def minOperations(self, boxes: str) -> List[int]:
        ans = []
        left = right = total = 0
        n = len(boxes)
        for i in range(n):
            if boxes[i] == '1':
                right += 1
                total += i
        print(right, total)
        for i in range(n):
            ans.append(total)
            if boxes[i] == '1':
                left += 1
                right -= 1
            total += left - right
        return ans


boxes = "110"
res = Solution().minOperations(boxes)
print(res)