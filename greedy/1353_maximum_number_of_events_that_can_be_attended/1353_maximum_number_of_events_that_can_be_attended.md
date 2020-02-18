1353. Maximum Number of Events That Can Be Attended


Medium


Given an array of events where events[i] = [startDayi, endDayi]. Every event i starts at startDayi and ends at endDayi.

You can attend an event i at any day d where startTimei <= d <= endTimei. Notice that you can only attend one event at any time d.

Return the maximum number of events you can attend.

 

Example 1:

```

Input: events = [[1,2],[2,3],[3,4]]
Output: 3
Explanation: You can attend all the three events.
One way to attend them all is as shown.
Attend the first event on day 1.
Attend the second event on day 2.
Attend the third event on day 3.
```


Example 2:

```
Input: events= [[1,2],[2,3],[3,4],[1,2]]
Output: 4
```

Example 3:

```
Input: events = [[1,4],[4,4],[2,2],[3,4],[1,1]]
Output: 4
```


Example 4:

```
Input: events = [[1,100000]]
Output: 1
```


Example 5:

```
Input: events = [[1,1],[1,2],[1,3],[1,4],[1,5],[1,6],[1,7]]
Output: 7
```
 

Constraints:

1 <= events.length <= 10^5  
events[i].length == 2  
1 <= events[i][0] <= events[i][1] <= 10^5  

## 方法

```go
import (
    "sort"
)

func maxEvents(events [][]int) int {
    var (
        i, j, k, maxEnd int
        length          = len(events)
        visitedDays     []bool
    )
    if length == 0 {
        return 0
    }
    maxEnd = events[0][1]
    sort.Slice(events, func(i, j int) bool {
        if events[i][1] > maxEnd {
            maxEnd = events[i][1]
        }
        if events[i][1] == events[j][1] {
            return events[i][0] > events[j][0]
        }
        return events[i][1] < events[j][1]
    })
    visitedDays = make([]bool, maxEnd+1)

    for i = 0; i < length; i++ {
        for j = events[i][0]; j <= events[i][1]; j++ {
            if !visitedDays[j] {
                visitedDays[j] = true
                break
            }
        }
    }

    for i = 0; i <= maxEnd; i++ {
        if visitedDays[i] {
            k++
        }
    }
    return k
}
```

```python
class Solution(object):
    def maxEvents(self, events):
        """
        :type events: List[List[int]]
        :rtype: int
        """
        events.sort(key=lambda event: event[1]) # sort by enddate
        attended = set() 
        for start, end in events:
            while start in attended and start < end:
                start += 1
            attended.add(start)
        return len(attended)
```


```python
class Solution(object):
    def maxEvents(self, events):
        """
        :type events: List[List[int]]
        :rtype: int
        """
        ans = 0
        end = list()
        events = sorted(events,reverse=True)
        for i in range(1,100010,1):
            while events and events[-1][0] == i:
                heapq.heappush(end, events.pop()[1])
            while end and end[0] < i:
                heapq.heappop(end)
            if end:
                heapq.heappop(end)
                ans += 1
        return ans
```