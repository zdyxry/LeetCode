# -*- coding: utf-8 -*-


class Solution(object):

    def fib(self, N):
        if N <= 1:
            return N
        return self.fib(N - 1) + self.fib(N - 2)

    def fib2(self, N):
        array = [i for i in xrange(N)]
        for i in xrange(2, N+1):
            array[i] = array[i-1] + array[i-2]
        return array[-1]

    def fib3(self, N):
        if N <=1:
            return N
        left = 0
        right =1
        for i in range(2,N+1):
            left, right = right, left + right
        return right

    def fib4(self, N):
        array =[i for i in range(N+1)]
        return self.fibola(array, N)

    def fibola(self, array, N):
        if N <= 1:
            return N
        array[N] = self.fibola(array, N-1) + array[N-2]
        return array[N]



print(Solution().fib4(6))