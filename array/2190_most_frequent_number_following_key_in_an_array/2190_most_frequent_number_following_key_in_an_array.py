class Solution:
    def mostFrequent(self, nums: List[int], key: int) -> int:
        l = len(nums)
        result = {}
        for idx, n in enumerate(nums):
            if n == key:
                if idx + 1 >= len(nums):
                    continue
                if nums[idx+1] in result:
                    result[nums[idx+1]] += 1
                else:
                    result[nums[idx+1]] = 1
        i = 0
        m = 0
        for k,v in result.items():
            if v > m:
                m = v
                i = k
        return i
            


