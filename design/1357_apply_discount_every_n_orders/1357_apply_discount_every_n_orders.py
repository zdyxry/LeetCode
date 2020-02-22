
class Solution(object):
    def __init(self, n, discount, products, prices):
        self.price = {p : prices[i] for i,p in enumerate(products)}
        self.discount = float(discount)
        self.n = n
        self.cnt = 0

    def getBill(self, product, amount):
        self.cnt += 1
        total = 0
        for i,p in enumerate(product):
            total += self.price[p] * amount[i]
        
        if self.cnt % self.n == 0 :
            total = total * (1 - self.discount/100)
        return total