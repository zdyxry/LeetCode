from typing import List


class Solution:
    def minOperations(self, logs: List[str]) -> int:
        count = 0
        for log in logs:
            if log == '../':
                if count > 0:
                    count -= 1
                else:
                    pass
            elif log == './':
                pass
            else:
                count += 1
        return count


logs = ["d1/","d2/","../","d21/","./"]
res = Solution().minOperations(logs)
print(res)