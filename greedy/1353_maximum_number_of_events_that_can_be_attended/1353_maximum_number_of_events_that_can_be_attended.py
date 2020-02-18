import heapq

class Solution(object):
    def maxEvents(self, events):
        ans = 0
        end = list()
        events = sorted(events, reverse=True)
        print events
        for i in range(1, 100010, 1):
            while events and events[-1][0] == i:
                heapq.heappush(end, events.pop()[1])
            while end and end[0] < i:
                heapq.heappop(end)
            if end:
                heapq.heappop(end)
                ans += 1
        return ans

events = [[1,2],[2,3],[3,4]]
res = Solution().maxEvents(events)
print(res)