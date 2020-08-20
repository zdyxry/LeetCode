from typing import List


class Solution:
    def reconstructQueue(self, people: List[List[int]]) -> List[List[int]]:
        people.sort(key = lambda x: (-x[0], x[1]))
        output = []
        for p in people:
            output.insert(p[1], p)
        return output


people = [[7,0],[4,4],[7,1],[5,0],[6,1],[5,2]]
res = Solution().reconstructQueue(people)
print(res)