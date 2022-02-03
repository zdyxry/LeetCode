class Solution:
    def sumOfBeauties(self, nums: List[int]) -> int:
        n = len(nums)
        ans = 0
        leftMax,rightMin = [0] * n, [0] * n
        leftMax[0] = nums[0]
        rightMin[-1] = nums[-1]
        # 计算左边的最大值
        for i in range(1, n):
            if nums[i] > leftMax[i-1]:
                leftMax[i] = nums[i]
            else:
                leftMax[i] = leftMax[i-1]
        # 计算右边最小值
        for j in range(n-2, -1, -1):
            if nums[j] < rightMin[j + 1]:
                rightMin[j] = nums[j]
            else:
                rightMin[j] = rightMin[j+1]
        # 判断美丽值
        for i in range(1,n-1):
            if leftMax[i-1] < nums[i] < rightMin[i+1]:
                ans += 2
            elif nums[i-1] < nums[i] < nums[i+1]:
                ans += 1
            else:
                pass
        return ans
        