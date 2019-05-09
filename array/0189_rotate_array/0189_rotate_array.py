# -*- coding: utf-8 -*-

class Solution(object):
    def rotate(self, nums, k):
        k %= len(nums)
        self.reverse(nums, 0, len(nums))
        self.reverse(nums, 0, k)
        self.reverse(nums, k, len(nums))

    def reverse(self, nums, start, end):
        while start < end:
            nums[start], nums[end - 1] = nums[end - 1], nums[start]
            start += 1
            end -= 1

    def rotate2(self, nums, k):
        n = len(nums)
        idx = 0
        distance = 0
        cur = nums[0]
        for x in range(n):
            idx = (idx + k) % n
            nums[idx], cur = cur, nums[idx]
            
            # print nums, cur
            distance = (distance + k) % n
            if distance == 0: #  若已循环一次，则重置索引，将索引+1，cur 等于索引值再次循环
                idx = (idx + 1) % n
                cur = nums[idx]



a = [1,2,3,4,5,6]
Solution().rotate2(a, 2)
print(a)