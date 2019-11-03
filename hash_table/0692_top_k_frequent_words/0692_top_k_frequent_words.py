import heapq

class Solution(object):
    def topKFrequent(self, words, k):
        u = {}
        for i in range(len(words)):
            if words[i] not in u:
                u[words[i]] = 1
            else:
                u[words[i]] += 1
        u_count = {}
        heap = set()
        for key, val in u.items():
            if val not in u_count:
                u_count[val] = []
            u_count[val].append(key)
            heap.add(val*-1)
        heap=list(heap)

        heapq.heapify(heap)
        res = []
        idx = 0
        while idx < k:
            vals = u_count[heapq.heappop(heap)*-1]
            heapq.heapify(heap)
            vals.sort()
            for i in range(len(vals)):
                if len(res) == k:
                    return res
                res.append(vals[i])
                idx += 1
        return res

words = ["i", "love", "leetcode", "i", "love", "coding"]
k = 2

res = Solution().topKFrequent(words, k)
print(res)