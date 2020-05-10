from typing import List


class Solution:
    def buildArray(self, target: List[int], n: int) -> List[str]:
        res = []
        keep = 0
        for num in range(1, n+1):
            res.append("Push")
            keep += 1
            if num not in target:
                res.append("Pop")
                keep -= 1
            if keep == len(target):
                break
        return res


target = [1,3]
n = 3
res = Solution().buildArray(target, n)
print(res)