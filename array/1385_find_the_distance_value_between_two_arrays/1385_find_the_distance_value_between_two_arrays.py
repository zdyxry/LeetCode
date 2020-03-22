
class Solution(object):
    def findTheDistanceValue(self, arr1, arr2, d):
        ans = 0
        for n1 in arr1:
            if all(abs(n1-n2) > d for n2 in arr2):
                ans+=1
        return ans


arr1= [4,5,8]
arr2 = [10,9,1,8]
d = 2
res = Solution().findTheDistanceValue(arr1, arr2, d)
print(res)