class Solution(object):
    def topKFrequent(self, nums, k):
        hs = {}
        frq = {}
        for i in xrange(0, len(nums)):
            if nums[i] not in hs:
                hs[nums[i]] = 1
            else:
                hs[nums[i]] += 1
        
        for z,v in hs.iteritems():
            if v not in frq:
                frq[v] = [z]
            else:
                frq[v].append(z)
        
        arr = []

        for x in xrange(len(nums), 0, -1):
            if x in frq:
                for i in frq[x]:
                    arr.append(i)
        
        return [arr[x] for x in xrange(0, k)]

nums = [1,1,1,2,2,3]
k = 2

res = Solution().topKFrequent(nums, k)
print(res)