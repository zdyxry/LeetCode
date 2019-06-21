class Solution(object):
    def findRestaurant(self, list1, list2):
        """
        :type list1: List[str]
        :type list2: List[str]
        :rtype: List[str]
        """
        lookup = {}
        for i, s in enumerate(list1):
            lookup[s] = i

        result = []
        min_sum = float("inf")
        for j, s in enumerate(list2):
            if j > min_sum:
                break
            if s in lookup:
                if j + lookup[s] < min_sum:
                    result = [s]
                    min_sum = j + lookup[s]
                elif j + lookup[s] == min_sum:
                    result.append(s)
        return result

list1 = ["Shogun","Tapioca Express","Burger King","KFC"]
list2 = ["Piatti","The Grill at Torrey Pines","Hungry Hunter Steakhouse","Shogun"]
print(Solution().findRestaurant(list1, list2))