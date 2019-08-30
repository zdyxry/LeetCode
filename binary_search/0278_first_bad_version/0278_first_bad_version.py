class Solution(object):
    def isBadVersion(self, num):
        return True

    def firstBadVersion(self, n):
        left, right = 1, n
        while left <= right:
            mid = left + (right - left)/2
            if self.isBadVersion(mid):
                right = mid - 1
            else:
                left = mid + 1
        return left