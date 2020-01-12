
class Solution(object):
    def numTrees(self, n):
        if n == 0 :
            return 1

        def combination(n, k):
            count = 1
            for i in xrange(1, k+1):
                count = count * (n-i+1)/i
            return count
        
        return combination(2 * n, n) - combination(2 * n, n-1)