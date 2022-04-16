class Solution:
    def largestInteger(self, num: int) -> int:
        even = []
        odd = []
        for idx, v in enumerate(str(num)):
            if int(v) % 2 == 0:
                even.append(v)
            if int(v) % 2 == 1:
                odd.append(v)
        even.sort(reverse=True)
        odd.sort(reverse=True)
        result = []
        for i in str(num):
            if int(i) % 2 == 0:
                result.append(even[0])
                even = even[1:]
            else:
                result.append(odd[0])
                odd = odd[1:]
        return int(''.join(result))