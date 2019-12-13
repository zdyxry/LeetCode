
import heapq 

class Solution(object):
    def nthUglyNumber1(self, n):
        u = self.nums()
        return u[n]

    def nums(self):
        ret = [0, 1]
        i2 = i3 = i5 = 1
        for i in range(2, 1691):
            n2 = ret[i2] * 2
            n3 = ret[i3] * 3
            n5 = ret[i5] * 5
            n = min(n2, n3, n5)
            ret.append(n)
            if n == n2:
                i2 += 1
            if n == n3:
                i3 += 1
            if n == n5:
                i5 += 1
        return ret

    def nthUglyNumber2(self, n):
        ugly_number = 0

        heap = []
        heapq.heappush(heap, 1)
        for i in xrange(n):
            ugly_number = heapq.heappop(heap)
            if ugly_number % 2 == 0:
                heapq.heappush(heap, ugly_number * 2)
            elif ugly_number % 3 == 0:
                heapq.heappush(heap, ugly_number * 2)
                heapq.heappush(heap, ugly_number * 3)
            else:
                heapq.heappush(heap, ugly_number * 2)
                heapq.heappush(heap, ugly_number * 3)
                heapq.heappush(heap, ugly_number * 5)
            print i, ugly_number, heap
        return ugly_number


n = 10
res = Solution().nthUglyNumber2(n)
print(res)