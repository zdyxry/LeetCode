
class Solution(object):
    def isPossible(self, target):
        while True:
            current_x = max(target)
            if current_x == 1:
                return True

            idx = target.index(current_x)
            s = sum(target)
            inc = s - current_x
            if inc > current_x or inc ==0:
                return False

            target[idx] = current_x % inc


target = [9,3,5]
res = Solution().isPossible(target)
print(res)