class Solution:
    def maximumTime(self, time: str) -> str:
        hh = time[:2]
        if hh[0] == '?' and hh[1] == '?':
            hh = '23'
        elif hh[0] == '?':
            if hh[1] < '4':
                hh = f'2{hh[1]}'
            else:
                hh = f'1{hh[1]}' 
        elif hh[1] == '?':
            if hh[0] == '2':
                hh = f'{hh[0]}3'
            else:
                hh = f'{hh[0]}9'
        mm = time[3:]
        if mm[0] == '?':
            mm = f'5{mm[1]}'
        if mm[1] == '?':
            mm = f'{mm[0]}9'
        return f'{hh}:{mm}'


time = "2?:?0"
res = Solution().maximumTime(time)
print(res)