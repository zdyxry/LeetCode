
class Solution(object):
    def asteroidCollision(self, asteroids):
        res = []
        for asteroid in asteroids:
            if len(res) == 0 or asteroid > 0:
                res.append(asteroid)
            elif asteroid < 0:
                while len(res) and res[-1] > 0:
                    if res[-1] == -asteroid:
                        res.pop()
                        break
                    elif res[-1] < -asteroid:
                        res.pop()
                        continue
                    elif res[-1] > -asteroid:
                        break
                else:
                    res.append(asteroid)
        return res


asteroids = [-2,-1,1,2]
res = Solution().asteroidCollision(asteroids)
print(res)