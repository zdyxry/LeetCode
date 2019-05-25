605. Can Place Flowers

Easy

Suppose you have a long flowerbed in which some of the plots are planted and some are not. However, flowers cannot be planted in adjacent plots - they would compete for water and both would die.



Given a flowerbed (represented as an array containing 0 and 1, where 0 means empty and 1 means not empty), and a number n, return if n new flowers can be planted in it without violating the no-adjacent-flowers rule.


Example 1:
```
Input: flowerbed = [1,0,0,0,1], n = 1
Output: True
```
Example 2:
```
Input: flowerbed = [1,0,0,0,1], n = 2
Output: False
```

Note:

The input array won't violate no-adjacent-flowers rule.

The input array size is in the range of [1, 20000].

n is a non-negative integer which won't exceed the input array size.

# 方法
 遍历数组，当当前数值为 0 且前后均为 0 时，赋值当前值为 1。


```go
func canPlaceFlowers(flowerbed []int, n int) bool {
    l := len(flowerbed)
	for i := 0; i < l; i++ {
		if flowerbed[i] == 0 && 
			((i+1 < l && flowerbed[i+1] == 0) || i+1 >= l) &&
			((i-1 >= 0 && flowerbed[i-1] == 0) || i-1 < 0) { 
			flowerbed[i] = 1
			n--
		}
	}

	return n <= 0
}
```


```python
class Solution(object):
    def canPlaceFlowers(self, flowerbed, n):
        """
        :type flowerbed: List[int]
        :type n: int
        :rtype: bool
        """
        size = len(flowerbed)
        for i in xrange(size):
            if flowerbed[i] == 0:
                avail = True
                if i - 1 >= 0 and flowerbed[i-1]==1:
                    avail = False
                if i + 1 < size and flowerbed[i+1]==1:
                    avail = False
                if avail:
                    flowerbed[i] = 1
                    n -= 1
        return n <= 0
```