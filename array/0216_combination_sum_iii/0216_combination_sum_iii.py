class Solution(object):
    def combinationSum3(self, k, n):
        """
        :type k: int
        :type n: int
        :rtype: List[List[int]]
        """
        paths = []
        self.helper(k, n, 0, 1, 0, [], paths)
        return paths
    
    def helper(self, k, n, cur_sum, cur_num, cur_cnt, path, paths):
        if cur_sum == n:
            if cur_cnt == k:
                paths.append(path[:])
            return
        
        for next_num in xrange(cur_num, 10):
            if cur_sum + next_num <= n and cur_cnt + 1 <= k:
                path.append(next_num)
                self.helper(k, n, cur_sum+next_num, next_num+1, cur_cnt+1, path, paths)
                path.pop()


k = 3
n = 7
res = Solution().combinationSum3(k, n)
print(res)