981. Time Based Key-Value Store


Medium


Create a timebased key-value store class TimeMap, that supports two operations.

```
1. set(string key, string value, int timestamp)

Stores the key and value, along with the given timestamp.
2. get(string key, int timestamp)

Returns a value such that set(key, value, timestamp_prev) was called previously, with timestamp_prev <= timestamp.
If there are multiple such values, it returns the one with the largest timestamp_prev.
If there are no values, it returns the empty string ("").
```

Example 1:

```
Input: inputs = ["TimeMap","set","get","get","set","get","get"], inputs = [[],["foo","bar",1],["foo",1],["foo",3],["foo","bar2",4],["foo",4],["foo",5]]
Output: [null,null,"bar","bar",null,"bar2","bar2"]
Explanation:   
TimeMap kv;   
kv.set("foo", "bar", 1); // store the key "foo" and value "bar" along with timestamp = 1   
kv.get("foo", 1);  // output "bar"   
kv.get("foo", 3); // output "bar" since there is no value corresponding to foo at timestamp 3 and timestamp 2, then the only value is at timestamp 1 ie "bar"   
kv.set("foo", "bar2", 4);   
kv.get("foo", 4); // output "bar2"   
kv.get("foo", 5); //output "bar2"   
```

Example 2:

```
Input: inputs = ["TimeMap","set","set","get","get","get","get","get"], inputs = [[],["love","high",10],["love","low",20],["love",5],["love",10],["love",15],["love",20],["love",25]]
Output: [null,null,null,"","high","high","low","low"]
```

Note:

All key/value strings are lowercase.  
All key/value strings have length in the range [1, 100]  
The timestamps for all TimeMap.set operations are strictly increasing.  
1 <= timestamp <= 10^7  
TimeMap.set and TimeMap.get functions will be called a total of 120000 times (combined) per test case.


## 方法


```go
type TimeMap struct {
    timestamp map[string][]int
    realValue map[string][]string
}

func Constructor() TimeMap {
    return TimeMap{
        timestamp: make(map[string][]int),
        realValue: make(map[string][]string),
    }
}

func (this *TimeMap) Set(key string, value string, timestamp int)  {
    this.timestamp[key] = append(this.timestamp[key], timestamp)
    this.realValue[key] = append(this.realValue[key], value)
}

func (this *TimeMap) Get(key string, timestamp int) string {
    if len(this.timestamp[key]) == 0 {
        return ""
    }
    index := findFirstLEIndex(this.timestamp[key], timestamp)
    if index < 0 {
        return ""
    }
    return this.realValue[key][index]
}

func findFirstLEIndex(nums []int, n int) int {
    l, r := 0, len(nums)-1
    for l <= r {
        m := (l + r) >> 1
        if nums[m] > n {
            r = m - 1
        } else {
            l = m + 1
        }
    }
    return r
}
```


```python
class TimeMap:

    def __init__(self):
        """
        Initialize your data structure here.
        """
        self.dic={}

    def set(self, key: str, value: str, timestamp: int) -> None:
        if key not in self.dic.keys():
            self.dic[key]=[[timestamp,value]]
        else:
            self.dic[key].append([timestamp,value])

    def get(self, key: str, timestamp: int) -> str:
        data=self.dic[key]
        if data[-1][0]<=timestamp:
            return data[-1][1]
        if data[0][0]>timestamp:
            return ""
        if data[0][0]==timestamp:
            return data[0][1]
        lo,ro=0,len(data)-1
        while lo<ro:
            mid=(lo+ro)//2
            if data[mid][0]<=timestamp:
                lo=mid+1
            else:
                ro=mid-1
        return data[mid][1]
        


# Your TimeMap object will be instantiated and called as such:
# obj = TimeMap()
# obj.set(key,value,timestamp)
# param_2 = obj.get(key,timestamp)
```