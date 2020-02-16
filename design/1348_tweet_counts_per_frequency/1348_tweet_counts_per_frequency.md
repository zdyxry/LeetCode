1348. Tweet Counts Per Frequency


Medium


Implement the class TweetCounts that supports two methods:

1. recordTweet(string tweetName, int time)

Stores the tweetName at the recorded time (in seconds).

2. getTweetCountsPerFrequency(string freq, string tweetName, int startTime, int endTime)

Returns the total number of occurrences for the given tweetName per minute, hour, or day (depending on freq) starting from the startTime (in seconds) and ending at the endTime (in seconds).

freq is always minute, hour or day, representing the time interval to get the total number of occurrences for the given tweetName.
The first time interval always starts from the startTime, so the time intervals are [startTime, startTime + delta*1>,  [startTime + delta*1, startTime + delta*2>, [startTime + delta*2, startTime + delta*3>, ... , [startTime + delta*i, min(startTime + delta*(i+1), endTime + 1)> for some non-negative number i and delta (which depends on freq).  
 

Example:
```
Input
["TweetCounts","recordTweet","recordTweet","recordTweet","getTweetCountsPerFrequency","getTweetCountsPerFrequency","recordTweet","getTweetCountsPerFrequency"]
[[],["tweet3",0],["tweet3",60],["tweet3",10],["minute","tweet3",0,59],["minute","tweet3",0,60],["tweet3",120],["hour","tweet3",0,210]]

Output
[null,null,null,null,[2],[2,1],null,[4]]

Explanation
TweetCounts tweetCounts = new TweetCounts();
tweetCounts.recordTweet("tweet3", 0);
tweetCounts.recordTweet("tweet3", 60);
tweetCounts.recordTweet("tweet3", 10);                             // All tweets correspond to "tweet3" with recorded times at 0, 10 and 60.
tweetCounts.getTweetCountsPerFrequency("minute", "tweet3", 0, 59); // return [2]. The frequency is per minute (60 seconds), so there is one interval of time: 1) [0, 60> - > 2 tweets.
tweetCounts.getTweetCountsPerFrequency("minute", "tweet3", 0, 60); // return [2, 1]. The frequency is per minute (60 seconds), so there are two intervals of time: 1) [0, 60> - > 2 tweets, and 2) [60,61> - > 1 tweet.
tweetCounts.recordTweet("tweet3", 120);                            // All tweets correspond to "tweet3" with recorded times at 0, 10, 60 and 120.
tweetCounts.getTweetCountsPerFrequency("hour", "tweet3", 0, 210);  // return [4]. The frequency is per hour (3600 seconds), so there is one interval of time: 1) [0, 211> - > 4 tweets.
```

Constraints:

There will be at most 10000 operations considering both recordTweet and getTweetCountsPerFrequency.

0 <= time, startTime, endTime <= 10^9

0 <= endTime - startTime <= 10^4

## 方法

```go
type TweetCounts struct {
    Tweet map[string][]int // key: name val: []int times
}

func Constructor() TweetCounts {
    return TweetCounts{
        Tweet: map[string][]int{},
    }
}


func (this *TweetCounts) RecordTweet(tweetName string, time int)  {
    if _, ok := this.Tweet[tweetName]; !ok {
        this.Tweet[tweetName] = []int{time}
    } else {
        this.Tweet[tweetName] = append(this.Tweet[tweetName], time)
    }
}


func (this *TweetCounts) GetTweetCountsPerFrequency(freq string, tweetName string, startTime int, endTime int) []int {
    interval := GetInterval(freq)
	
    if _, ok := this.Tweet[tweetName]; !ok {
        return []int{}
    }
    
    result := make([]int, (endTime - startTime + interval) / interval)
    for _, time := range this.Tweet[tweetName] {
        if time > endTime || time < startTime {
            continue
        }
        
        thisInterval := (time-startTime)/interval
        result[thisInterval]++
    }
    
    return result
}


func GetInterval(freq string) int {
    if freq == "minute" {
        return 60
    }
    
    if freq == "hour" {
        return 3600
    }
    
    if freq == "day" {
        return 86400
    }
    
    return 0 // never reach here
}

/**
 * Your TweetCounts object will be instantiated and called as such:
 * obj := Constructor();
 * obj.RecordTweet(tweetName,time);
 * param_2 := obj.GetTweetCountsPerFrequency(freq,tweetName,startTime,endTime);
 */
 ```



```python
class TweetCounts(object):

    def __init__(self):
        self.__records = collections.defaultdict(list)
        self.__lookup = {"minute":60, "hour":3600, "day":86400}

    def recordTweet(self, tweetName, time):
        """
        :type tweetName: str
        :type time: int
        :rtype: None
        """
        self.__records[tweetName].append(time)

    def getTweetCountsPerFrequency(self, freq, tweetName, startTime, endTime):
        """
        :type freq: str
        :type tweetName: str
        :type startTime: int
        :type endTime: int
        :rtype: List[int]
        """
        delta = self.__lookup[freq]
        result = [0]*((endTime- startTime)//delta+1)
        for t in self.__records[tweetName]:
            if startTime <= t <= endTime:
                result[(t-startTime)//delta] += 1
        return result
```