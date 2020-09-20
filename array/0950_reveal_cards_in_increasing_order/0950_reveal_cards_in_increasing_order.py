from typing import List
import collections


class Solution:
    def deckRevealedIncreasing(self, deck: List[int]) -> List[int]:
        N = len(deck)
        index = collections.deque(range(N))
        ans = [None] * N

        for card in sorted(deck):
            ans[index.popleft()] = card
            if index:
                index.append(index.popleft())

        return ans


deck = [17,13,11,2,3,5,7]
res = Solution().deckRevealedIncreasing(deck)
print(res)