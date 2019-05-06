class Solution(object):
    def generate(self, numRows):
        """
        :type numRows: int
        :rtype: List[List[int]]
        """
        result = []
        for i in range(numRows):
            result.append([])
            for j in range(0, i+1):
                if j in (0,i):
                    result[i].append(1)
                else:
                    result[i].append(result[i-1][j-1] + result[i-1][j])
                    
        return result


print(Solution().generate(5))