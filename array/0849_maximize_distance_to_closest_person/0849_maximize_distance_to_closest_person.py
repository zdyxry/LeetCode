# -*- coding: utf-8 -*-


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


print(Solution().maxDistToClosest([1,0,0,0,1,0,1]))