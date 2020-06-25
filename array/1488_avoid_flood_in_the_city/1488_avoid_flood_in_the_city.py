from typing import List

class Solution:
    def avoidFlood(self, rains: List[int]) -> List[int]:
        ans = [1] * len(rains)
        left = []
        record = dict()

        for i, val in enumerate(rains):
            if val > 0:
                ans[i] = -1
                if val in record:
                    pos = bisect_left(left, record[val])
                    if pos > len(left):
                        return []
                    ans[left.pop(pos)] = val
                record[val] = i
            else:
                left.append(i)
        return ans


rains = [1,2,3,4]
res = Solution().avoidFlood(rains)
print(res)