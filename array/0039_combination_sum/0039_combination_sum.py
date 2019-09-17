class Solution(object):
    def combinationSum(self, candidates, target):
        result = []
        self.combinationSumRecu(sorted(candidates), result, 0, [], target)
        return result

    def combinationSumRecu(self, candidates, result, start, intermediate, target):
        if target == 0:
            result.append(list(intermediate))
        while start < len(candidates) and candidates[start] <= target:
            intermediate.append(candidates[start])
            self.combinationSumRecu(candidates, result, start, intermediate, target-candidates[start])
            intermediate.pop()
            start += 1

    
candidates = [2,3,6,7]
target = 7

res = Solution().combinationSum(candidates, target)
print(res)