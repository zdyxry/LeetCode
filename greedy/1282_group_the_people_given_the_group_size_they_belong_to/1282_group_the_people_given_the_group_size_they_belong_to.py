import collections

from typing import List

class Solution:
    def groupThePeople(self, groupSizes: List[int]) -> List[List[int]]:
        groups = collections.defaultdict(list)
        for i, _id in enumerate(groupSizes):
            groups[_id].append(i)
        
        ans = list()
        for gsize, users in groups.items():
            for it in range(0, len(users), gsize):
                ans.append(users[it : it + gsize])
        return ans


groupSizes = [3,3,3,3,3,1,3]
res = Solution().groupThePeople(groupSizes)
print(res)