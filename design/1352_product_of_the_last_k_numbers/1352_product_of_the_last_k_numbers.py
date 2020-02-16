
class ProductOfNumbers(object):
    def __init__(self):
        self.A = [1]

    def add(self, a):
        if a == 0:
            self.A = [1]
        else:
            self.A.append(self.A[-1] * a)
    
    def getProduct(self, k):
        if k >= len(self.A):
            return 0
        return self.A[-1] / self.A[-k-1]

