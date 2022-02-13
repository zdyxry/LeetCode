class Solution:
    def timeRequiredToBuy(self, tickets: List[int], k: int) -> int:
        result = 0
        while tickets[k] > 0:
            for i in range(len(tickets)):
                if tickets[i] >= 1:
                    tickets[i] -= 1
                    result += 1
                if tickets[k] == 0:
                    break
        return result
