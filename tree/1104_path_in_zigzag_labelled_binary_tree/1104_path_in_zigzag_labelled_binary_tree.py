from typing import List

class Solution:
    def pathInZigZagTree(self, label: int) -> List[int]:
        layer = 0
        for i in range(1, 33):
            if (1 << i) - 1 >= label:
                layer = i
                break

        cur = label
        ans = []
        while cur > 0:
            ans.append(cur)

            if layer % 2 == 1:
                offset = (cur - (1 << (layer - 1))) // 2
                layer -= 1
                cur = (1 << layer) - 1 - offset

            else:
                offset = ((1 << layer)-1 - cur) // 2
                layer -= 1
                cur = (1 << (layer - 1)) + offset

        ans.reverse()
        return ans


label = 14
res = Solution().pathInZigZagTree(label)
print(res)