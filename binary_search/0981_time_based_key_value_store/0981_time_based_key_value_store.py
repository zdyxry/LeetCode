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