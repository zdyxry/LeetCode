class Solution:
    def findKDistantIndices(self, nums: List[int], key: int, k: int) -> List[int]:
        res = []
        n = len(nums)
        # 遍历数对
        for i in range(n):
            for j in range(n):
                if nums[j] == key and abs(i - j) <= k:
                    res.append(i)
                    break   # 提前结束以防止重复添加
        return res
