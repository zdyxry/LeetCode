
class Solution(object):
    def fizzBuzz(self, n):
        result = []

        for i in xrange(1, n+1):
            if i % 15 == 0:
                result.append('FizzBuzz')
            elif i % 5 == 0:
                result.append('Buzz')
            elif i % 3 == 0:
                result.append('Fizz')
            else:
                result.append(str(i))
        return result


n = 15
res = Solution().fizzBuzz(n)
print(res)