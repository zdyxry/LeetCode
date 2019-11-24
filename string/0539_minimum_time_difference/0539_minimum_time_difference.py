
class Solution(object):
    def findMinDifference1(self, timePoints):
        if len(timePoints) > 1440:
            return 0
        s = sorted(map(lambda t: int(t[:2]) * 60 + int(t[3:]), timePoints))
        return min(s2 - s1 for s1,s2 in zip(s, s[1:] + [1440+s[0]]))

    def findMinDifference2(self, timePoints):
        if len(timePoints) > 1440:
            return 0
        buckets = [0] * 1440
        for tp in timePoints:
            seconds = int(tp[:2]) * 60 + int(tp[3:])
            buckets[seconds] += 1
            if buckets[seconds] > 1:
                return 0
        s = [i for i, cnt in enumerate(buckets) if cnt]
        return min(s2-s1 for s1, s2 in zip(s, s[1:] + [1440 + s[0]]))


timePoints = ["23:59","00:00"]
res1 = Solution().findMinDifference1(timePoints)
print(res1)
res2 = Solution().findMinDifference2(timePoints)
print(res2)