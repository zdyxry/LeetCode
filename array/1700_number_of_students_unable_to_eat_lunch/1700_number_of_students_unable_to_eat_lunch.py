from typing import List

class Solution:
    def countStudents(self, students: List[int], sandwiches: List[int]) -> int:
        while students:
            if students[0] == sandwiches[0]:
                students = students[1:]
                sandwiches = sandwiches[1:]
            elif sandwiches[0] not in students:
                return len(students)
            else:
                students = students[1:] + [students[0]]
        return 0


students = [1,1,0,0]
sandwiches = [0,1,0,1]
res = Solution().countStudents(students,sandwiches)
print(res)