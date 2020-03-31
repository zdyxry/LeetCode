class UndergroundSystem:

    def __init__(self):
        self.count = defaultdict(int)
        self.time = defaultdict(int)
        self.traveling = dict()
        
    def checkIn(self, id: int, stationName: str, t: int) -> None:
        self.traveling[id] = (stationName, t)

    def checkOut(self, id: int, stationName: str, t: int) -> None:
        (prev_station, prev_t) = self.traveling[id]
        del self.traveling[id]
        key = (prev_station, stationName)
        self.count[key] += 1
        self.time[key] += (t-prev_t)

    def getAverageTime(self, startStation: str, endStation: str) -> float:
        key = (startStation, endStation)
        
        return self.time[key] / self.count[key]
        


# Your UndergroundSystem object will be instantiated and called as such:
# obj = UndergroundSystem()
# obj.checkIn(id,stationName,t)
# obj.checkOut(id,stationName,t)
# param_3 = obj.getAverageTime(startStation,endStation)