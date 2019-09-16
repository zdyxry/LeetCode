class Solution(object):
    def searchRange(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        def get_start_index(nums, target):
            start, end = 0, len(nums) - 1
            while start <= end:
                mid = start + (end - start) // 2
                if nums[mid] == target:
                    if mid == 0 or nums[mid-1] != target:
                        return mid
                if nums[mid] < target:
                    start = mid + 1
                else:
                    end = mid - 1
            return -1
        
        def get_end_index(nums, target):
            start, end = 0, len(nums) - 1
            while start <= end:
                mid = start + (end - start)//2
                if nums[mid] == target: 
                    if mid == len(nums)-1 or nums[mid + 1] != target:
                        return mid
                if nums[mid] <= target:
                    start = mid + 1
                else:
                    end = mid - 1
            return -1
        
        return [get_start_index(nums, target), get_end_index(nums, target)]

nums = [5,7,7,8,8,10]
target = 8 

res = Solution().searchRange(nums, target)
print(res)