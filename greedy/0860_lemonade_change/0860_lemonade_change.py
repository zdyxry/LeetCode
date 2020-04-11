from typing import List


class Solution:
    def lemonadeChange(self, bills: List[int]) -> bool:
        previous_5 = 0
        previous_10 = 0
        price = 5
        for bill in bills:
            if bill - price == 5:
                previous_5 -= 1
                previous_10 += 1
            elif bill - price == 15:
                previous_5 -= 1
                if previous_10 > 0:
                    previous_10 -= 1
                else:
                    previous_5 -= 2
            else:
                previous_5 += 1
            if previous_5 < 0:
                return False
        return True

bills = [5,5,5,10,20]
res = Solution().lemonadeChange(bills)
print(res)