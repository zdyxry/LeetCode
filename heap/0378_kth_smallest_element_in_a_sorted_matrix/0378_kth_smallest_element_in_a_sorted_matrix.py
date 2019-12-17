
class Solution(object):
    def kthSmallest(self, matrix, k):
        left = matrix[0][0]
        right = matrix[-1][-1] +1
        while left < right:
            mid = (left + right) / 2
            i = len(matrix) - 1
            j = 0
            count = 0
            while i >= 0 and j < len(matrix[0]):
                if matrix[i][j] <= mid:
                    j += 1
                    count += (i+1)
                else:
                    i -= 1

            if count < k:
                left = mid + 1
            else:
                right = mid
        return left

    
matrix = [[1,5,9],[10,11,13],[12,13,15]]
k = 8
res = Solution().kthSmallest(matrix,k)
print(res)