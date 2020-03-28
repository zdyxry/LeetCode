
class Solution(object):
    def carPooling(self, trips, capacity):
        dcap = [0] * 1001
        for p, start, end in trips:
            dcap[start] -= p
            dcap[end] += p
        
        for delta in dcap:
            capacity += delta
        
            if capacity < 0:
                return False
        
        return True


trips = [[2,1,5],[3,3,7]]
capacity = 4
res = Solution().carPooling(trips, capacity)
print(res)