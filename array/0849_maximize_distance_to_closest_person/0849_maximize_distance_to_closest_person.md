849. Maximize Distance to Closest Person

Easy

In a row of seats, 1 represents a person sitting in that seat, and 0 represents that the seat is empty. 

There is at least one empty seat, and at least one person sitting.

Alex wants to sit in the seat such that the distance between him and the closest person to him is maximized. 

Return that maximum distance to closest person.

Example 1:

```
Input: [1,0,0,0,1,0,1]
Output: 2
Explanation: 
If Alex sits in the second open seat (seats[2]), then the closest person has distance 2.
If Alex sits in any other open seat, the closest person has distance 1.
Thus, the maximum distance to the closest person is 2.
```

Example 2:

```
Input: [1,0,0,0]
Output: 3
Explanation: 
If Alex sits in the last seat, the closest person is 3 seats away.
This is the maximum distance possible, so the answer is 3.
```

Note:

1 <= seats.length <= 20000

seats contains only 0s or 1s, at least one 0, and at least one 1.


# 方法
遍历数组，双指针，左指针指向最后 1 出现的索引，右指针指向当前索引，若当前索引为 1 时，则计算两个 1 之间的距离，除以 2，并更新做指针。

需要考虑两侧无人情况。

```go
func maxDistToClosest(seats []int) int {
    start := 0
	max := 0
	var i int
	for ; i < len(seats); i++ {
		if seats[i] == 1 {
			tmp := i - start
			if start != 0 {
				tmp = (tmp + 1) / 2
			}
			if max < tmp {
				max = tmp
			}
			start = i + 1
		}
	}
	if max < i-start {
		max = i - start
	}

	return max
}

```


```python
class Solution(object):
    def maxDistToClosest(self, seats):
        """
        :type seats: List[int]
        :rtype: int
        """
        left, count, _max_distance = -1, 0, 1
        for i, x in enumerate(seats):
            if x == 0:
                count += 1
            else:
                if left < 0:
                    # 左边没有1最大距离就是count
                    distance = count
                else:
                    # 左边有1最大距离就是一半左右
                    # 奇数个0正好在中间
                    # 偶数个0要比一半小1
                    distance = count // 2 + count % 2
                # 重置左边第一个1的位置
                # 遇到1重新计数0的个数
                left, count = i, 0
                # 更新最大距离
                _max_distance = max(_max_distance, distance)
        # 末尾一连串0的情况也可能刷新最大距离
        return max(_max_distance, count)

```