import collections
import heapq


class Solution(object):
    def reorganizeString(self, S):
        counts = collections.Counter(S)
        if any(v > (len(S) + 1) / 2 for k, v in counts.iteritems()):
            return ""

        result = []
        max_heap = []
        for k, v in counts.iteritems():
            heapq.heappush(max_heap, (-v, k))
        print([i for i in counts.iteritems()])
        while len(max_heap) > 1:
            count1, c1 = heapq.heappop(max_heap)
            count2, c2 = heapq.heappop(max_heap)
            print(count1, c1, count2, c2)
            if not result or c1 != result[-1]:
                result.extend([c1, c2])
                if count1 + 1: 
                    heapq.heappush(max_heap, (count1 + 1, c1))
                if count2 + 1: 
                    heapq.heappush(max_heap, (count2 + 1, c2))
        return "".join(result) + (max_heap[0][1] if max_heap else '')


S = "aab"
res = Solution().reorganizeString(S)
print(res)