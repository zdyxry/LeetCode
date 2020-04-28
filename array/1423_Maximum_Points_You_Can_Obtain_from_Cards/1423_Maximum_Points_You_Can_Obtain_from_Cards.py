from typing import List


class Solution:
    def maxScore(self, cardPoints: List[int], k: int) -> int:
        size = len(cardPoints) - k
        minSubArraySum = float('inf')
        j = curr = 0

        for i, v in enumerate(cardPoints):
            curr += v
            if i - j + 1 > size:
                curr -= cardPoints[j]
            if i - j + 1 == size:
                minSubArraySum = min(minSubArraySum, curr)
        
        return sum(cardPoints) - minSubArraySum

cardPoints = [1,2,3,4,5,6,1]
k = 3
res = Solution().maxScore(cardPoints, k)
print(res)