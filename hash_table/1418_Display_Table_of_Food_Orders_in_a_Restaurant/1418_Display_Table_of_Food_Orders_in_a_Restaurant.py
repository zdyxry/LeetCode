from typing import List

from collections import defaultdict


class Solution:
    def displayTable(self, orders: List[List[str]]) -> List[List[str]]:
        data = defaultdict(lambda: defaultdict(int))
        food_items = set()
        for order in orders:
            table, food = int(order[1]), order[2]
            data[table][food] += 1
            food_items.add(food)
        sorted_food_items = sorted(list(food_items))
        result = [["Table"] + sorted_food_items]
        for table in sorted(data.keys()):
            result.append([str(table)] + [str(data[table][f]) for f in sorted_food_items])
        return result


orders = [["David","3","Ceviche"],["Corina","10","Beef Burrito"],["David","3","Fried Chicken"],["Carla","5","Water"],["Carla","5","Ceviche"],["Rous","3","Ceviche"]]
res = Solution().displayTable(orders)
print(res)