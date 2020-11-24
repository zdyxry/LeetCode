from typing import List


class Solution:
    def largestValsFromLabels(self, values: List[int], labels: List[int], num_wanted: int, use_limit: int) -> int:
        l = [(val, label) for val, label in zip(values, labels)]
        l.sort(reverse = True)

        c = Counter()
        cnt = 0
        sum = 0
        for val, label in l:
            if cnt == num_wanted:
                break

            if c[label] < use_limit:
                c[label] += 1
                sum += val
                cnt += 1

        return sum

