from typing import List

class Solution:
    def slowestKey(self, releaseTimes: List[int], keysPressed: str) -> str:
        n = len(releaseTimes)
        last = longest = releaseTimes[0]
        ans = keysPressed[0]
        for i in range(1, n):
            last_time = releaseTimes[i] - last
            if last_time > longest:
                longest = last_time
                ans = keysPressed[i]
            elif last_time == longest:
                ans = max(ans, keysPressed[i])
            last = releaseTimes[i]
        return ans


releaseTimes = [9,29,49,50]
keysPressed = "cbcd"
res = Solution().slowestKey(releaseTimes, keysPressed)
print(res)