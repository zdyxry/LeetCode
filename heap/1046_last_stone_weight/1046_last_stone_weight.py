
import heapq 

class Solution(object):
    def lastStoneWeight(self, stones):
        max_heap = [-x for x in stones]
        heapq.heapify(max_heap)
        for i in xrange(len(stones)-1):
            x, y = -heapq.heappop(max_heap), -heapq.heappop(max_heap)
            heapq.heappush(max_heap, -abs(x-y))
        return -max_heap[0]


stones = [2,7,4,1,8,1]
res = Solution().lastStoneWeight(stones)
print(res)