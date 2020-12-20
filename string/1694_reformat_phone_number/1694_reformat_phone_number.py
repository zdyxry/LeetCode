

class Solution(object):
    def reformatNumber(self, number):
        """
        :type number: str
        :rtype: str
        """
        number = number.replace('-' ,'').replace(' ', '')
        ans = []
        for i in range(0, len(number), 3):
            ans.append(number[i:i+3])

        if len(ans) >= 2 and len(ans[-1]) == 1:
            ans[-1] = ans[-2][2] + ans[-1]
            ans[-2] = ans[-2][:2]
        return '-'.join(ans)


number = "1-23-45 6"
res = Solution().reformatNumber(number)
print(res)